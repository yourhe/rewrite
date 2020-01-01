package rewrite

import (
	"fmt"
	"net"
	"strings"

	"golang.org/x/net/publicsuffix"
)

// ToLetDigHyp ...
// www.wanfangdata.com.cn => www-wanfangdata-com-cn
// www.wanfangdata.com.cn:80 => www-wanfangdata-com-cn-port-80
func ToLetDigHyp(host string) string {
	h, p, _ := VaildHTTPPort(host)
	d, _ := publicsuffix.EffectiveTLDPlusOne(h)
	l := len(h) - len(d)
	if l < 0 {
		l = 0
	}
	h = h[:l] + strings.Replace(h[l:], ".", "-", -1)
	if p != "" {
		h = fmt.Sprintf("%s-port-%s", h, p)
	}

	return h
}

//VaildHTTPPort host中存在port返回true
func VaildHTTPPort(host string) (string, string, bool) {
	h, p, err := net.SplitHostPort(host)
	if err != nil {
		return host, "", false
	}
	if p != "" {
		return h, p, true
	}
	return h, p, false

}
