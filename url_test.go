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
		{"http://www.webofknowledge.com?", "http://www.webofknowledge.com.wf"},
	})

	rw := NewUrlRewriter("//www.webofknowledge.com", "http://${host}.wf")
	rw.ProtocolOnQuery(true)
	testRewriteCases(t, rw, cases)
}

func TestWOSNewUrlsinRewriter(t *testing.T) {
	cases := stringTestCases([]stringTestCase{
		{"http://.www.webofknowledge.com?", "https://www.webofknowledge.com.wf"},
		{"https://www.webofknowledge.com?", "https://www.webofknowledge.com.wf?__dp=https"},
		{"http://www.webofknowledge.com:84?", "https://www.webofknowledge.com.wf?__dp=http|84"},
		{"//www.webofknowledge.com?", "https://www.webofknowledge.com.wf"},
		{"/a/b?", "/a/b?"},
		{"a/b?", "a/b?"},
		{"../a/b?", "../a/b?"},
		{"http://images.drcnet-sod.com/javascript/jquery.min.js", "../a/b?"},
	})
	rw := NewURLRewriter("https://www.webofknowledge.com.wf/abc/d", "wf", "https", true, 0)
	// rw := NewUrlRewriter("//www.webofknowledge.com", "http://${host}.wf")
	// rw.ProtocolOnQuery(true)
	testRewriteCases(t, rw, cases)
}

func TestWOSNewUrlsinSWRRewriter(t *testing.T) {
	cases := stringTestCases([]stringTestCase{
		{"http://www.webofknowledge.com?", "https://wf/--/com/webofknowledge/www/_"},
		{"http://www.webofknowledge.com/a?", "https://wf/--/com/webofknowledge/www/_/a"},
		{"https://www.webofknowledge.com?a=b", "https://wf/--/com/webofknowledge/www/_?a=b&__dp=https"},
		{"/a/b?", "https://wf/--/com/webofknowledge/www/_/a/b?__dp=https"},
		{"a/b?", "https://wf/--/com/webofknowledge/www/_/abc/a/b?__dp=https"},
		{"/../a/b?", "https://wf/--/com/webofknowledge/www/_/../a/b?__dp=https"},
		// {"http://jggw.cnki.net/sso/home/check?returnurl=http%3A%2F%2Fjggw.cnki.net%2FODCC&timestamp=1584241472.59793&appid=odcc_81&sign=ba51d40cdc4817122c8dd5cd3b79974d8e2074ed", ""},
		{"../images/gb/icon-d.gif", "https://wf/--/com/webofknowledge/www/_/images/gb/icon-d.gif?__dp=https"},
		{"images/gb/icon-d.gif", "https://wf/--/com/webofknowledge/www/_/abc/images/gb/icon-d.gif?__dp=https"},
		{"?curpage=2&RecordsPerPage=20&QueryID=9&ID=&turnpage=1&tpagemode=L&dbPrefix=SCDB&Fields=&DisplayMode=listmode&PageName=ASP.brief_default_result_aspx&isinEn=1&", "https://wf/--/com/webofknowledge/www/_/abc/a?curpage=2&RecordsPerPage=20&QueryID=9&ID=&turnpage=1&tpagemode=L&dbPrefix=SCDB&Fields=&DisplayMode=listmode&PageName=ASP.brief_default_result_aspx&isinEn=1&&__dp=https"},
		{"javascript:__doPostBack('Button1','')", "javascript:__doPostBack('Button1','')"},
		{"javascript:__doPostBack('Button1','')", "javascript:__doPostBack('Button1','');"},
		{"&#xA;                    http://doi.cnki.net/doi/Resolution/Handler?doi= 10.1016/j.bbapap.2020.140410", "https://wf/--/net/cnki/doi/_/doi/Resolution/Handler?doi= 10.1016/j.bbapap.2020.140410&__dp=http"},
		{"javascript:void();", "javascript:void();"},
	})
	rw := NewURLRewriter("https://www.webofknowledge.com/abc/a?s", "wf", "https", true, 1)
	// rw := NewUrlRewriter("//www.webofknowledge.com", "http://${host}.wf")
	// rw.ProtocolOnQuery(true)
	testRewriteCases(t, rw, cases)
}
