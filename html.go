package rewrite

import (
	"bytes"
	"io"
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
	// rdr := bytes.NewReader(p)

	tokenizer := html.NewTokenizer(rdr)
	// w := &bytes.Buffer{}
	// var token html.Token
	var isScript bool

	for {
		tt := tokenizer.Next()
		// token := tokenizer.Token()
		switch tt {
		// case html.TextToken:
		// case html.CommentToken:
		case html.ErrorToken:

			// ErrorToken means that an error occurred during tokenization.
			// most common is end-of-file (EOF)

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
			isScript = bytes.Compare(name, []byte(RwTypeScript.String())) == 0
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
		// fmt.Println(string(tokenizer.Text()))
		// b := tokenizer.Raw()
		// if len(b) < 4096 {
		// 	ww.Write(b)
		// 	continue
		// }

		// buf := bytes.NewBuffer(b)
		// for {
		// 	b := buf.Next(4096)
		// 	if len(b) == 0 {
		// 		break
		// 	}
		// 	ww.Write(b)
		// }
		// w.WriteString(token.String())

	}

	return nil
}

func (hrw *HtmlRewriter) rewriteMetaRefresh(p []byte, metaRefresh *regexp.Regexp) {

}

func (hrw *HtmlRewriter) rewriteToken(t *html.Token, tok *html.Tokenizer) {
	attrs := hrw.rewriteTags[t.Data]
	for {
		key, val, more := tok.TagAttr()
		repl := attrs[string(bytes.ToLower(key))]
		if repl != nil {
			val = repl.Rewrite(val)
		}

		t.Attr = append(t.Attr, html.Attribute{
			Key: string(key),
			Val: string(val),
		})

		if !more {
			return
		}
	}
	return
}

// func (hrw *HtmlRewriter) rewriteJSToken(t *html.Token, tok *html.Tokenizer) {
// 	if t.Data == "script" {
// 		tok.NextIsNotRawText()
// 	}
// }

func rewriteTags(defmod Rewriter) map[string]map[string]Rewriter {
	oe := PrefixRewriter{Prefix: []byte("oe_")}
	im := PrefixRewriter{Prefix: []byte("im_")}
	// if_ := PrefixRewriter{Prefix: []byte("if_")}
	fr_ := PrefixRewriter{Prefix: []byte("fr_")}
	// js_ := PrefixRewriter{Prefix: []byte("js_")}

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
	return NewRewriteReader(r,
		SetHtmlRewriter(hrw),
	)
}

type ReaderOption func(*RewriteReader)

func SetHtmlRewriter(htmlrw *HtmlRewriter) ReaderOption {
	return func(rr *RewriteReader) {
		rr.SetHtmlRewriter(htmlrw)
	}
}

type RewriteReader struct {
	tokenizer *html.Tokenizer
	r         io.Reader
	htmlrw    *HtmlRewriter
}

// NewReader 接受一个reader返回TokenReader
// in -> Response.Body.(io.ReadCloser)
// out -> HtmlRewriter.(io.Reader)
func NewRewriteReader(r io.Reader, opts ...ReaderOption) *RewriteReader {
	tokenizer := html.NewTokenizer(r)
	rr := &RewriteReader{
		tokenizer: tokenizer,
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

//Read implmement io.Reader interface , that can support stream
func (r *RewriteReader) Read(p []byte) (n int, err error) {
	tokenizer := r.tokenizer
	tt := tokenizer.Next()
	var raw []byte
	switch tt {
	case html.ErrorToken:
		return 0, tokenizer.Err()
	case html.StartTagToken, html.SelfClosingTagToken:
		name, hasAttr := tokenizer.TagName()
		token := html.Token{
			Type: tt,
			Data: string(name),
		}
		if hasAttr {
			r.htmlrw.rewriteToken(&token, tokenizer)
		}
		tkstring := String(token)
		raw = []byte(tkstring)
	default:
		raw = tokenizer.Raw()
	}

	n = copy(p, raw)
	return n, nil
}
