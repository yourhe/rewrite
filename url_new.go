package rewrite

import (
	"bytes"
	"io"
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
	if bytes.Index(p, []byte("g.alicdn.com")) > -1 ||
		bytes.Index(p, []byte("www.recaptcha.net")) > -1 {
		return p
	}

	originHost := []byte(urw.host)
	originProtocol := []byte(urw.protocol)
	baseURI := []byte(urw.baseURI)

	switch {
	case urw.mode == 1:
		prr := newPortToPathReduceRewriteReader(originHost, originProtocol, baseURI, p)
		bf := bytes.NewBuffer(nil)
		_, err := io.Copy(bf, prr)

		if err != nil {
			urw.err = err
		}
		return bf.Bytes()
	default:
		drr := newPortToDomainRewriteReader(originHost, originProtocol, baseURI, p)
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

type BytesName []byte

var UnRewriteTable = []BytesName{
	BytesName("javascript:"),
	BytesName("data:"),
	BytesName("#"),
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

func recodePath(path []byte) []byte {
	if path[0] != '/' {
		path = append([]byte{'/'}, path...)
	}
	segments := bytes.Split(path, []byte{'/'})
	for i := 0; i < len(segments); i++ {
		v := segments[i]
		// fmt.Println(i, string(v))

		if bytes.Compare(v, []byte{'.', '.'}) == 0 {
			back := i - 1
			if back < 1 {
				back = 1
			}
			// fmt.Println(string(bytes.Join(segments[:back], []byte{'/'})))
			// fmt.Println(string(bytes.Join(segments[i+1:], []byte{'/'})))
			segments = append(segments[:back], segments[i+1:]...)
			i = back - 1
			// fmt.Println(string(bytes.Join(segments, []byte{'/'})))
		}

	}

	// return append([]byte{'/'}, bytes.Join(segments, []byte{'/'})...)
	return bytes.Join(segments, []byte{'/'})
}
