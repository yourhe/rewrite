package rewrite

import (
	"bytes"
	"io"

	"github.com/valyala/fasthttp"
)

type URLRewriter struct {
	baseURI             string
	host                string
	protocol            string
	protocolOnQuery     bool
	mode                int
	err                 error
	processRelativePath bool
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
func NewURLRewriterRelativePath(baseURI, host, protocol string, protocolOnQuery bool, mode int) *URLRewriter {
	return &URLRewriter{
		baseURI:             baseURI,
		host:                host,
		protocol:            protocol,
		protocolOnQuery:     protocolOnQuery,
		mode:                mode,
		processRelativePath: true,
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
		drr.ProcessRelativePath = urw.processRelativePath
		drr.NotSetProtocolOnQuery = !urw.protocolOnQuery
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
	BytesName("#"),
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

type domainRewriteReader struct {
	src                   []byte
	originHost            []byte
	originProtocol        []byte
	NotSetProtocolOnQuery bool
	baseURI               []byte
	ProcessRelativePath   bool
	buf                   *bytes.Buffer
}

func newDomainRewriteReader(originHost, originProtocol, baseURI, src []byte) *domainRewriteReader {
	return &domainRewriteReader{
		src:            src,
		originHost:     originHost,
		originProtocol: originProtocol,
		baseURI:        baseURI,
	}
}

func newDomainRewriteReaderProcessRelative(originHost, originProtocol, baseURI, src []byte) *domainRewriteReader {
	return &domainRewriteReader{
		src:                 src,
		originHost:          originHost,
		originProtocol:      originProtocol,
		baseURI:             baseURI,
		ProcessRelativePath: true,
	}
}

func (drw *domainRewriteReader) read(p []byte) (n int, err error) {
	if drw.buf != nil {
		return drw.buf.Read(p)
	}
	return 0, io.EOF
}
func (drw *domainRewriteReader) Read(p []byte) (n int, err error) {
	if len(drw.src) == 0 {
		return 0, io.EOF
	}

	if drw.buf != nil {
		return drw.read(p)
	}

	for _, n := range UnRewriteTable {
		if len(drw.src) < len(n) {
			continue
		}
		if bytes.Index(bytes.ToLower(drw.src[:len(n)]), n) == 0 {
			// return copy(p, drw.src), io.EOF
			if drw.buf == nil {
				drw.buf = bytes.NewBuffer(drw.src)
			}

			return drw.read(p)
		}
	}
	var protocolTag = "__dp="
	var u fasthttp.URI
	var baseURI = fasthttp.AcquireURI()
	defer fasthttp.ReleaseURI(baseURI)
	baseURI.Parse(nil, drw.baseURI)
	baseURI.DisablePathNormalizing = true

	u.Parse(nil, drw.src)

	u.DisablePathNormalizing = true

	if len(drw.src) > 1 && drw.src[0] == '/' && drw.src[1] == '/' {
		// drw.baseURI
		u.SetSchemeBytes(baseURI.Scheme())
	}
	// fmt.Println(string(u.RequestURI()))
	// fmt.Println(string(u.FullURI()), string(drw.src))
	if len(u.Host()) == 0 && !drw.ProcessRelativePath {
		if drw.buf == nil {
			drw.buf = bytes.NewBuffer(drw.src)
		}

		return drw.read(p)
	}

	if len(u.Host()) == 0 {

		if len(u.PathOriginal()) > 0 {
			switch u.PathOriginal()[0] {
			case '/':
			default:
				if drw.buf == nil {
					drw.buf = bytes.NewBuffer(drw.src)
				}
				return drw.read(p)
			}
		} else {
			u.SetPathBytes(baseURI.PathOriginal())
		}
		u.SetHostBytes(baseURI.Host())
		u.SetSchemeBytes(baseURI.Scheme())

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
	drw.buf = bytes.NewBuffer(uri)
	return drw.read(p)
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
