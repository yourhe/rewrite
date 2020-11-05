package rewrite

import (
	"bytes"
	"io"

	"github.com/valyala/fasthttp"
)

type pathReduceRewriteReader struct {
	src                   []byte
	originHost            []byte
	originProtocol        []byte
	NotSetProtocolOnQuery bool
	baseURI               []byte
	buf                   *bytes.Buffer
}

func newPathReduceRewriteReader(originHost, originProtocol, baseURI, src []byte) *pathReduceRewriteReader {
	return &pathReduceRewriteReader{
		src:            src,
		originHost:     originHost,
		originProtocol: originProtocol,
		baseURI:        baseURI,
		// buf:            bytes.NewBuffer(src),
	}
}

func (prr *pathReduceRewriteReader) read(p []byte) (n int, err error) {
	if prr.buf != nil {
		return prr.buf.Read(p)
	}
	return 0, io.EOF
}
func (prr *pathReduceRewriteReader) Read(p []byte) (n int, err error) {
	if len(prr.src) == 0 {
		return 0, io.EOF
	}
	if prr.buf != nil {
		return prr.read(p)
	}

	for _, n := range UnRewriteTable {
		if len(prr.src) < len(n) {
			continue
		}
		if bytes.Index(bytes.ToLower(prr.src[:len(n)]), n) == 0 {

			if prr.buf == nil {
				prr.buf = bytes.NewBuffer(prr.src)
			}

			return prr.read(p)
		}
	}

	var protocolTag = "__dp="

	var u fasthttp.URI
	u.Parse(nil, prr.src)

	u.DisablePathNormalizing = true
	if len(u.Host()) == 0 {

		var u2 fasthttp.URI
		u2.Parse(nil, prr.baseURI)
		// 处理相对路径 https://www/a/b/c + ../abc/html => https:/www/a/abc/html
		if len(u.PathOriginal()) > 0 {
			switch u.PathOriginal()[0] {
			case '/':
			default:
				if prr.buf == nil {
					prr.buf = bytes.NewBuffer(prr.src)
				}

				return prr.read(p)
			}

		} else {
			u.SetPathBytes(u2.PathOriginal())
		}
		u.SetHostBytes(u2.Host())
		u.SetSchemeBytes(u2.Scheme())
	}

	if !prr.NotSetProtocolOnQuery {
		protocol := FormatProtocol(&u)
		if protocol != "" && protocol != "http" {
			query := u.QueryString()
			if len(query) == 0 {
				query = append(query[:0], []byte(protocolTag+protocol)...)
			} else {
				query = append(query, []byte("&"+protocolTag+protocol)...)
			}
			u.SetQueryStringBytes(query)
		}

		//remove origin port
		// like: www.ex.com:81 => www.ex.com
		if i := bytes.IndexByte(u.Host(), ':'); i > -1 {
			u.SetHostBytes(u.Host()[:i])
		}
	}

	pathHost := ToSWReduceLetDigHyp(string(u.Host()), string(u.Scheme()))
	path := u.PathOriginal()

	if len(path) > 0 && path[0] != '/' {
		path = append([]byte{'/'}, path...)

	}

	u.SetPath("/--/" + pathHost + "/_" + string(path))
	u.SetHostBytes(prr.originHost)
	u.SetSchemeBytes(prr.originProtocol)
	prr.buf = bytes.NewBuffer(u.FullURI())
	return prr.read(p)
}

type pathPortToReduceRewriteReader struct {
	src                   []byte
	originHost            []byte
	originProtocol        []byte
	NotSetProtocolOnQuery bool
	baseURI               []byte
	buf                   *bytes.Buffer
}

func newPortToPathReduceRewriteReader(originHost, originProtocol, baseURI, src []byte) *pathPortToReduceRewriteReader {
	return &pathPortToReduceRewriteReader{
		src:            src,
		originHost:     originHost,
		originProtocol: originProtocol,
		baseURI:        baseURI,
		// buf:            bytes.NewBuffer(src),
	}
}

func (prr *pathPortToReduceRewriteReader) read(p []byte) (n int, err error) {
	if prr.buf != nil {
		return prr.buf.Read(p)
	}
	return 0, io.EOF
}
func (prr *pathPortToReduceRewriteReader) Read(p []byte) (n int, err error) {
	if len(prr.src) == 0 {
		return 0, io.EOF
	}
	if prr.buf != nil {
		return prr.read(p)
	}

	for _, n := range UnRewriteTable {
		if len(prr.src) < len(n) {
			continue
		}
		if bytes.Index(bytes.ToLower(prr.src[:len(n)]), n) == 0 {

			if prr.buf == nil {
				prr.buf = bytes.NewBuffer(prr.src)
			}

			return prr.read(p)
		}
	}

	var protocolTag = "__dp="

	var u fasthttp.URI
	u.Parse(nil, prr.src)

	u.DisablePathNormalizing = true
	if len(u.Host()) == 0 {

		var u2 fasthttp.URI
		u2.Parse(nil, prr.baseURI)
		// 处理相对路径 https://www/a/b/c + ../abc/html => https:/www/a/abc/html
		if len(u.PathOriginal()) > 0 {
			switch u.PathOriginal()[0] {
			case '/':
			default:
				// CNKI 某些相对路径的url不能重写。
				// 例如
				// https://navi.cnki.net/KNavi/JournalDetail?pcode=CJFD&pykm=ZRZY&Year=&Issue=
				// https://navi.cnki.net/knavi/JournalDetail/GetArticleList?year=2020&issue=09&pykm=ZRZY&pageIdx=0&pcode=CJFD
				// ！！！！！如需修改请考虑冲突

				prr.buf = bytes.NewBuffer(prr.src)
				return prr.read(p)
				absolutePath := append(u2.PathOriginal(), '/', '.', '.', '/')
				absolutePath = append(absolutePath, u.PathOriginal()...)
				u.SetPathBytes(recodePath(absolutePath))
			}

		} else {
			u.SetPathBytes(u2.PathOriginal())
		}
		u.SetHostBytes(u2.Host())
		u.SetSchemeBytes(u2.Scheme())
	}
	if !prr.NotSetProtocolOnQuery {
		protocol := FormatProtocol(&u)
		if protocol != "" && protocol != "http" {
			query := u.QueryString()
			if len(query) == 0 {
				query = append(query[:0], []byte(protocolTag+protocol)...)
			} else {
				query = append(query, []byte("&"+protocolTag+protocol)...)
			}
			u.SetQueryStringBytes(query)
		}

		//remove origin port
		// like: www.ex.com:81 => www.ex.com
		// if i := bytes.IndexByte(u.Host(), ':'); i > -1 {
		// 	u.SetHostBytes(u.Host()[:i])
		// }
	}

	pathHost := ToSWReduceLetDigHyp(string(u.Host()), string(u.Scheme()))
	path := u.PathOriginal()

	if len(path) > 0 && path[0] != '/' {
		path = append([]byte{'/'}, path...)

	}

	u.SetPath("/--/" + pathHost + "/_" + string(path))
	u.SetHostBytes(prr.originHost)
	u.SetSchemeBytes(prr.originProtocol)
	prr.buf = bytes.NewBuffer(u.FullURI())
	return prr.read(p)
}
