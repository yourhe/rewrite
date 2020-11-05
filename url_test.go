package rewrite

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
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
		{"http://.www.webofknowledge.com?", "http://www.webofknowledge.com.wf"},
		{"https://www.webofknowledge.com?", "http://www.webofknowledge.com.wf?__dp=https"},
		{"http://www.webofknowledge.com:84?", "http://www.webofknowledge.com.wf?__dp=http|84"},
		{"//www.webofknowledge.com?", "https://www.webofknowledge.com.wf"},
		{"/a/b?", "/a/b?"},
		{"a/b?", "a/b?"},
		{"../a/b?", "../a/b?"},
		{"https://images.drcnet-sod.com/javascript/jquery.min.js", "https://images.drcnet-sod.com.wf/javascript/jquery.min.js?__dp=https"},
		{"https://onlinelibrary.wiley.com/action/ajaxShowPubInfo?accordionHeadingWrapper=h2&ajax=true&displayAlmetricDropzone=true&displayCitedByLink=true&doi=10.1002%2Fjps.23423&pbContext=%3BrequestedJournal%3Ajournal%3A15206017%3Barticle%3Aarticle%3Adoi%5C%3A10.1002%2Fjps.23423%3Bpage%3Astring%3AArticle%2FChapter+View%3Bctype%3Astring%3AJournal+Content%3BsubPage%3Astring%3AAbstract%3Bwebsite%3Awebsite%3Apericles%3Bjournal%3Ajournal%3A19302169%3Bissue%3Aissue%3Adoi%5C%3A10.1002%2Fjps.v102.3%3Bwgrou", "a"},
	})
	rw := NewURLRewriter("https://www.webofknowledge.com/abc/d", "wf", "http", true, 0)
	// rw := NewUrlRewriter("//www.webofknowledge.com", "http://${host}.wf")
	// rw.ProtocolOnQuery(true)
	testRewriteCases(t, rw, cases)
}
func TestPortAndProtocolToDomianNewUrlsinRewriter(t *testing.T) {
	cases := stringTestCases([]stringTestCase{
		{"http://www.webofknowledge.com:80", "http://h-port-80.www.webofknowledge.com.wf"},
		{"http://www.webofknowledge.com", "http://h-port-80.www.webofknowledge.com.wf"},
		{"https://www.webofknowledge.com", "http://hs.www.webofknowledge.com.wf"},
	})
	rw := NewURLRewriter("https://www.webofknowledge.com/abc/d", "wf", "http", true, 0)
	// rw := NewUrlRewriter("//www.webofknowledge.com", "http://${host}.wf")
	// rw.ProtocolOnQuery(true)
	testRewriteCases(t, rw, cases)
}

func TestPortAndProtocolToPreduceathNewUrlsinRewriter(t *testing.T) {
	cases := stringTestCases([]stringTestCase{
		{"http://www.webofknowledge.com:80", "http://wf/--/com/webofknowledge/www/_/?__dp=http|80"},
		{"http://www.webofknowledge.com", "http://wf/--/com/webofknowledge/www/_/"},
		{"https://www.webofknowledge.com", "http://wf/--/com/webofknowledge/www/hs/_/?__dp=https"},
	})
	rw := NewURLRewriter("https://www.webofknowledge.com/abc/d", "wf", "http", true, 1)
	// rw := NewUrlRewriter("//www.webofknowledge.com", "http://${host}.wf")
	// rw.ProtocolOnQuery(true)
	testRewriteCases(t, rw, cases)
}
func TestWOSNewUrlsinRewriterProcessRelative(t *testing.T) {
	cases := stringTestCases([]stringTestCase{
		{"http://.www.webofknowledge.com?", "https://www.webofknowledge.com.wf"},
		{"https://www.webofknowledge.com?", "https://www.webofknowledge.com.wf?__dp=https"},
		{"http://www.webofknowledge.com:84?", "https://www.webofknowledge.com.wf?__dp=http|84"},
		{"//www.webofknowledge.com?", "https://www.webofknowledge.com.wf"},
		{"/a/b?", "https://www.webofknowledge.com.wf/a/b?__dp=https"},
		{"a/b?", "https://www.webofknowledge.com.wf/abc/a/b?__dp=https"},
		{"../a/b?", "https://www.webofknowledge.com.wf/a/b?__dp=https"},
		{"https://images.drcnet-sod.com/javascript/jquery.min.js", "https://images.drcnet-sod.com.wf/javascript/jquery.min.js?__dp=https"},
		{"https://ieeexplore.ieee.org/Xplore/home.jsp", "https://ieeexplore.ieee.org.wf/Xplore/home.jsp"},
	})
	rw := NewURLRewriterRelativePath("https://www.webofknowledge.com/abc/d", "wf", "https", false, 0)
	// rw := NewUrlRewriter("//www.webofknowledge.com", "http://${host}.wf")
	// rw.ProtocolOnQuery(true)
	testRewriteCases(t, rw, cases)
}

func TestWOSNewUrlsinSWRRewriter(t *testing.T) {

	cases := stringTestCases([]stringTestCase{
		{"http://www.webofknowledge.com", "https://wf/--/com/webofknowledge/www/_"},
		{"http://www.webofknowledge.com?", "https://wf/--/com/webofknowledge/www/_?"},
		{"http://www.webofknowledge.com/", "https://wf/--/com/webofknowledge/www/_/"},
		{"http://www.webofknowledge.com/?", "https://wf/--/com/webofknowledge/www/_/?"},
		{"http://www.webofknowledge.com/a", "https://wf/--/com/webofknowledge/www/_/a"},
		{"https://www.webofknowledge.com?a=b", "https://wf/--/com/webofknowledge/www/_?a=b&__dp=https"},
		{"/a/b?", "https://wf/--/com/webofknowledge/www/_/a/b?__dp=https"},
		{"a/b?", "https://wf/--/com/webofknowledge/www/_/abc/a/b?__dp=https"},
		{"/../a/b?", "https://wf/--/com/webofknowledge/www/_/../a/b?__dp=https"},
		// {"http://jggw.cnki.net/sso/home/check?returnurl=http%3A%2F%2Fjggw.cnki.net%2FODCC&timestamp=1584241472.59793&appid=odcc_81&sign=ba51d40cdc4817122c8dd5cd3b79974d8e2074ed", ""},
		{"../images/gb/icon-d.gif", "https://wf/--/com/webofknowledge/www/_/images/gb/icon-d.gif?__dp=https"},
		{"images/gb/icon-d.gif", "https://wf/--/com/webofknowledge/www/_/abc/images/gb/icon-d.gif?__dp=https"},
		{"?curpage=2&RecordsPerPage=20&QueryID=9&ID=&turnpage=1&tpagemode=L&dbPrefix=SCDB&Fields=&DisplayMode=listmode&PageName=ASP.brief_default_result_aspx&isinEn=1&", "https://wf/--/com/webofknowledge/www/_/abc/a?curpage=2&RecordsPerPage=20&QueryID=9&ID=&turnpage=1&tpagemode=L&dbPrefix=SCDB&Fields=&DisplayMode=listmode&PageName=ASP.brief_default_result_aspx&isinEn=1&&__dp=https"},
		{"javascript:__doPostBack('Button1','')", "javascript:__doPostBack('Button1','')"},
		{"javascript:__doPostBack('Button1','');", "javascript:__doPostBack('Button1','');"},
		{"&#xA;                    http://doi.cnki.net/doi/Resolution/Handler?doi= 10.1016/j.bbapap.2020.140410", "https://wf/--/net/cnki/doi/_/doi/Resolution/Handler?doi= 10.1016/j.bbapap.2020.140410"},
		{"javascript:void();", "javascript:void();"},
		{"#", "#"},
	})
	rw := NewURLRewriter("https://www.webofknowledge.com/abc/a?s", "wf", "https", true, 1)
	// rw := NewUrlRewriter("//www.webofknowledge.com", "http://${host}.wf")
	// rw.ProtocolOnQuery(true)
	testRewriteCases(t, rw, cases[:4])
}

func TestWanfangLocation(t *testing.T) {
	cases := stringTestCases([]stringTestCase{
		{"http://common.wanfangdata.com.cn/pay/submitWeb.do?webDownResourceRequest=eyJyZXF1ZXN0X3VybCI6Imh0dHA6Ly9jb21tb24ud2FuZmFuZ2RhdGEuY29tLmNuL3BheS9zdWJtaXRXZWIuZG8iLCJ3ZWJEb3duUmVzb3VyY2VQYXJhbSI6eyJ1bml0IjoiNSIsImxhbmd1YWdlIjoiY2hpIiwicmVzb3VyY2VfdHlwZSI6InBlcmlvIiwic291cmNlIjoiV0YiLCJkb3duX3Jlc291cmNlX2lkIjoicnloeGd5MjAxOTEyMDA0IiwiZG93bl9yZXNvdXJjZV90aXRsZSI6IkFTUOS4ieWFg-WkjeWQiOmpseayueS9k-ezu-WcqOmVv-WyqeW_g-S4reeahOi_kOenu-inhOW-iyIsInN0YXR1cyI6MSwiaXNvYSI6ZmFsc2UsImZpcnN0cHVibGlzaCI6bnVsbCwiaXNyZXN1bHQiOmZhbHNlLCJiYWNrdXJsIjoiIiwicGllY2UiOjAsImNvcHkiOjAsIml0ZW0iOjAsIm5leHQiOjAsImt3b3JkIjowLCJtd29yZCI6MCwiaW5kaXZpZHVhbCI6MCwicGFnZSI6MCwicmVzb3VyY2VMaW1pdHNEVE8iOm51bGwsInJlc291cmNlTGltaXRzU1REIjpudWxsfX0", ""},
		{"//oup.silverchair-cdn.com/UI/app/vendor/jquery-2.2.4.js", "http://oup.silverchair-cdn.com/UI/app/vendor/jquery-2.2.4.js?__dp=https"},
	})
	rw := NewURLRewriter("https://www.webofknowledge.com/abc/a?s", "wf", "http", true, 0)
	testRewriteCases(t, rw, cases)
}

func TestNGTemplate(t *testing.T) {
	// http://www.specialsci.cn/searchresult?complex=%2Basp&from=Fast

	cases := stringTestCases([]stringTestCase{
		{"/../detail/{{s.gui}}?view=detailed", ""},
		{"../detail/{{s.gui}}?view=detailed", ""},
		{"https://www.sciencedirect.com", ""},
	})
	rw := NewURLRewriterRelativePath("http://www.specialsci.cn/searchresult/a?complex=%2Basp&from=Fast", "wf:8989", "http", true, 0)

	testRewriteCases(t, rw, cases)
	u := fasthttp.AcquireURI()
	u.Parse(nil, []byte("http://www.specialsci.cn/detail/202001100740289732848460?view=detailed"))
	u.Update("app/app.1b070094.js")
	fmt.Println(u.String())

}

func TestCNKI(t *testing.T) {
	cases := stringTestCases([]stringTestCase{
		// {"//bianke.cnki.net/adfiles/ad/R018.js?sc=B024", ""},
		{"Common/RedirectPage?sfield=FN&", "Common/RedirectPage?sfield=FN&"},
	})
	rw := NewURLRewriterRelativePath("https://kns.cnki.net/KCMS/detail/detail.aspx?dbcode=CJFQ&dbname=CJFDAUTO&filename=HNXB202009019&v=MDA1NjZZUzdEaDFUM3FUcldNMUZyQ1VSN3FmWU9Sb0Z5dmtVN3JPTFNQVGJMRzRITkhNcG85RWJZUjhlWDFMdXg=", "wf", "http", true, 1)
	testRewriteCases(t, rw, cases)

}

func TestQdexam(t *testing.T) {
	cases := stringTestCases([]stringTestCase{
		{"/js/home.js?v=1.3", "http://www.qdexam.com.wf/js/home.js?v=1.3"},
		{"js/home.js?v=1.3", "http://www.qdexam.com.wf/js/home.js?v=1.3"},
	})
	rw := NewURLRewriterRelativePath("http://www.qdexam.com/main", "wf", "http", true, 0)
	testRewriteCases(t, rw, cases)
	cases = stringTestCases([]stringTestCase{
		{"/js/home.js?v=1.3", "http://wf/--/com/qdexam/www/_/js/home.js?v=1.3"},
		{"js/home.js?v=1.3", "http://wf/--/com/qdexam/www/_/js/home.js?v=1.3"},
	})
	rw = NewURLRewriterRelativePath("http://www.qdexam.com/main", "wf", "http", true, 1)
	testRewriteCases(t, rw, cases)

}
