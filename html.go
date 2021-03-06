package rewrite

import (
	"bytes"
	"fmt"
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
	after   bool
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
	// rr.tokenizer = html.NewTokenizerFragment(r, "title")
	return rr
}

func (r *RewriteReader) SetHtmlRewriter(hrw *HtmlRewriter) *RewriteReader {
	r.htmlRewriter = hrw
	return r
}

func readRegexStrEnd(str string) int {
	i := strings.Index(str, "/")
	if i < 1 {
		return i
	}
	if str[i-1] == '\\' {
		n := readRegexStrEnd(str[i+1:])
		if n > -1 {
			return i + n + 1
		}
		return -1
	}
	return i
}

func readAttrName(src string) (name string, regex *regexp.Regexp, rwfn Rewriter, err error) {
	name = src
	i := strings.Index(src, "/")
	if i > -1 { //has regext
		j := readRegexStrEnd(src[i+1:])
		if j > -1 {
			name = src[:i]
			rgStr := src[i+1 : i+1+j]
			regex, err = regexp.Compile(rgStr)

			r := readRegexStrEnd(src[i+1+j+1:])
			if r > -1 { // has replace
				replaceStr := src[i+1+j+1 : i+1+j+1+r]
				rwfn = &RegexRewriter{
					Re: regex,
					Rw: RewriteFunc(func(in []byte) []byte {
						return regex.ReplaceAll(in, []byte(replaceStr))
						return []byte("-" + replaceStr + "-")
					}),
				}
			}
		}

		// j := strings.LastIndex(src, "/")

	}

	// i = strings.Index(src, ":")
	// if i > -1 && src[len(src)-1] == ':' {
	// 	name = src[:i]
	// 	j := strings.LastIndex(src, "/")
	// 	rgStr := src[i+1 : j]
	// 	regex, err = regexp.Compile(rgStr)
	// 	rwfn = &RegexRewriter{
	// 		Re: regex,
	// 		Rw: RewriteFunc(func(in []byte) []byte {
	// 			fmt.Println(string(in))
	// 			fmt.Println(string(in))
	// 			fmt.Println(string(in))
	// 			fmt.Println(string(in))
	// 			return in
	// 		}),
	// 	}
	// 	return
	// }
	return
}

func (r *RewriteReader) SetTagRewriter(query, attrName string, rw Rewriter) *RewriteReader {
	if query == "" || attrName == "" {
		return r
	}
	attrName, regex, rwfn, err := readAttrName(attrName)
	if err != nil {
		fmt.Println(err, "SetTagRewriter")
		return r
	}
	if rwfn != nil {
		rw = rwfn
	}
	if r.tagTable == nil {
		r.tagTable = map[string][]matcher{}
	}
	// fmt.Println(rw)
	matchs := []match{}
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
	matchs = append(matchs, match{key: attrName, rewrite: rw, regex: regex})

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

func (r *RewriteReader) AddInsertAfter(name, value string, once bool) *RewriteReader {
	r.inserts = append(r.inserts, insert{
		Name:    name,
		Value:   value,
		NotOnce: !once,
		after:   true,
	})
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

type HTMLTokenizer struct {
	*html.Tokenizer
}

// func TagName(tt html.TokenType,z *html.Tokenizer) (name []byte, hasAttr bool) {
// 	if z.data.start < z.data.end {
// 		switch tt {
// 		case html.StartTagToken, html.EndTagToken, html.SelfClosingTagToken:
// 			s := z.buf[z.data.start:z.data.end]
// 			z.data.start = z.raw.end
// 			z.data.end = z.raw.end
// 			return lower(s), z.nAttrReturned < len(z.attr)
// 		}
// 	}
// 	return nil, false
// }
func Token(tt html.TokenType, z *html.Tokenizer) html.Token {
	t := html.Token{}
	switch tt {
	case html.TextToken, html.CommentToken, html.DoctypeToken:
		t.Data = string(z.Text())
	case html.StartTagToken, html.SelfClosingTagToken, html.EndTagToken:
		name, moreAttr := z.TagName()
		for moreAttr {
			var key, val []byte
			key, val, moreAttr = z.TagAttr()
			t.Attr = append(t.Attr, html.Attribute{"", atom.String(key), string(val)})
		}
		if a := atom.Lookup(name); a != 0 {
			t.DataAtom, t.Data = a, a.String()
		} else {
			// fmt.Println(rawName)
			t.DataAtom, t.Data = 0, string(name)
		}
	}
	return t
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
func isBaseToken(token *html.Token) bool {
	return token.DataAtom == atom.Base
}

func readRawTagName(raw []byte) string {
	// fmt.Println(string(raw))
	if raw[0] != '<' {
		return string(raw)
	}
	var tmp = []byte{}
	for i := 1; i < len(raw); i++ {
		switch raw[i] {
		case ' ', '\n', '\r', '\t', '\f', '>':
			return string(tmp)
		case '/':
			break
		default:
			tmp = append(tmp, raw[i])
		}
	}
	return string(tmp)
}

//Read implmement io.Reader interface , that can support stream
func (r *RewriteReader) Read(p []byte) (n int, err error) {

loop:
	tokenizer := r.tokenizer
	tt := tokenizer.Next()
	// fmt.Println(tt.String())
	var raw []byte
	switch tt {
	case html.ErrorToken:
		r.stackToBuf()
		r.eof = true
		return r.read(p)
		// return 0, io.EOF
	case html.StartTagToken, html.SelfClosingTagToken:
		rawTagname := readRawTagName(tokenizer.Raw())
		var token html.Token = tokenizer.Token()

		token.Data = rawTagname

		//process base token
		// 如果在html中发现了base标签并且有href attr。
		// 替换url rewrite 中的baseURI
		if isBaseToken(&token) && hasAttr(&token) {
			for _, attr := range token.Attr {

				if attr.Key == "href" {
					// fmt.Println("token", attr, strings.Index(attr.Val, "//"))

					if strings.Index(attr.Val, "//") == -1 {
						// fmt.Println("////")
						goto loop
					}
					baseURI := attr.Val
					if r.htmlRewriter != nil && r.htmlRewriter.urlrw != nil {
						_, ok := r.htmlRewriter.urlrw.(*URLRewriter)
						if ok {
							// fmt.Println("try replace baseURI", baseURI)
							// fmt.Println("try replace baseURI", baseURI)
							// fmt.Println("try replace baseURI", baseURI)
							r.htmlRewriter.urlrw.(*URLRewriter).baseURI = baseURI

						}

					}

					for _, val := range r.tagTable {
						for _, val := range val {
							for _, val := range val.matchs {
								_, ok := val.rewrite.(*URLRewriter)
								if ok {
									if val.rewrite.(*URLRewriter).baseURI != baseURI {
										val.rewrite.(*URLRewriter).baseURI = baseURI
									}
								}
							}
						}
					}
					goto out
				}

			}
		}
	out:
		if hasAttr(&token) {
			r.rewriteToken(&token, tokenizer)
		}
		if tt == html.SelfClosingTagToken {
			tokenizer.NextIsNotRawText()
			// raw = r.waitTagTokenClose(token, "script")
			// raw = append(raw, r.ProcessInsert(token, tokenizer)...)
		} else {
			// fmt.Println("close", token)
			// raw = append(raw, r.ProcessInsert(token, tokenizer)...)
		}
		raw = r.waitTagTokenClose(token, "script")
		raw = append(raw, r.ProcessInsert(token, tokenizer)...)
		r.buf.Write(raw)
		// fmt.Println("****", string(r.buf.String()))
		// raw = r.insert(token, "head", "<base href='http://www.baidu.com'></base>")

	case html.EndTagToken:
		rawTagname := readRawTagName(tokenizer.Raw())
		token := tokenizer.Token()
		// fmt.Println(token.Data)
		token.Data = rawTagname
		raw = r.ProcessInsertAfter(token, tokenizer)

		raw = append(raw, r.waitTagTokenClose(token, "script")...)
		r.buf.Write(raw)
	// case html.TextToken:
	// 	raw = tokenizer.Raw()
	// 	r.buf.Write(raw)
	case html.CommentToken, html.DoctypeToken:
		r.buf.Write(tokenizer.Raw())
	default:
		raw = r.waitTagTokenzerClose(tokenizer, "script")
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

func (r *RewriteReader) queryTags2(token *html.Token) []Tags {
	tag := r.tagTable[token.Data]
	if tag == nil {
		return nil
	}
	result := []Tags{}
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

				var rewriter *match
				var idx int
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
					result = append(result, Tags{
						match: rewriter,
						index: idx,
					})
				}
			}

		}
	}
	return result
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

type Tags struct {
	match *match
	index int
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

	// match, i := r.queryTags(t)
	matchs := r.queryTags2(t)
	for _, item := range matchs {
		match := item.match
		i := item.index
		if match != nil && i > -1 {
			repl := match.rewrite
			// fmt.Println(t)
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

			if bytes.Compare(newVal, bytes.TrimSpace(val)) != 0 {
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

func (r *RewriteReader) isCloseType(tp html.TokenType) bool {
	return tp == html.EndTagToken
}
func (r *RewriteReader) insert(tz html.Token, name, value string, after bool) (raw []byte, matched bool) {

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

func (r *RewriteReader) insert2(tz html.Token, name, value string, after bool) (raw []byte, matched bool) {

	var bs *bytes.Buffer = bytes.NewBufferString("")
	var hasStack bool
	if r.stack.Len() > 0 {
		last := r.stack.Pop().(string)
		bs = bytes.NewBufferString(last)
		hasStack = true
	} else {
		if !after {
			// bs = bytes.NewBufferString(tz.String())
		}
	}

	if r.isMatchTagName(name, tz.Data) && r.isCloseType(tz.Type) == after {
		bs.WriteString(value)
		matched = true
	}
	if after {
		// bs.WriteString(tz.String())
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
func (r *RewriteReader) ProcessInsert(tz html.Token, tokenizer *html.Tokenizer) (raw []byte) {
	return r.processInsert(tz, tokenizer, false)
}
func (r *RewriteReader) ProcessInsertAfter(tz html.Token, tokenizer *html.Tokenizer) (raw []byte) {
	return r.processInsert(tz, tokenizer, true)
}
func (r *RewriteReader) processInsert(tz html.Token, tokenizer *html.Tokenizer, after bool) (raw []byte) {
	// tz := tokenizer.Token()
	if len(r.inserts) < 1 {
		return nil //r.getRawDataFromToken(tz)
	}
	var insertRaw []byte

	for _, insert := range r.inserts {

		// if !insert.NotOnce && insert.matched {
		// 	continue
		// }
		var rr []byte
		var matched bool
		rr, matched = r.insert2(tz, insert.Name, insert.Value, insert.after)
		// 如果没有发现head标签，在满足当前标签条件下insert
		if !matched && !insert.after &&
			insert.Name == atom.Head.String() &&
			(tz.DataAtom == atom.Link ||
				tz.DataAtom == atom.Script ||
				// tz.DataAtom == atom.Style ||
				tz.DataAtom == atom.Meta) {
			matched = true
			insertRaw = append(insertRaw, []byte(insert.Value)...)
		}

		if !insert.NotOnce && matched {
			// r.inserts[i].matched = matched
			r.inserts = append(r.inserts[:0], r.inserts[1:]...)
		}

		insertRaw = append(insertRaw, rr...)
	}
	if after {
		return insertRaw //append(insertRaw, r.getRawDataFromToken(tz)...)
	}

	return insertRaw //append(r.getRawDataFromToken(tz), insertRaw...)
}

func (r *RewriteReader) waitTagTokenzerClose(tokenizer *html.Tokenizer, tagName string) (raw []byte) {
	rawTagname := readRawTagName(tokenizer.Raw())
	tz := tokenizer.Token()
	tz.Data = rawTagname
	return r.waitTagTokenClose(tz, tagName)
}
func (r *RewriteReader) waitTagTokenClose(tz html.Token, tagName string) (raw []byte) {
	if r.isMatchTagName(tz.Data, tagName) {
		switch tz.Type {
		case html.StartTagToken:
			r.stack.Push(String(tz))
			break
		case html.EndTagToken, html.SelfClosingTagToken:
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
					data := strings.Replace(tz.Data, "<!--", "//<!--", -1)
					jsBuf := bytes.NewReader([]byte(data))
					// fmt.Println(tz.Data)

					r.jsRewriter.NewReader(jsBuf)
					io.Copy(bs, bytes.NewBufferString("(function(window){"))
					io.Copy(bs, r.jsRewriter)
					io.Copy(bs, bytes.NewBufferString("})(window)"))
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
