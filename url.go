package rewrite

import (
	"bytes"
	"fmt"
	"net/url"
	"strings"

	"github.com/valyala/fasthttp"
)

const ReplaceHostFlag = "${host}"
const PortPrefix = "port-"

type UrlRewriter struct {
	hostRelative    bool
	fromHostName    string
	pathDepth       int
	to              *url.URL
	portOnQuery     bool
	protocolOnQuery bool
}

func (ur *UrlRewriter) PortOnQuery(on bool) {
	ur.portOnQuery = on
}
func (ur *UrlRewriter) ProtocolOnQuery(on bool) {
	ur.protocolOnQuery = on
}

// func NewUrlRewriter2(from, to string) *UrlRewriter {
// 	return &UrlRewriter{
// 		fromHostName: from,
// 		to:           to,
// 	}
// }
func NewUrlRewriter(from, to string) *UrlRewriter {
	f, err := url.Parse(from)
	if err != nil {
		// TODO - ugh.
		panic(err)
	}
	host := f.Host
	host = EncodeHost(host, PortPrefix, ".")
	host = strings.TrimSpace(host)
	if host[0] == '.' {
		host = host[1:]
	}
	to = strings.Replace(to, ReplaceHostFlag, host, -1)

	t, err := url.Parse(to)
	if err != nil {
		// TODO
		panic(err)
	}

	return &UrlRewriter{
		fromHostName: f.Hostname(),
		to:           t,
	}
}

func NewLetDigHostUrlRewriter(from, to string) *UrlRewriter {
	f, err := url.Parse(from)
	if err != nil {
		// TODO - ugh.
		panic(err)
	}
	host := ToLetDigHyp(f.Host)
	to = strings.Replace(to, ReplaceHostFlag, host, -1)
	t, err := url.Parse(to)
	if err != nil {
		// TODO
		panic(err)
	}
	return &UrlRewriter{
		fromHostName: f.Hostname(),
		to:           t,
	}
}

func NewSWHostUrlRewriter(from, to string) *UrlRewriter {
	f, err := url.Parse(from)
	if err != nil {
		// TODO - ugh.
		panic(err)
	}

	host := ToSWLetDigHyp(f.Host, f.Scheme)
	to = strings.Replace(to, ReplaceHostFlag, host, -1)

	t, err := url.Parse(to)
	if err != nil {
		// TODO
		panic(err)
	}

	return &UrlRewriter{
		fromHostName: f.Hostname(),
		to:           t,
	}
}

func NewSWReduceHostUrlRewriter(from, to string) *UrlRewriter {
	f, err := url.Parse(from)
	if err != nil {
		// TODO - ugh.
		panic(err)
	}

	host := ToSWReduceLetDigHyp(f.Host, f.Scheme)
	to = strings.Replace(to, ReplaceHostFlag, host, -1)
	t, err := url.Parse(to)
	if err != nil {
		// TODO
		panic(err)
	}

	return &UrlRewriter{
		fromHostName: f.Hostname(),
		to:           t,
	}
}

// NewRelativeUrlRewriter turns urls that match from's
// hostname into relative urls
func NewRelativeUrlRewriter(from string) *UrlRewriter {
	f, err := url.Parse(from)
	if err != nil {
		// TODO - ugh.
		panic(err)
	}

	return &UrlRewriter{
		fromHostName: f.Hostname(),
		to:           &url.URL{},
	}
}

func (urw *UrlRewriter) RewriteString(p string) string {
	return string(urw.Rewrite([]byte(p)))
}

func hasSlash(p []byte) bool {
	j := bytes.IndexByte(p, '?')
	if j > -1 {
		return p[j-1] == '/'
	}
	// hasSlash := p[len(p)-1] == '/'
	return p[len(p)-1] == '/'
}

func (urw *UrlRewriter) Rewrite(p []byte) []byte {
	// call to rewrite with empty slice is a no-op
	if len(p) == 0 {
		return nil
	}

	var u fasthttp.URI
	u.Parse(nil, p)

	u.DisablePathNormalizing = true
	if urw.protocolOnQuery {
		protocol := FormatProtocol(&u)
		if protocol != "" && protocol != "http" {
			query := u.QueryString()
			if len(query) == 0 {
				query = append(query[:0], []byte("__dp="+protocol)...)
			} else {
				query = append(query, []byte("&__dp="+protocol)...)
			}
			u.SetQueryStringBytes(query)
		}
		//remove origin port
		if i := bytes.IndexByte(u.Host(), ':'); i > -1 {
			u.SetHostBytes(u.Host()[:i])
		}
	}

	vPath := strings.Replace(urw.to.Path, "${procotol}", string(u.Scheme()), -1)
	if vPath != "" {
		u.SetPath(vPath + string(u.PathOriginal()))

	}

	host := u.Host()

	hostname, _ := releaseHost(host)

	if len(host) == 0 || string(hostname) == urw.fromHostName {
		u.SetHost(urw.to.Host)
		u.SetScheme(urw.to.Scheme)
	}

	// relative urls should be "directory relative"
	if len(u.Host()) == 0 {
		u.SetPath("." + string(u.Path()))
	}

	if urw.hostRelative {
		u.SetScheme("") // = ""
		return append(urw.pathPrefix(), []byte(u.String())[2:]...)
	}
	uri := u.FullURI()
	return uri
	if hasSlash(p) {
		return uri
	}
	fmt.Println(string(uri), ",", string(u.LastPathSegment()))
	if len(u.LastPathSegment()) > 1 {
		return uri
	}
	if j := bytes.IndexByte(uri, '?'); j > -1 {
		return append(uri[:j-1], uri[j:]...)
	}
	return uri[:len(uri)-1]
}

func releaseHost(host []byte) (hostname, port []byte) {
	i := bytes.IndexByte(host, ':')
	if i > -1 {
		return host[:i], host[i+1:]
	}
	return host, nil
}

func FormatProtocol(u *fasthttp.URI) string {
	// port := u.Port()
	var port string
	host := u.Host()
	i := bytes.Index(host, []byte{':'})
	if i > -1 {
		port = "|" + string(host[i+1:])
	}
	protocol := u.Scheme()
	return fmt.Sprintf("%s%s", protocol, port)
}
func NewHostRelativeUrlRewriter(from string) *UrlRewriter {
	f, err := url.Parse(from)
	if err != nil {
		// TODO - ugh.
		panic(err)
	}

	if f.Path == "" {
		f.Path = "/"
	}

	return &UrlRewriter{
		fromHostName: f.Host,
		hostRelative: true,
		pathDepth:    bytes.Count([]byte(f.Path), []byte{'/'}),
		to:           f,
	}
}

func (urw *UrlRewriter) pathPrefix() []byte {
	return bytes.Repeat([]byte("../"), urw.pathDepth)
}
