package rewrite

import (
	"testing"
)

func TestUrlRewriter(t *testing.T) {
	cases := stringTestCases([]stringTestCase{
		{"", ""},
		// {"http://youtube.com/{}", "http://youtube.com/{}"},
		{"https://a.com/{{afaf.com+dd}}}", "http://b.tv/{{afaf.com}}}"},
		{"http://a.com/", "http://b.tv/"},
		{"/relative/url", "http://b.tv/relative/url"},
		{"http://a.com/path?query=a", "http://b.tv/path?query=a"},
		{"http://a.com:80//static/bootstrap/bootstrap.css", "http://b.tv:80//static/bootstrap/bootstrap.css"},
	})

	rw := NewUrlRewriter("http://a.com", "http://${host}b.tv")
	testRewriteCases(t, rw, cases)
}

func TestSchemeChangeUrlRewriter(t *testing.T) {
	cases := stringTestCases([]stringTestCase{
		{"", ""},
		{"http://youtube.com", "http://youtube.com/"},
		{"http://a.com", "https://b.tv/"},
		{"/relative/url", "https://b.tv/relative/url"},
		{"https://a.com/path?query=a", "https://b.tv/path?query=a"},
	})

	rw := NewUrlRewriter("http://a.com", "https://b.tv")
	testRewriteCases(t, rw, cases)
}

func TestRelativeUrlRewriter(t *testing.T) {
	cases := stringTestCases([]stringTestCase{
		{"", ""},
		{"http://youtube.com", "http://youtube.com/"},
		{"http://a.com", "."},
		{"/relative/url", "./relative/url"},
		{"https://a.com/path?query=a", "./path?query=a"},
	})

	rw := NewRelativeUrlRewriter("http://a.com")
	testRewriteCases(t, rw, cases)
}

func TestHostRelativeUrlRewriter(t *testing.T) {
	cases := stringTestCases([]stringTestCase{
		{"", ""},
		{"http://youtube.com", "../youtube.com"},
		{"http://a.com", "../a.com"},
		{"/relative/url", "../a.com/relative/url"},
		{"https://a.com/path?query=a", "../a.com/path?query=a"},
	})

	rw := NewHostRelativeUrlRewriter("http://a.com")
	testRewriteCases(t, rw, cases)
}

func TestLetDigHOSTUrlRewriter(t *testing.T) {
	cases := stringTestCases([]stringTestCase{
		{"http://wanfangdata.com.cn", "http://wanfangdata-com-cn.pk.com/"},
	})

	rw := NewLetDigHostUrlRewriter("http://wanfangdata.com.cn", "http://${host}.pk.com")
	testRewriteCases(t, rw, cases)
}

func TestSWUrlRewriter(t *testing.T) {
	cases := stringTestCases([]stringTestCase{
		{"http://wanfangdata.com.cn", "http://pk.com/--wanfangdata-com-cn/"},
		{"http://wanfangdata.com.cn/sns", "http://pk.com/--wanfangdata-com-cn/sns"},
		{"http://wanfangdata.com.cn:80/sns/?bb=1", "http://pk.com/--wanfangdata-com-cn/sns/?bb=1&__dp=http|80"},
		{"https://wanfangdata.com.cn/sns/", "http://pk.com/--wanfangdata-com-cn/sns/?__dp=https"},
		{"https://wanfangdata.com.cn:444/sns/", "http://pk.com/--wanfangdata-com-cn/sns/?__dp=https|444"},
	})

	rw := NewSWHostUrlRewriter("//wanfangdata.com.cn", "http://pk.com/--${host}")
	rw.ProtocolOnQuery(true)
	testRewriteCases(t, rw, cases)
}

func TestSWReduceHostUrlRewriter(t *testing.T) {
	cases := stringTestCases([]stringTestCase{
		{"http://wanfangdata.com.cn", "http://pk.com/--/cn/com/wanfangdata/@"},
		{"http://wanfangdata.com.cn/sns", "http://pk.com/--/cn/com/wanfangdata/@/sns"},
		{"http://wanfangdata.com.cn:80/sns/?bb=1", "http://pk.com/--/cn/com/wanfangdata/@/sns/?bb=1&__dp=http|80"},
		{"https://wanfangdata.com.cn/sns/", "http://pk.com/--/cn/com/wanfangdata/@/sns/?__dp=https"},
		{"https://wanfangdata.com.cn:444/sns/", "http://pk.com/--/cn/com/wanfangdata/@/sns/?__dp=https|444"},
	})

	rw := NewSWReduceHostUrlRewriter("//wanfangdata.com.cn", "http://pk.com/--/${host}/@")
	rw.ProtocolOnQuery(true)
	testRewriteCases(t, rw, cases)

	cases = stringTestCases([]stringTestCase{
		{"https://www.wanfangdata.com.cn:444/sns/", "http://pk.com/--/cn/com/wanfangdata/www/@/sns/?__dp=https|444"},
	})

	rw = NewSWReduceHostUrlRewriter("//www.wanfangdata.com.cn", "http://pk.com/--/${host}/@")
	rw.ProtocolOnQuery(true)
	testRewriteCases(t, rw, cases)
}

func TestWwwSinomedAcCnUrlsinRewriter(t *testing.T) {
	cases := stringTestCases([]stringTestCase{
		{"http://www.sinomed.ac.cn:80//static/bootstrap/bootstrap.css", "http://www.sinomed.ac.cn.wf//static/bootstrap/bootstrap.css?__dp=http|80"},
		{"http://www.sinomed.ac.cn:81//static/bootstrap/bootstrap.css", "http://www.sinomed.ac.cn.wf//static/bootstrap/bootstrap.css?__dp=http|81"},
	})

	rw := NewUrlRewriter("//www.sinomed.ac.cn", "http://${host}.wf")
	rw.ProtocolOnQuery(true)
	testRewriteCases(t, rw, cases)
}

func TestWOSUrlsinRewriter(t *testing.T) {
	cases := stringTestCases([]stringTestCase{
		{"http://www.webofknowledge.com?", "http://www.webofknowledge.com.wf"},
	})

	rw := NewUrlRewriter("//www.webofknowledge.com", "http://${host}.wf")
	rw.ProtocolOnQuery(true)
	testRewriteCases(t, rw, cases)
}
