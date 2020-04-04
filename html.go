package rewrite

import (
	"bytes"
	"io"
	"strings"

	// "io/ioutil"

	"regexp"
	"strconv"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
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
	transfrom     *Transform
	// rewriters     map[RewriterType]Rewriter
}

func NewHtmlRewriter(urlrw Rewriter, configs ...func(*Config)) *HtmlRewriter {
	c := makeConfig(configs...)
	return &HtmlRewriter{
		urlrw: urlrw,
		// jsRewriter:  c.Rewriters[RwTypeJavascript],
		rewriteTags: rewriteTags(urlrw),
		transfrom:   c.Transform,
	}
}

func NewHtmlJSRewriter(urlrw Rewriter, jsrw Rewriter, configs ...func(*Config)) *HtmlRewriter {
	c := makeConfig(configs...)
	return &HtmlRewriter{
		urlrw:       urlrw,
		jsRewriter:  jsrw,
		rewriteTags: rewriteTags(urlrw),
		transfrom:   c.Transform,
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
		ww.Write(tokenizer.Raw())
	}
	return nil
}

func (hrw *HtmlRewriter) rewriteMetaRefresh(p []byte, metaRefresh *regexp.Regexp) {

}

func (hrw *HtmlRewriter) rewriteToken(t *html.Token, tok *html.Tokenizer) {
	if hrw.rewriteTags == nil {
		return
	}
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
		buf.WriteString(html.EscapeString(a.Val))
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

func SetTransform(transform *Transform) ReaderOption {
	return func(r *RewriteReader) {
		r.transform = transform
	}
}
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

type TagsRewriter map[string][]matcher

type match struct {
	key     string
	value   *string
	regex   *regexp.Regexp
	rewrite Rewriter
}

type matcher struct {
	matchs []match
}

type AttrRewriter map[string]matcher

type AttrsRewriter struct {
	AttrRewriter
}

type RewriteReader struct {
	tokenizer    *html.Tokenizer
	r            io.Reader
	htmlRewriter *HtmlRewriter
	stack        *Stack
	inserts      []insert
	tagTable     TagsRewriter
	buf          *bytes.Buffer
	jsRewriter   *JavaScriptRewrite
	transform    *Transform
	eof          bool
	outreader    io.Reader
}

type insert struct {
	Name    string
	Value   string
	NotOnce bool
	matched bool
}

// NewRewriteReader 接受一个reader返回TokenReader
// in -> Response.Body.(io.ReadCloser)
// out -> HtmlRewriter.(io.Reader)
func NewRewriteReader(r io.Reader, opts ...ReaderOption) *RewriteReader {
	// var tokenizer *html.Tokenizer

	rr := &RewriteReader{
		// tokenizer: tokenizer,
		stack:    NewStack(),
		tagTable: make(TagsRewriter),
		buf:      bytes.NewBuffer(nil),
	}
	// handler opts ....
	for _, opt := range opts {
		opt(rr)
	}
	// if r != nil {
	// 	if rr.transform != nil {
	// 		rr.tokenizer = html.NewTokenizer(rr.transform.NewReader(r))
	// 	} else {
	// 		rr.tokenizer = html.NewTokenizer(r)

	// 	}
	// }
	rr.NewReader(r)
	return rr
}

func (rr *RewriteReader) SetJavascriptRewriter(jsRewriter *JavaScriptRewrite) io.Reader {
	rr.jsRewriter = jsRewriter

	return rr
}
func (rr *RewriteReader) Reset() {
	rr.buf = bytes.NewBuffer(nil)
}
func (rr *RewriteReader) NewReader(r io.Reader) io.Reader {
	if r == nil {
		return rr
	}
	rr.Reset()

	if rr.transform != nil {
		r = rr.transform.NewDecodeReader(r)
		rr.outreader = rr.transform.NewEncodeReader(rr.buf)
	}

	rr.tokenizer = html.NewTokenizer(r)
	return rr
}

func (r *RewriteReader) SetHtmlRewriter(hrw *HtmlRewriter) *RewriteReader {
	r.htmlRewriter = hrw
	return r
}

func readAttrName(src string) (name string, regex *regexp.Regexp) {
	name = src
	i := strings.Index(src, "/")
	if i > -1 {
		name = src[:i]
		j := strings.LastIndex(src, "/")
		rgStr := src[i+1 : j]
		regex = regexp.MustCompile(rgStr)
	}
	return
}

func (r *RewriteReader) SetTagRewriter(query, attrName string, rw Rewriter) *RewriteReader {
	if query == "" || attrName == "" || rw == nil {
		return r
	}
	attrName, regex := readAttrName(attrName)
	if r.tagTable == nil {
		r.tagTable = map[string][]matcher{}
	}

	matchs := []match{match{key: attrName, rewrite: rw, regex: regex}}
	var tagName = query
	i := strings.Index(query, "[")
	if i > -1 {
		j := strings.LastIndex(query, "]")
		if j > -1 {
			tagName = query[:i]
			expr := query[i+1 : j]
			parts := strings.Split(expr, "=")
			key := strings.TrimSpace(parts[0])
			var val *string
			if len(parts) > 1 {
				v := strings.TrimSpace(parts[1])
				if v[0] == '"' || v[0] == '\'' {
					v = v[1 : len(v)-1]
				}
				val = &v
			}
			matchs = append(matchs, match{
				key:   key,
				value: val,
			})
		}

	}

	r.tagTable[tagName] = append(r.tagTable[tagName], matcher{
		matchs: matchs,
	})

	// r.tagTable[tagName] = &AttrsRewriter{
	// 	AttrRewriter: AttrRewriter{
	// 		attrName: matcher{key: attrName, rewrite: rw},
	// 	},
	// }
	// r.tagTable[tagName].AttrRewriter
	return r
}
func (r *RewriteReader) AddInsert(name, value string, once bool) *RewriteReader {
	r.inserts = append(r.inserts, insert{Name: name, Value: value, NotOnce: !once})
	return r
}
func (r *RewriteReader) stackToBuf() {

	for {
		s := r.stack.Pop()
		if s == nil {
			break
		}

		if r.buf == nil {
			r.buf = bytes.NewBufferString(s.(string))
		} else {
			r.buf.WriteString(s.(string))
		}
	}

}
func (r *RewriteReader) read(p []byte) (n int, err error) {
	if r.eof && r.buf.Len() == 0 {
		return 0, io.EOF
	}
	if r.outreader != nil {
		return r.outreader.Read(p)
	}

	// if r.transform != nil {

	// 	lr := io.LimitReader(r.buf, 4096)
	// 	er := r.transform.NewEncodeReader(lr)

	// 	i := 0
	// 	for {
	// 		n, err := er.Read(p[i:])
	// 		i += n
	// 		if err != nil || i >= len(p) {
	// 			break
	// 		}

	// 	}
	// 	return i, nil

	// }
	n, err = r.buf.Read(p)
	// if r.buf.Len() != 0 {
	// 	fmt.Printf("buf:%d, read:%d\n", r.buf.Len(), n)

	// }
	return n, nil
}

func hasAttr(token *html.Token) bool {
	return len(token.Attr) > 0
}

//Read implmement io.Reader interface , that can support stream
func (r *RewriteReader) Read(p []byte) (n int, err error) {

loop:
	tokenizer := r.tokenizer
	tt := tokenizer.Next()
	var raw []byte
	switch tt {
	case html.ErrorToken:
		r.stackToBuf()
		r.eof = true
		return r.read(p)
		// return 0, io.EOF
	case html.StartTagToken, html.SelfClosingTagToken:

		var token html.Token = tokenizer.Token()

		if hasAttr(&token) {
			r.rewriteToken(&token, tokenizer)
		}

		r.waitTagTokenClose(token, "script")
		raw = r.processInsert(token, tokenizer)
		r.buf.Write(raw)

		// raw = r.insert(token, "head", "<base href='http://www.baidu.com'></base>")

	case html.EndTagToken:
		token := tokenizer.Token()
		raw = r.waitTagTokenClose(token, "script")
		r.buf.Write(raw)

	default:

		token := tokenizer.Token()
		raw = r.waitTagTokenClose(token, "script")
		r.buf.Write(raw)

	}
	if len(raw) == 0 {
		goto loop
	}

	return r.read(p)
	// n = copy(p, raw)
	// return n, nil
}

func (r *RewriteReader) getTagRewriter(tagName string, attrName []byte) Rewriter {
	if tagName == "" || len(attrName) == 0 {
		return nil
	}
	attrName = (bytes.ToLower(attrName))
	tag := r.tagTable[tagName]

	if tag == nil {
		return nil
	}
	return nil
	// return tag[string(attrName)]

}

func (r *RewriteReader) queryTags(token *html.Token) (rewriter *match, idx int) {
	tag := r.tagTable[token.Data]
	if tag == nil {
		return nil, -1
	}

	for _, tagMatchs := range tag {
		var matched = make([]bool, len(tagMatchs.matchs))
		for i, matcher := range tagMatchs.matchs {
			for j, attr := range token.Attr {

				if attr.Key != matcher.key {
					continue
				}

				if matcher.value != nil {
					if attr.Val != *matcher.value {
						continue
					}
				}
				if matcher.rewrite != nil {

					rewriter = &tagMatchs.matchs[i]

					idx = j
				}
				matched[i] = true
				var e = true
				for _, b := range matched {
					e = e && b
				}
				if e {
					return
				}
			}

		}
	}
	return nil, -1
}

func (r *RewriteReader) rewriteToken(t *html.Token, tok *html.Tokenizer) {
	// attrs := hrw.rewriteTags[t.Data]
	// for {

	// 	key, val, more := tok.TagAttr()
	// 	if len(key) > 0 {

	// 		t.Attr = append(t.Attr, html.Attribute{
	// 			Key: string(key),
	// 			Val: string(val),
	// 		})
	// 	}
	// 	if !more {
	// 		break
	// 	}
	// }

	match, i := r.queryTags(t)

	if match != nil && i > -1 {
		repl := match.rewrite
		if repl == nil {
			return
		}
		val := []byte(t.Attr[i].Val)
		newVal := val
		if match.regex != nil {
			findpart := match.regex.Copy().FindSubmatchIndex(val)
			if len(findpart) > 0 {
				i := findpart[2]
				j := findpart[3]
				bf := bytes.NewBuffer(val[:i])
				bf.Write(repl.Rewrite(val[i:j]))
				bf.Write(val[j:])
				newVal = bf.Bytes()

			}

		} else {
			newVal = repl.Rewrite(val)
		}

		if bytes.Compare(newVal, val) != 0 {
			t.Attr[i].Val = string(newVal)
			t.Attr = append(t.Attr, html.Attribute{
				Key: "__cpp",
				Val: "1",
			})
			t.Attr = append(t.Attr, html.Attribute{
				Key: "__dp",
				Val: "1",
			})
		}
	}
	return
	for {
		key, oldval, more := tok.TagAttr()
		var newVal []byte
		repl := r.getTagRewriter(t.Data, key)
		// repl := attrs[string(bytes.ToLower(key))]
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
func (r *RewriteReader) isScript(tz html.Token) bool {
	n := tz.Data
	return r.isMatchTagName(n, RwTypeCss.String())
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

func (r *RewriteReader) getRawData(tz *html.Tokenizer) (raw []byte) {
	if r.stack.Len() > 0 {
		return nil
	}
	return tz.Raw() // bytes.NewBufferString(String(tz)).Bytes()
	// return bytes.NewBufferString(tz.String()).Bytes()
}
func (r *RewriteReader) getRawDataFromToken(tk html.Token) (raw []byte) {
	if r.stack.Len() > 0 {
		return nil
	}
	return bytes.NewBufferString(String(tk)).Bytes()
	// return bytes.NewBufferString(tz.String()).Bytes()
}
func (r *RewriteReader) processInsert(tz html.Token, tokenizer *html.Tokenizer) (raw []byte) {
	// tz := tokenizer.Token()
	if len(r.inserts) < 1 {
		return r.getRawDataFromToken(tz)
	}
	for i, insert := range r.inserts {
		// if !insert.NotOnce && insert.matched {
		// 	continue
		// }
		rr, matched := r.insert(tz, insert.Name, insert.Value)
		// 如果没有发现head标签，在满足当前标签条件下insert
		if !matched &&
			insert.Name == atom.Head.String() &&
			(tz.DataAtom == atom.Link ||
				tz.DataAtom == atom.Script ||
				tz.DataAtom == atom.Style ||
				tz.DataAtom == atom.Meta) {
			matched = true
			raw = []byte(insert.Value)
		}

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
			r.stack.Push(String(tz))
			break
		case html.EndTagToken:
			last, ok := r.stack.Pop().(string)

			var bs *bytes.Buffer
			if ok {
				bs = bytes.NewBufferString(last)
				bs.WriteString(String(tz))

			} else {
				bs = bytes.NewBufferString(String(tz))
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
				if tagName == "script" && r.jsRewriter != nil {
					jsBuf := bytes.NewReader([]byte(tz.Data))
					r.jsRewriter.NewReader(jsBuf)
					io.Copy(bs, r.jsRewriter)
				} else {
					bs.WriteString(tz.Data)
				}
			default:
				bs.WriteString(String(tz))
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
		raw = bytes.NewBufferString(String(tz)).Bytes()

	}
	return
}
