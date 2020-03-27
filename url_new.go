package rewrite

import (
	"bytes"
	"io"

	"github.com/valyala/fasthttp"
)

type URLRewriter struct {
	baseURI         string
	host            string
	protocol        string
	protocolOnQuery bool
	mode            int
	err             error
}

func NewURLRewriter(baseURI, host, protocol string, protocolOnQuery bool, mode int) *URLRewriter {
	return &URLRewriter{
		baseURI:         baseURI,
		host:            host,
		protocol:        protocol,
		protocolOnQuery: protocolOnQuery,
		mode:            mode,
	}
}

func (urw *URLRewriter) Rewrite(p []byte) (o []byte) {
	p = bytes.TrimPrefix(p, []byte("&#xD;"))
	p = bytes.TrimPrefix(p, []byte("&#xA;"))
	// p = bytes.Trim(p, "&#xA;")
	p = bytes.TrimSpace(p)
	if len(p) == 0 {
		return p
	}

	originHost := []byte(urw.host)
	originProtocol := []byte(urw.protocol)
	baseURI := []byte(urw.baseURI)
	switch {
	case urw.mode == 1:
		prr := newPathReduceRewriteReader(originHost, originProtocol, baseURI, p)
		bf := bytes.NewBuffer(nil)
		_, err := io.Copy(bf, prr)
		if err != nil {
			urw.err = err
		}
		return bf.Bytes()
	default:
		drr := newDomainRewriteReader(originHost, originProtocol, baseURI, p)
		bf := bytes.NewBuffer(nil)
		_, err := io.Copy(bf, drr)
		if err != nil {
			urw.err = err
		}
		return bf.Bytes()
	}
}

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

type BytesName []byte

var UnRewriteTable = []BytesName{
	BytesName("javascript:"),
	BytesName("data:"),
}

func (prr *pathReduceRewriteReader) read(p []byte) (n int, err error) {
	if prr.buf != nil {
		return prr.buf.Read(p)
	}
	return 0, io.EOF
}
func (prr *pathReduceRewriteReader) Read(p []byte) (n int, err error) {
	if len(prr.src) == 0 {
		// return prr.buf.Read(p)
		return 0, io.EOF
		// return copy(p, prr.src), io.EOF

	}
	if prr.buf != nil {
		return prr.read(p)
	}
	for _, n := range UnRewriteTable {
		if len(prr.src) < len(n) {
			continue
		}
		if bytes.Index(bytes.ToLower(prr.src[:len(n)]), n) == 0 {
			// prr.buf.w
			// return prr.buf.Write(p)
			// fmt.Println(string(prr.src), len(p))
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
				return copy(p, prr.src), io.EOF
			}
			// i := bytes.LastIndexByte(u2.PathOriginal(), '/')
			// if i > -1 {
			// 	p := u2.PathOriginal()[:i+1]
			// 	u.SetPathBytes(append(p, u.PathOriginal()...))
			// 	u.SetPathBytes(u.Path())
			// }
		} else {
			u.SetPathBytes(u2.PathOriginal())
		}
		u.SetHostBytes(u2.Host())
		u.SetSchemeBytes(u2.Scheme())
	}

	if !prr.NotSetProtocolOnQuery {
		protocol := FormatProtocol(&u)
		if protocol != "" {
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
	// if bytes.IndexByte(path, '.') > -1 {
	// 	fmt.Println(string(path))
	// 	fmt.Println(string(u.Path()))

	// 	path = u.Path()
	// }
	if len(path) > 0 && path[0] != '/' {
		path = append([]byte{'/'}, path...)

	}

	u.SetPath("/--/" + pathHost + "/_" + string(path))
	u.SetHostBytes(prr.originHost)
	u.SetSchemeBytes(prr.originProtocol)
	prr.buf = bytes.NewBuffer(u.FullURI())
	return prr.read(p)
}

type domainRewriteReader struct {
	src                   []byte
	originHost            []byte
	originProtocol        []byte
	NotSetProtocolOnQuery bool
	baseURI               []byte
}

func newDomainRewriteReader(originHost, originProtocol, baseURI, src []byte) *domainRewriteReader {
	return &domainRewriteReader{
		src:            src,
		originHost:     originHost,
		originProtocol: originProtocol,
		baseURI:        baseURI,
	}
}

func (drw *domainRewriteReader) Read(p []byte) (n int, err error) {
	if len(drw.src) == 0 {
		return 0, io.EOF
	}
	var protocolTag = "__dp="
	var u fasthttp.URI
	u.Parse(nil, drw.src)
	u.DisablePathNormalizing = true
	if len(u.Host()) == 0 {
		return copy(p, drw.src), io.EOF
	}

	if !drw.NotSetProtocolOnQuery {
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

	formatHost := formatHostBytes(u.Host(), []byte(PortPrefix), []byte{'.'})
	formatHost = append(formatHost, '.')
	u.SetHostBytes(append(formatHost, drw.originHost...))
	u.SetSchemeBytes(drw.originProtocol)
	uri := u.FullURI()
	return copy(p, uri), io.EOF
}

func formatHostBytes(host, portprefix, sep []byte) []byte {
	func() {
		for {
			if len(host) == 0 {
				return
			}
			if host[0] == '.' {
				host = host[1:]
				continue
			}
			return
		}
	}()
	hostname, port := releaseHost(host)

	if len(port) > 0 {
		hostname = append(portprefix, '.')
		hostname = append(portprefix, hostname...)
	}
	if bytes.Compare(sep, []byte{'.'}) != 0 {
		bf := bytes.NewBuffer(nil)
		for i := 0; i < len(hostname); i++ {
			if hostname[i] == '.' {
				bf.Write(sep)
				continue
			}
			bf.WriteByte(hostname[i])
		}
		// return bytes.ReplaceAll(hostname, []byte{'.'}, sep)
		return bf.Bytes()
	}

	return hostname
}
