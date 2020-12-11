package rewrite

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/valyala/fasthttp"
	// proxycontext "gitlab.iyorhe.com/wfgz/reverseproxy/context"
)

type RewriteRule int

const (
	Keep RewriteRule = iota
	PrefixIfUrlRewrite
	Prefix
	UrlRewrite
	PrefixIfContentRewrite
	ContentLength
	Cookie
	Subfix
	StrictTransportSecurity
	AccessControlAllowOrigin
	XFrameOptonis
	Refresh
)

type HeaderRewriter struct {
	Prefix                    string
	Rules                     map[string]RewriteRule
	Urlrw                     Rewriter
	Cookierw                  Rewriter
	RewritingContent          bool
	req                       *http.Request
	GetLastOriginContextKeyFn func(*http.Request) string
}

func NewHeaderRewriter(configs ...func(cfg *Config)) *HeaderRewriter {
	c := makeConfig(configs...)
	return &HeaderRewriter{
		Prefix: c.HeaderPrefix,
		Rules:  c.HeaderRules,
		Urlrw:  c.Defmod,
		// Cookierw: c.CookieRewriter,
		// RewritingContent: c.ContentRewriter != nil,
	}
}
func (hrw *HeaderRewriter) SetRequest(r *http.Request) {
	hrw.req = r
}

func (hrw *HeaderRewriter) RewriteHeaders(headers http.Header) http.Header {
	rewritten := http.Header{}
	for key, _ := range headers {
		if val, ok := headers[key]; ok && len(val) > 0 {
			for _, v := range val {
				newkey, newval := hrw.rewriteHeader(key, v)
				rewritten.Add(newkey, newval)
			}
		}
		// } else {
		// 	newkey, newval := hrw.rewriteHeader(key, headers.Get(key))
		// 	rewritten.Add(newkey, newval)
		// }
	}
	//如果服务端不反悔AccessAllowOrigin ，动态追加
	if hrw.GetLastOriginContextKeyFn == nil {
		return rewritten
	}
	lastOrigin := hrw.GetLastOriginContextKeyFn(hrw.req)
	if headers.Get("Access-Control-Allow-Origin") == "" && lastOrigin != "" {
		rewritten.Set("Access-Control-Allow-Origin", lastOrigin)
	}
	return rewritten
}

func (hrw HeaderRewriter) rewriteHeader(name, value string) (string, string) {
	if name == "Location" {
		// ？？？如果没有http，补充主RequestURI？
		if !strings.HasPrefix(value, "http") {
			u, err := hrw.req.URL.Parse(value)
			if err == nil {
				value = u.String()
			}
		}
	}
	switch hrw.Rules[name] {
	case Keep:
		return name, value
	case Refresh:
		re, err := regexp.Compile(`.*?(?:URL|url)\s*=\s*(.+?)`)
		if hrw.Urlrw != nil && err == nil {
			match := re.FindAllStringSubmatchIndex(value, -1)
			if len(match) > 0 {
				v := hrw.Urlrw.Rewrite([]byte(value[match[0][2]:]))
				return name, value[:match[0][2]] + string(v)
			}

		}
		return name, value
	case UrlRewrite:

		if hrw.Urlrw != nil {
			v := hrw.Urlrw.Rewrite([]byte(value))
			return name, string(v)
		}
		return name, value
	case PrefixIfContentRewrite:
		if hrw.RewritingContent {
			return hrw.Prefix + name, value
		}
		return name, value
	case PrefixIfUrlRewrite:
		if hrw.Urlrw != nil {
			return hrw.Prefix + name, value
		}
		return name, value
	case ContentLength:
		if value == "0" {
			return name, value
		}
		// if not rewriting content, attempt to use the
		// length value
		if !hrw.RewritingContent {
			if lenth, err := strconv.Atoi(value); err == nil {
				return name, strconv.FormatInt(int64(lenth), 10)
			}
		}
		return hrw.Prefix + name, value
	case Cookie:
		if hrw.Cookierw != nil {
			return name, string(hrw.Cookierw.Rewrite([]byte(value)))
			//               return self.rwinfo.cookie_rewriter.rewrite(value)
		}
		return name, value
	case Prefix:
		return hrw.Prefix + name, value
	case Subfix:
		return "XX-" + name + "-Report-Only", "default-src 'self' *;"
	case StrictTransportSecurity:
		return "Strict-Transport-Security", "max-age=0; includeSubDomains"
	case AccessControlAllowOrigin:
		if value == "*" || hrw.GetLastOriginContextKeyFn == nil {
			return name, value
		}
		fmt.Println("AccessControlAllowOrigin")
		// fmt.Println()
		lastOrigin := hrw.GetLastOriginContextKeyFn(hrw.req)
		if lastOrigin != "" {
			return name, lastOrigin
		}
		fmt.Println(hrw.req.Header.Get("Referer"))
		fmt.Println(hrw.req.Header.Get("Origin"))
		fmt.Println(hrw.req.Header.Get("Origin"))
		fmt.Println(hrw.GetLastOriginContextKeyFn(hrw.req))
		fmt.Println(hrw.GetLastOriginContextKeyFn(hrw.req))
		fmt.Println(hrw.req.Referer())
		if hrw.Urlrw != nil {
			referer := hrw.req.Header.Get("Referer")
			if referer != "" {
				value = referer
			}
			v := hrw.Urlrw.Rewrite([]byte(value))
			uri := fasthttp.AcquireURI()
			uri.Parse(nil, v)
			uri.SetPathBytes(nil)
			uri.SetQueryStringBytes(nil)
			vv := uri.FullURI()
			if vv[len(vv)-1] == '/' {
				vv = vv[:len(vv)-1]
			}
			value = string(vv)
			fasthttp.ReleaseURI(uri)
			return name, value
		}

		// only from www.emerald.com

		return name + "GGG", "*"
	case XFrameOptonis:
		return "_D_" + name, value + ",ALLOW-FROM=*"
	}
	return name, value
}

// func (hrw *HeaderRewriter) AddCacheHeaders(headers map[string]string) {
// }

var DefaultHeaderRewriters = map[string]RewriteRule{
	"Access-Control-Allow-Origin":      AccessControlAllowOrigin,
	"Access-Control-Allow-Credentials": PrefixIfUrlRewrite,
	"Access-Control-Expose-Headers":    PrefixIfUrlRewrite,
	"Access-Control-Max-Age":           PrefixIfUrlRewrite,
	"Access-Control-Allow-Methods":     PrefixIfUrlRewrite,
	"Access-Control-Allow-Headers":     PrefixIfUrlRewrite,

	"Accept-Patch":  Keep,
	"Accept-Ranges": Keep,

	"Age": Prefix,

	"Allow": Keep,

	"Alt-Svc":       Prefix,
	"Cache-Control": Prefix,

	"Connection": Prefix,

	"Content-Base":                        UrlRewrite,
	"Content-Disposition":                 Keep,
	"Content-Encoding":                    PrefixIfContentRewrite,
	"Content-Language":                    Keep,
	"Content-Length":                      ContentLength,
	"Content-Location":                    UrlRewrite,
	"Content-Md5":                         Prefix,
	"Content-Range":                       Keep,
	"Content-Security-Policy":             Subfix,
	"Content-Security-Policy-Report-Only": Prefix,
	"Content-Type":                        Keep,

	"Date": Keep,

	"Etag":    Prefix,
	"Expires": Prefix,

	"Last-Modified": Prefix,
	"Link":          Keep,
	"Location":      UrlRewrite,

	"P3p":    Prefix,
	"Pragma": Prefix,

	"Proxy-Authenticate": Keep,

	"Public-Key-Pins": Prefix,
	"Retry-After":     Prefix,
	"Server":          Prefix,

	"Set-Cookie": Cookie,

	"Strict-Transport-Security": StrictTransportSecurity,

	"Trailer":           Prefix,
	"Transfer-Encoding": Prefix,
	"Tk":                Prefix,

	"Upgrade":                   Prefix,
	"Upgrade-Insecure-Requests": Prefix,

	"Vary": Prefix,

	"Via": Prefix,

	"Warning": Prefix,

	"Www-Authenticate": Keep,

	"X-Frame-Options":  XFrameOptonis,
	"X-Xss-Protection": Prefix,
	"Refresh":          Refresh,
}
