package rewrite

import (
	"bytes"
	"fmt"
	"io"

	// "io/ioutil"

	"regexp"
	"strconv"

	"golang.org/x/net/html"
)

var headTags = []string{"html", "head", "base", "link", "meta", "title", "style", "script", "object", "bgsound"}
var beforeHeadTags = []string{"html", "head"}
var dataRwProtocols = []string{"http://", "https://", "//"}

type HtmlRewriter struct {
	urlrw         Rewriter
	jsRewriter    Rewriter
	cssRewriter   Rewriter
	url           string
	defmod        Rewriter
	parseComments bool
	rewriteTags   map[string]map[string]Rewriter
	r             io.Reader
	inserts       []insert
	// rewriters     map[RewriterType]Rewriter
}

func NewHtmlRewriter(urlrw Rewriter, configs ...func(*Config)) *HtmlRewriter {
	// c := makeConfig(configs...)
	return &HtmlRewriter{
		urlrw: urlrw,
		// jsRewriter:  c.Rewriters[RwTypeJavascript],
		rewriteTags: rewriteTags(urlrw),
	}
}

func NewHtmlJSRewriter(urlrw Rewriter, jsrw Rewriter, configs ...func(*Config)) *HtmlRewriter {
	// c := makeConfig(configs...)
	return &HtmlRewriter{
		urlrw:       urlrw,
		jsRewriter:  jsrw,
		rewriteTags: rewriteTags(urlrw),
	}
}

func (hr *HtmlRewriter) AddInsert(name, value string) {
	hr.inserts = append(hr.inserts, insert{Name: name, Value: value})
}

func (hrw *HtmlRewriter) Rewrite(p []byte) []byte {
	// rdr := bytes.NewReader(p)

	w := &bytes.Buffer{}
	err := hrw.rewrite(bytes.NewReader(p), w)
	if err != nil {
		return p
	}
	return w.Bytes()

}

func (hrw *HtmlRewriter) rewrite(rdr io.Reader, ww io.Writer) error {
	r := hrw.NewReader(rdr)
	_, err := io.Copy(ww, r)
	return err
	tokenizer := html.NewTokenizer(rdr)
	// tokenizer := html.NewTokenizerFragment(rdr, "head")
	var isScript bool
	for {
		tt := tokenizer.Next()
		switch tt {
		case html.ErrorToken:
			if tokenizer.Err().Error() == "EOF" {
				return nil
			}
			return tokenizer.Err()
		case html.StartTagToken:
			name, hasAttr := tokenizer.TagName()
			token := html.Token{
				Type: html.StartTagToken,
				Data: string(name),
			}
			if hasAttr {
				hrw.rewriteToken(&token, tokenizer)
			}
			tkstring := String(token)
			ww.Write([]byte(tkstring))
			// isScript = bytes.Compare(name, []byte(RwTypeScript.String())) == 0
			isScript = r.isScript(token)
			continue
		case html.SelfClosingTagToken:
			name, hasAttr := tokenizer.TagName()
			token := html.Token{
				Type: html.SelfClosingTagToken,
				Data: string(name),
			}

			if hasAttr {
				hrw.rewriteToken(&token, tokenizer)
			}

			ww.Write([]byte(String(token)))
			continue
		case html.TextToken:
			if isScript && hrw.jsRewriter != nil {
				isScript = false
				ww.Write(hrw.jsRewriter.Rewrite(tokenizer.Raw()))
				continue
			}
			isScript = false

		}
		fmt.Println("tokenizer.Raw()", string(tokenizer.Raw()))
		ww.Write(tokenizer.Raw())
	}
	return nil
}

func (hrw *HtmlRewriter) rewriteMetaRefresh(p []byte, metaRefresh *regexp.Regexp) {

}

func (hrw *HtmlRewriter) rewriteToken(t *html.Token, tok *html.Tokenizer) {
	attrs := hrw.rewriteTags[t.Data]
	for {
		key, oldval, more := tok.TagAttr()
		var newVal []byte
		repl := attrs[string(bytes.ToLower(key))]
		if repl != nil {
			newVal = repl.Rewrite(oldval)
		}
		if newVal == nil {
			newVal = oldval
		}

		t.Attr = append(t.Attr, html.Attribute{
			Key: string(key),
			Val: string(newVal),
		})

		if bytes.Compare(newVal, oldval) != 0 {
			t.Attr = append(t.Attr, html.Attribute{
				Key: "__cpp",
				Val: "1",
			})
		}

		if !more {
			return
		}
	}
	return
}

func rewriteTags(defmod Rewriter) map[string]map[string]Rewriter {
	oe := PrefixRewriter{Prefix: []byte("oe_")}
	im := PrefixRewriter{Prefix: []byte("im_")}
	fr_ := PrefixRewriter{Prefix: []byte("fr_")}

	return map[string]map[string]Rewriter{
		"a":          {"href": defmod},
		"applet":     {"codebase": oe, "archive": oe},
		"area":       {"href": defmod},
		"audio":      {"src": oe},
		"base":       {"href": defmod},
		"blockquote": {"cite": defmod},
		"body":       {"background": im},
		"button":     {"formaction": defmod},
		"command":    {"icon": im},
		"del":        {"cite": defmod},
		"embed":      {"src": oe},
		"head":       {"": defmod}, // for head rewriting
		// "iframe":     {"src": if_},
		"iframe": {"src": defmod},
		"image":  {"src": defmod, "xlink:href": im},
		"img":    {"src": defmod, "srcset": im},
		"ins":    {"cite": defmod},
		"input":  {"src": im, "formaction": defmod},
		"form":   {"action": defmod},
		"frame":  {"src": fr_},
		"link":   {"href": defmod},
		"meta":   {"content": defmod},
		"object": {"codebase": oe, "data": oe},
		"param":  {"value": oe},
		"q":      {"cite": defmod},
		"ref":    {"href": oe},
		"script": {"src": defmod},
		"source": {"src": oe},
		"video":  {"src": oe, "poster": im},
	}
}

// TODO - resolve older verion:
// func rewriteTags(defmod Rewriter) map[string]map[string]Rewriter {
// 	oe := PrefixRewriter{Prefix: []byte("oe_")}
// 	im := PrefixRewriter{Prefix: []byte("im_")}
// 	if_ := PrefixRewriter{Prefix: []byte("if_")}
// 	fr_ := PrefixRewriter{Prefix: []byte("fr_")}
// 	js_ := PrefixRewriter{Prefix: []byte("js_")}

// 	return map[string]map[string]Rewriter{
// 		"a":          {"href": defmod},
// 		"applet":     {"codebase": oe, "archive": oe},
// 		"area":       {"href": defmod},
// 		"audio":      {"src": oe},
// 		"base":       {"href": defmod},
// 		"blockquote": {"cite": defmod},
// 		"body":       {"background": im},
// 		"button":     {"formaction": defmod},
// 		"command":    {"icon": im},
// 		"del":        {"cite": defmod},
// 		"embed":      {"src": oe},
// 		"head":       {"": defmod}, // for head rewriting
// 		"iframe":     {"src": if_},
// 		"image":      {"src": im, "xlink:href": im},
// 		"img":        {"src": im, "srcset": im},
// 		"ins":        {"cite": defmod},
// 		"input":      {"src": im, "formaction": defmod},
// 		"form":       {"action": defmod},
// 		"frame":      {"src": fr_},
// 		"link":       {"href": oe},
// 		"meta":       {"content": defmod},
// 		"object":     {"codebase": oe, "data": oe},
// 		"param":      {"value": oe},
// 		"q":          {"cite": defmod},
// 		"ref":        {"href": oe},
// 		"script":     {"src": js_},
// 		"source":     {"src": oe},
// 		"video":      {"src": oe, "poster": im},
// 	}
// }

func tagString(t html.Token) string {
	bf := tagBuffer(t)
	if bf != nil {
		return bf.String()
	}
	return ""
}

func tagBuffer(t html.Token) *bytes.Buffer {
	buf := bytes.NewBufferString(t.Data)

	if len(t.Attr) == 0 {
		return buf
	}
	for _, a := range t.Attr {
		buf.WriteByte(' ')
		buf.WriteString(a.Key)
		buf.WriteString(`="`)
		buf.WriteString(a.Val)
		// escape(buf, a.Val)
		buf.WriteByte('"')
	}
	return buf
}

// String returns a string representation of the Token.
func String(t html.Token) string {
	switch t.Type {
	case html.ErrorToken:
		return ""
	case html.TextToken:
		return html.EscapeString(t.Data)
	case html.StartTagToken:
		return "<" + tagString(t) + ">"
	case html.EndTagToken:
		return "</" + tagString(t) + ">"
	case html.SelfClosingTagToken:
		return "<" + tagString(t) + "/>"
	case html.CommentToken:
		return "<!--" + t.Data + "-->"
	case html.DoctypeToken:
		return "<!DOCTYPE " + t.Data + ">"
	}
	return "Invalid(" + strconv.Itoa(int(t.Type)) + ")"
}

func (hrw *HtmlRewriter) RewriteStream(f io.Reader, t io.Writer) error {
	return hrw.rewrite(f, t)
}

func (hrw *HtmlRewriter) NewReader(r io.Reader) *RewriteReader {
	var opts = []ReaderOption{
		SetHtmlRewriter(hrw),
	}
	for _, insert := range hrw.inserts {
		opts = append(opts, AddInsert(insert))
	}
	return NewRewriteReader(r,
		opts...,
	)
}

type ReaderOption func(*RewriteReader)

func SetHtmlRewriter(htmlrw *HtmlRewriter) ReaderOption {
	return func(rr *RewriteReader) {
		rr.SetHtmlRewriter(htmlrw)
	}
}
func AddInsert(i insert) ReaderOption {
	return func(rr *RewriteReader) {
		rr.AddInsert(i.Name, i.Value, !i.NotOnce)
	}
}

type RewriteReader struct {
	tokenizer *html.Tokenizer
	r         io.Reader
	htmlrw    *HtmlRewriter
	stack     *Stack
	inserts   []insert
}

type insert struct {
	Name    string
	Value   string
	NotOnce bool
	matched bool
}

// NewReader 接受一个reader返回TokenReader
// in -> Response.Body.(io.ReadCloser)
// out -> HtmlRewriter.(io.Reader)
func NewRewriteReader(r io.Reader, opts ...ReaderOption) *RewriteReader {
	tokenizer := html.NewTokenizer(r)
	rr := &RewriteReader{
		tokenizer: tokenizer,
		stack:     NewStack(),
	}
	// handler opts ....
	for _, opt := range opts {
		opt(rr)
	}
	return rr
}

func (r *RewriteReader) SetHtmlRewriter(hrw *HtmlRewriter) *RewriteReader {
	r.htmlrw = hrw
	return r
}

func (r *RewriteReader) AddInsert(name, value string, once bool) *RewriteReader {
	r.inserts = append(r.inserts, insert{Name: name, Value: value, NotOnce: !once})
	return r
}

//Read implmement io.Reader interface , that can support stream
func (r *RewriteReader) Read(p []byte) (n int, err error) {
	tokenizer := r.tokenizer
	tt := tokenizer.Next()
	if tt == html.ErrorToken {
		return 0, tokenizer.Err()
	}
	var raw []byte
	switch tt {
	// case html.ErrorToken:
	// 	return 0, tokenizer.Err()
	case html.StartTagToken, html.SelfClosingTagToken:
		name, hasAttr := tokenizer.TagName()
		token := html.Token{
			Type: tt,
			Data: string(name),
		}
		if hasAttr {
			r.htmlrw.rewriteToken(&token, tokenizer)
		}
		r.waitTagTokenClose(token, "script")
		raw = r.processInsert(token)
		// raw = r.insert(token, "head", "<base href='http://www.baidu.com'></base>")

	case html.EndTagToken:
		token := tokenizer.Token()
		raw = r.waitTagTokenClose(token, "script")

	default:
		token := tokenizer.Token()
		raw = r.waitTagTokenClose(token, "script")

	}

	n = copy(p, raw)
	return n, nil
}

func (r *RewriteReader) isScript(tz html.Token) bool {
	n := tz.Data
	return r.isMatchTagName(n, RwTypeCss.String())
	// fmt.Println(string(n), RwTypeScript.String(), bytes.Compare(n, []byte(RwTypeScript.String())))
	// return bytes.Compare(n, []byte(RwTypeScript.String())) == 0
}
func (r *RewriteReader) isMatchTagName(a, b string) bool {
	return a == b
}
func (r *RewriteReader) insert(tz html.Token, name, value string) (raw []byte, matched bool) {
	var bs *bytes.Buffer
	var hasStack bool
	if r.stack.Len() > 0 {
		last := r.stack.Pop().(string)
		bs = bytes.NewBufferString(last)
		hasStack = true
	} else {
		bs = bytes.NewBufferString(tz.String())
	}

	if r.isMatchTagName(name, tz.Data) {
		bs.WriteString(value)
		matched = true
	}

	if hasStack {
		r.stack.Push(bs.String())
		return
	}

	raw = bs.Bytes()
	return
}
func (r *RewriteReader) getRawData(tz html.Token) (raw []byte) {
	if r.stack.Len() > 0 {
		return nil
	}
	return bytes.NewBufferString(tz.String()).Bytes()
}
func (r *RewriteReader) processInsert(tz html.Token) (raw []byte) {
	if len(r.inserts) < 1 {
		return r.getRawData(tz)
	}
	for i, insert := range r.inserts {
		// if !insert.NotOnce && insert.matched {
		// 	continue
		// }
		rr, matched := r.insert(tz, insert.Name, insert.Value)
		if !insert.NotOnce && matched {
			r.inserts[i].matched = matched
			r.inserts = append(r.inserts[:i], r.inserts[i+1:]...)
		}
		raw = append(raw, rr...)
	}
	return
}
func (r *RewriteReader) waitTagTokenClose(tz html.Token, tagName string) (raw []byte) {
	if r.isMatchTagName(tz.Data, tagName) {
		switch tz.Type {
		case html.StartTagToken:
			r.stack.Push(tz.String())
			break
		case html.EndTagToken:
			last, ok := r.stack.Pop().(string)

			var bs *bytes.Buffer
			if ok {
				bs = bytes.NewBufferString(last)
				bs.WriteString(tz.String())

			} else {
				bs = bytes.NewBufferString(tz.String())
			}

			if r.stack.Len() == 0 {
				raw = bs.Bytes()
			} else {
				last, ok := r.stack.Pop().(string)
				if ok {

					r.stack.Push(last + bs.String())
				}
			}
			break
		}
		return

	} else {
		if r.stack.Len() > 0 {
			last, _ := r.stack.Pop().(string)
			bs := bytes.NewBufferString(last)
			// bs.WriteString(tz.String())
			switch tz.Type {
			case html.TextToken:
				bs.WriteString(tz.Data)
			default:
				bs.WriteString(tz.String())

			}
			r.stack.Push(bs.String())
		}
	}

	if r.stack.Len() > 0 {
		return
	}
	switch tz.Type {
	case html.TextToken:
		raw = bytes.NewBufferString(tz.Data).Bytes()
	default:
		raw = bytes.NewBufferString(tz.String()).Bytes()

	}
	return
}
