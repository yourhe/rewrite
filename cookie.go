package rewrite

import "net/http"

type CookieRewriter struct {
}

func NewCookieRewriter(configs ...func(*Config)) *CookieRewriter {
	// c := makeConfig(configs...)
	return &CookieRewriter{}
}

func (crw *CookieRewriter) Rewrite(p []byte) []byte {
	// TODO
	ck := ParseCookies(string(p))
	ck.Secure = false
	ck.HttpOnly = false
	ck.Domain = ""
	return []byte(ck.String())
}
func ParseCookies(s string) *http.Cookie {
	return (&http.Response{Header: http.Header{"Set-Cookie": {s}}}).Cookies()[0]
}
