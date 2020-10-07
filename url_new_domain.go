package rewrite

import (
	"bytes"
	"io"
	"strings"

	"github.com/valyala/fasthttp"
)

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

type portToDomainRewriteReader struct {
	src                   []byte
	originHost            []byte
	originProtocol        []byte
	NotSetProtocolOnQuery bool
	baseURI               []byte
	ProcessRelativePath   bool
	buf                   *bytes.Buffer
}

func newPortToDomainRewriteReader(originHost, originProtocol, baseURI, src []byte) *portToDomainRewriteReader {
	return &portToDomainRewriteReader{
		src:            src,
		originHost:     originHost,
		originProtocol: originProtocol,
		baseURI:        baseURI,
	}
}

func newPortToDomainRewriteReaderProcessRelative(originHost, originProtocol, baseURI, src []byte) *portToDomainRewriteReader {
	return &portToDomainRewriteReader{
		src:                 src,
		originHost:          originHost,
		originProtocol:      originProtocol,
		baseURI:             baseURI,
		ProcessRelativePath: true,
	}
}
func (drw *portToDomainRewriteReader) read(p []byte) (n int, err error) {
	if drw.buf != nil {
		return drw.buf.Read(p)
	}
	return 0, io.EOF
}
func (drw *portToDomainRewriteReader) Read(p []byte) (n int, err error) {
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

	// absoluteTo
	if len(u.Host()) == 0 {
		// fmt.Println("absoluteTo", string(drw.src))
		// fmt.Println(len(u.PathOriginal()) > 0, string(u.PathOriginal()), string(drw.src))
		if len(u.PathOriginal()) > 0 {
			switch u.PathOriginal()[0] {
			case '/':

			default:
				absolutePath := append(baseURI.PathOriginal(), '/', '.', '.', '/')
				absolutePath = append(absolutePath, u.PathOriginal()...)
				u.SetPathBytes(recodePath(absolutePath))

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
			protocol = strings.Replace(protocol, "|", "-port-", -1)
			protocol = strings.Replace(protocol, "http", "h", -1)

			if formatBProtocol(protocol) != "" {
				host := append([]byte(protocol+"."), u.Host()...)
				u.SetHostBytes(host)
			}

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
