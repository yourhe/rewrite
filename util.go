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

func ToSWLetDigHyp(host string, protocol string) string {
	if protocol == "" {
		protocol = "http"
	}
	h, _, _ := VaildHTTPPort(host)
	h = strings.Replace(h, ".", "-", -1)
	// if p != "" {
	// 	p = fmt.Sprintf("%s[%s]", protocol, p)
	// } else {
	// 	p = protocol
	// }

	// if p != "" {
	// 	h = fmt.Sprintf("%s-%s", h, p)
	// }

	return h
}

func ToSWReduceLetDigHyp(host string, protocol string) string {
	return toSWReduceLetDigHyp(host, protocol, true)
}
func NotProcessProtocolToSWReduceLetDigHyp(host string, protocol string) string {
	return toSWReduceLetDigHyp(host, protocol, false)
}
func toSWReduceLetDigHyp(host string, protocol string, processProtocol bool) string {
	if protocol == "" {
		protocol = "http"
	}
	h, p, _ := VaildHTTPPort(host)
	for {
		if h[0] == '.' {
			h = h[1:]
			continue
		}
		break
	}
	rh := strings.Split(h, ".")

	for i := 0; i < len(rh); i++ {

		if rh[i] == "" {
			continue
		}
		j := len(rh) - i - 1
		if j <= i {
			break
		}
		t := rh[i]
		rh[i] = rh[j]
		rh[j] = t
	}
	h = strings.Join(rh, "/")

	protocol = strings.Replace(protocol, "http", "h", -1)
	if p != "" {
		protocol = protocol + "-port-" + p

	}
	if formatBProtocol(protocol) != "" && processProtocol {
		h = h + "/" + protocol

	}
	return h
}

func formatBProtocol(protocol string) string {
	switch protocol {
	case "h", "h-port-80":
		return ""
	case "hs-port-443":
		protocol = "hs"
	}
	return protocol
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

//EncodeHost xxx.xx:89 => port-89.xxx.xx
func EncodeHost(host, perfix, s string) string {
	// set port to begin
	host, port, vailded := VaildHTTPPort(host)
	if !vailded {
		return host
	}
	if port != "" && (port != "80" && port != "443") {
		port = "port-" + port
	} else {
		port = ""
	}

	return fmt.Sprintf("%s%s%s", port, s, host)
}

//DecodeHost port-89.xxx.xx => xxx.xx:89
func DecodeHost(host, perfix, s string) string {
	if perfix == "" {
		perfix = "port-"
	}

	if !strings.Contains(host, perfix) {
		return host
	}

	if i := strings.Index(host, "."+perfix); i > -1 {
		host = host[i+1:] + "." + host[:i]
	}
	if !strings.HasPrefix(host, perfix) {
		return host
	}

	idx := strings.Index(host, s)
	if idx < len(perfix) {
		return host
	}
	strPort := host[len(perfix):idx]
	if strPort != "" {
		strPort = ":" + strPort
	}
	// port,_ := strconv.Atoi(strPort)
	return fmt.Sprintf("%s%s", host[idx+1:], strPort)
}
