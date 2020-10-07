package rewrite

import (
	"bytes"
	"io"
)

type CssRewriter struct {
	Rw    *UrlRewriter
	r     io.Reader
	buf   *bytes.Buffer
	urlrw *URLRewriter
}

func NewCssRewriter(urlrw *UrlRewriter) *CssRewriter {
	return &CssRewriter{
		Rw:  urlrw,
		buf: bytes.NewBuffer(nil),
	}
}

func (rerw *CssRewriter) Rewrite(p []byte) []byte {
	return ReplaceAllSubmatchFunc(CssUrlRegex, p, func(i []byte) []byte {
		o := rerw.Rw.Rewrite(i)
		return append([]byte("url(\""), append(o, []byte("\")")...)...)
	})
}
func NewNewCssRewriterReader(r io.Reader, urlrw *URLRewriter, opts ...ReaderOption) *CssRewriter {
	return &CssRewriter{
		r:     r,
		buf:   bytes.NewBuffer(nil),
		urlrw: urlrw,
	}
}
func (cssrw *CssRewriter) NewReader(r io.Reader) io.Reader {
	cssrw.r = r
	cssrw.buf = bytes.NewBuffer(nil)
	return cssrw
}
func (cssrw *CssRewriter) Read(p []byte) (int, error) {
	bf := bytes.NewBuffer(nil)
	bf.ReadFrom(cssrw.r)
	// tmp := bytes.ReplaceAll(bf.Bytes(), []byte("url(../"), []byte("url(/../"))
	tmp := bf.Bytes()
	tmp2 := bytes.NewBuffer(nil)
	var end int = 0
	for {
		u, s, i, _ := tq(tmp[end:])
		if i == -1 {
			tmp2.Write(tmp[end:])
			break
		}
		tmp2.Write(tmp[end : s+end])
		if cssrw.urlrw != nil && len(u) > 0 &&
			(u[0] == '/' || u[0] == '.') {
			tmp2.Write(cssrw.urlrw.Rewrite(u))
		} else {
			tmp2.Write(u)
		}
		end = i + end
	}
	cssrw.buf.Write(tmp2.Bytes())
	// bufio.NewReader(cssrw.)

	return cssrw.read(p)
}

func (cssrw *CssRewriter) read(p []byte) (n int, err error) {
	if cssrw.buf != nil {
		return cssrw.buf.Read(p)
	}
	return 0, io.EOF
}

func tq(s []byte) ([]byte, int, int, error) {
	startToken := "url("
	closeToken := ")"
	startTokenLen := len(startToken)
	i := bytes.Index(s, []byte(startToken))
	if i > -1 {
		end := bytes.Index(s[i+4:], []byte(closeToken))
		if end > -1 {
			var err error = io.EOF
			if end < len(s) {
				err = nil
			}
			start := i + startTokenLen
			return s[start : start+end], start, start + end, err
		}
	}
	return nil, -1, -1, io.EOF
}
