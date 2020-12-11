package rewrite

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"testing"
)

func TestRewriteHtml(t *testing.T) {
	urlrw := NewUrlRewriter("http://a.com", "https://b.tv")
	rw := NewHtmlRewriter(urlrw)
	cases := stringTestCases([]stringTestCase{
		// {"", ""},
		{htmlNoChange, htmlNoChange},
		// {basicHtmlRewriteIn, basicHtmlRewriteOut},
	})
	testRewriteCases(t, rw, cases)
}

func TestRewriteReader_Read(t *testing.T) {
	urlrw := NewUrlRewriter("http://a.com:90", "https://b.tv")
	rw := NewHtmlRewriter(urlrw)

	// bf := bytes.NewBufferString(basicHtmlRewriteIn)
	bf := bytes.NewBufferString(basicHTMLRewriteIn)
	r := rw.NewReader(bf)
	r.AddInsert("head", "<base></base>", true)
	b, err := ioutil.ReadAll(r)
	// b := make([]byte, 512)
	// l, err := r.Read(b)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(b))
}

func TestNewReader(t *testing.T) {
	bf := bytes.NewBufferString(basicHTMLRewriteIn)
	r := NewRewriteReader(bf)
	// NewUrlRewriter("", "")
	r.SetTagRewriter("a", "href", nil)
	b, err := ioutil.ReadAll(r)
	fmt.Println(string(b), err)
}

func TestScriptTagTextRewrite(t *testing.T) {
	var raw = []byte(`<script>location.href = "abc" </script>`)
	bf := bytes.NewBuffer(raw)
	r := NewRewriteReader(bf)
	r.SetJavascriptRewriter(NewJavaScriptRewrite())
	got, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}
	if reflect.DeepEqual(raw, got) {
		t.Errorf("want %s, but got %s", raw, got)
	}
	fmt.Println(string(got))
}
func TestCnkiPlan(t *testing.T) {
	f, err := os.Open("./testdata/cnki-loginapi-get.html")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	r := NewRewriteReader(f)
	// r.SetTagRewriter("a", "href", nil)
	b, err := ioutil.ReadAll(r)
	fmt.Println(string(b), err)
}

func TestCnkiHome(t *testing.T) {
	var raw = `<ul class="yw">
                        <li id="ywk" class="selected" onclick="LinkDb(this,'http://ref.cnki.net/ref')"><i></i><a href="javascript:void(0)">中国引文数据库</a></li>
        <input type="hidden" name="list" id="list" value="[&quot;清华大学&quot;,&quot;北京大学&quot;,&quot;北京师范大学&quot;,&quot;中国人民大学&quot;,&quot;山东大学&quot;,&quot;华东师范大学&quot;,&quot;浙江大学&quot;,&quot;天津大学&quot;,&quot;华南理工大学&quot;,&quot;华中科技大学&quot;,&quot;长春理工大学&quot;,&quot;上海财经大学&quot;,&quot;深圳大学&quot;,&quot;沈阳农业大学&quot;,&quot;浙江工商大学&quot;,&quot;华北水利水电大学&quot;,&quot;上海师范大学&quot;,&quot;武汉大学&quot;,&quot;中国科学技术大学&quot;,&quot;云南大学&quot;,&quot;济南大学&quot;,&quot;北京中医药大学&quot;,&quot;浙江中医药大学&quot;,&quot;厦门大学&quot;,&quot;北京交通大学&quot;,&quot;西安建筑科技大学&quot;,&quot;华中师范大学&quot;,&quot;福建师范大学&quot;,&quot;武汉纺织大学&quot;,&quot;华东理工大学&quot;,&quot;北京联合大学&quot;,&quot;北京林业大学&quot;,&quot;重庆建筑工程职业学院&quot;,&quot;上海交通大学&quot;,&quot;江西旅游商贸职业学院&quot;,&quot;上海对外经贸大学&quot;,&quot;江苏大学&quot;,&quot;南京工程学院&quot;,&quot;湖北警官学院&quot;,&quot;西南财经大学&quot;,&quot;南开大学&quot;,&quot;河北大学&quot;,&quot;南方科技大学&quot;,&quot;江苏师范大学&quot;,&quot;东南大学&quot;,&quot;肇庆学院&quot;,&quot;佛山科学技术学院&quot;,&quot;浙江财经大学&quot;,&quot;华南农业大学&quot;,&quot;广西民族大学&quot;,&quot;燕山大学&quot;,&quot;上海海洋大学&quot;,&quot;天津商业大学&quot;,&quot;青海民族大学&quot;,&quot;中国农业大学&quot;,&quot;郑州轻工业大学&quot;,&quot;合肥工业大学&quot;,&quot;广东科技学院&quot;,&quot;南京理工大学&quot;,&quot;乐山师范学院&quot;,&quot;河南大学&quot;,&quot;湖南大学&quot;,&quot;黄河水利职业技术学院&quot;,&quot;南京工业大学&quot;,&quot;北京石油化工学院&quot;,&quot;广西大学&quot;,&quot;陕西理工大学&quot;,&quot;东华大学&quot;,&quot;杭州师范大学&quot;,&quot;中国石油大学（北京）&quot;,&quot;华南理工大学广州学院&quot;,&quot;韩山师范学院&quot;,&quot;顺德职业技术学院&quot;,&quot;天津医科大学&quot;,&quot;淮阴工学院&quot;,&quot;华中农业大学&quot;,&quot;南京农业大学&quot;,&quot;湘潭大学&quot;,&quot;中央音乐学院&quot;,&quot;广州中医药大学&quot;,&quot;济南大学泉城学院&quot;,&quot;武汉工程大学&quot;,&quot;郑州市现代教育信息技术中心&quot;,&quot;暨南大学&quot;,&quot;湖南工学院&quot;,&quot;广西经贸职业技术学院&quot;,&quot;南京审计大学&quot;,&quot;中国矿业大学&quot;,&quot;西南大学（调试中）&quot;,&quot;东北林业大学&quot;,&quot;湖南科技大学&quot;,&quot;四川大学&quot;,&quot;上海电机学院&quot;,&quot;西北工业大学&quot;,&quot;广州大学&quot;,&quot;内蒙古工业大学&quot;,&quot;中华女子学院&quot;,&quot;中央民族大学&quot;,&quot;复旦大学&quot;,&quot;中国人民警察大学&quot;,&quot;山东女子学院&quot;,&quot;广西师范大学&quot;,&quot;安徽工程大学&quot;,&quot;中国海洋大学&quot;,&quot;河南科技大学&quot;,&quot;郑州大学&quot;,&quot;吉首大学&quot;,&quot;陕西师范大学&quot;,&quot;哈尔滨工业大学&quot;,&quot;青海建筑职业技术学院&quot;,&quot;上海第二工业大学&quot;,&quot;香港中文大学(深圳)&quot;,&quot;集美大学&quot;,&quot;哈尔滨商业大学&quot;,&quot;吉林体育学院&quot;,&quot;中南大学&quot;,&quot;上海外国语大学&quot;,&quot;西北农林科技大学&quot;,&quot;贵州师范大学&quot;,&quot;河北师范大学&quot;,&quot;江西中医药大学&quot;,&quot;安徽农业大学&quot;,&quot;湖南工程学院&quot;,&quot;东北大学&quot;,&quot;西安石油大学&quot;,&quot;大连理工大学&quot;,&quot;北京化工大学&quot;,&quot;辽宁师范大学&quot;,&quot;哈尔滨工程大学&quot;,&quot;湖北经济学院&quot;,&quot;中南民族大学&quot;,&quot;上海海事大学&quot;,&quot;海南师范大学&quot;,&quot;桂林旅游学院&quot;,&quot;上海理工大学&quot;,&quot;上海大学&quot;,&quot;山东师范大学&quot;,&quot;许昌学院&quot;,&quot;四川大学锦城学院&quot;,&quot;中南林业科技大学&quot;,&quot;四川省乐山第一中学&quot;,&quot;长沙民政职业技术学院&quot;,&quot;浙江水利水电学院&quot;,&quot;宁波职业技术学院&quot;,&quot;沈阳航空航天大学&quot;,&quot;南京林业大学&quot;,&quot;广西交通技师学院&quot;,&quot;苏州科技大学&quot;,&quot;东北师范大学&quot;,&quot;西南民族大学&quot;,&quot;中国传媒大学&quot;,&quot;广东财经大学&quot;,&quot;潍坊医学院&quot;,&quot;湖南财政经济学院&quot;,&quot;贵州财经大学&quot;,&quot;山东中医药大学&quot;,&quot;怀化学院&quot;,&quot;阜阳师范大学&quot;,&quot;北京社会管理职业学院&quot;,&quot;铜陵学院&quot;,&quot;成都理工大学&quot;,&quot;河北科技大学&quot;,&quot;嘉应学院&quot;,&quot;北京理工大学&quot;,&quot;扬州大学&quot;,&quot;西安邮电大学&quot;,&quot;西安工业大学&quot;,&quot;重庆医科大学&quot;,&quot;湖南师范大学&quot;,&quot;西安电子科技大学&quot;,&quot;兰州文理学院&quot;,&quot;中山大学&quot;,&quot;南通大学&quot;,&quot;南京大学&quot;,&quot;镇江市高等专科学校&quot;,&quot;山东英才学院&quot;,&quot;浙江传媒学院&quot;,&quot;深圳技术大学&quot;,&quot;宁波大学&quot;,&quot;中国地质大学(武汉)&quot;,&quot;中南财经政法大学&quot;,&quot;河池学院&quot;,&quot;西北大学&quot;,&quot;山东体育学院&quot;,&quot;东莞理工学院&quot;,&quot;华南师范大学&quot;,&quot;三江学院&quot;,&quot;广西中医药大学&quot;,&quot;内蒙古科技大学&quot;,&quot;河南城建学院&quot;,&quot;江西财经大学&quot;,&quot;安徽理工大学&quot;,&quot;淮南师范学院&quot;,&quot;兰州大学&quot;,&quot;三峡大学&quot;,&quot;郑州升达经贸管理学院&quot;,&quot;江西师范高等专科学校&quot;,&quot;长春大学&quot;,&quot;武汉理工大学&quot;,&quot;福建农林大学&quot;,&quot;广东职业技术学院&quot;,&quot;新乡医学院&quot;,&quot;四川美术学院&quot;,&quot;北京语言大学&quot;,&quot;黄山学院&quot;,&quot;安徽科技学院&quot;,&quot;唐山学院&quot;,&quot;兰州交通大学&quot;,&quot;巢湖学院&quot;,&quot;安徽大学&quot;,&quot;北京师范大学珠海分校&quot;,&quot;安徽工商职业学院&quot;,&quot;华侨大学&quot;,&quot;中国石油大学（华东）&quot;,&quot;西北民族大学&quot;,&quot;中国药科大学&quot;,&quot;山西大学&quot;,&quot;华北科技学院&quot;,&quot;同济大学&quot;,&quot;滁州学院&quot;,&quot;中国民航大学&quot;,&quot;浙江工业大学&quot;,&quot;重庆邮电大学&quot;,&quot;天津财经大学&quot;,&quot;江西环境工程职业学院&quot;,&quot;绍兴文理学院&quot;,&quot;曲阜师范大学&quot;,&quot;河北工业大学&quot;,&quot;周口师范学院&quot;,&quot;池州学院&quot;,&quot;上海海关学院&quot;,&quot;深圳大学城（调试中）&quot;,&quot;四川传媒学院&quot;,&quot;南昌大学&quot;,&quot;沧州医学高等专科学校&quot;,&quot;上海工程技术大学&quot;,&quot;山东农业大学&quot;,&quot;浙江师范大学&quot;,&quot;南昌工程学院&quot;,&quot;咸阳师范学院&quot;,&quot;宁波财经学院&quot;,&quot;湖北大学&quot;,&quot;吉林农业大学&quot;,&quot;武夷学院&quot;,&quot;安徽国际商务职业学院&quot;,&quot;内蒙古大学&quot;,&quot;华东政法大学&quot;,&quot;西安工程大学&quot;,&quot;重庆师范大学&quot;,&quot;重庆交通大学&quot;,&quot;西南医科大学&quot;,&quot;昆明理工大学&quot;,&quot;河南水利与环境职业学院&quot;,&quot;南华大学&quot;,&quot;浙江海洋大学&quot;,&quot;山西医科大学&quot;,&quot;苏州工艺美术职业技术学院&quot;,&quot;江西师范大学&quot;,&quot;山西师范大学&quot;,&quot;湖南工程职业技术学院&quot;,&quot;安徽新华学院&quot;,&quot;合肥师范学院&quot;,&quot;牡丹江师范学院&quot;,&quot;西北政法大学&quot;,&quot;常熟理工学院&quot;,&quot;河南财政金融学院&quot;,&quot;南方医科大学&quot;,&quot;温州医科大学&quot;,&quot;中央财经大学&quot;,&quot;北京城市学院&quot;,&quot;长春工业大学&quot;,&quot;重庆大学&quot;,&quot;中国科学院文献情报中心&quot;,&quot;成都大学&quot;,&quot;武汉音乐学院&quot;,&quot;福建医科大学&quot;,&quot;安徽工业大学&quot;,&quot;广州美术学院&quot;,&quot;四川工程职业技术学院&quot;,&quot;洛阳理工学院&quot;,&quot;仲恺农业工程学院&quot;,&quot;兰州理工大学&quot;,&quot;延边大学&quot;,&quot;桂林医学院&quot;,&quot;中国科学院大学&quot;,&quot;南昌航空大学&quot;,&quot;郑州西亚斯学院&quot;,&quot;中国科学院地理科学与资源研究所&quot;,&quot;中国科学院数学与系统科学研究院&quot;,&quot;中国科学院心理研究所&quot;,&quot;中国科学院生态环境研究中心&quot;,&quot;重庆工商大学&quot;,&quot;云南民族大学&quot;,&quot;齐鲁工业大学&quot;,&quot;天津工业大学&quot;,&quot;内蒙古师范大学&quot;,&quot;南阳理工学院&quot;,&quot;河北交通职业技术学院&quot;,&quot;安徽师范大学&quot;,&quot;北京服装学院&quot;,&quot;黑龙江大学&quot;,&quot;青海师范大学&quot;,&quot;天津城建大学&quot;,&quot;江苏建筑职业技术学院&quot;,&quot;南京师范大学&quot;,&quot;宁夏大学&quot;,&quot;四川建筑职业技术学院&quot;,&quot;上海政法学院&quot;,&quot;广西科技大学&quot;,&quot;重庆文理学院&quot;,&quot;桂林理工大学&quot;,&quot;鹤壁职业技术学院&quot;,&quot;福建中医药大学&quot;,&quot;厦门医学院&quot;,&quot;盐城工学院&quot;,&quot;长安大学&quot;,&quot;中原工学院&quot;,&quot;台州学院&quot;,&quot;太原科技大学&quot;,&quot;湖南农业大学&quot;,&quot;中国科学院云南天文台&quot;,&quot;中国科学院上海药物研究所&quot;,&quot;中国科学院新疆生态与地理研究所&quot;,&quot;中国科学院新疆理化技术研究所&quot;,&quot;中国科学院西北高原生物研究所&quot;,&quot;中国科学院青海盐湖研究所&quot;,&quot;中国科学院寒区旱区环境与工程研究所&quot;,&quot;中国科学院兰州化学物理研究所&quot;,&quot;中国科学院近代物理研究所&quot;,&quot;中国科学院兰州文献情报中心&quot;,&quot;中国科学院地球环境研究所&quot;,&quot;中国科学院国家授时中心&quot;,&quot;中国科学院西安光学精密机械研究所&quot;,&quot;中国科学院武汉病毒研究所&quot;,&quot;中国科学院西双版纳热带植物园&quot;,&quot;中国科学院昆明植物研究所&quot;,&quot;中国科学院昆明动物研究所&quot;,&quot;中国科学院光电技术研究所&quot;,&quot;中国科学院成都生物研究所&quot;,&quot;中国科学院成都文献情报中心&quot;,&quot;中国科学院深圳先进技术研究院&quot;,&quot;中国科学院广州生物医药与健康研究院&quot;,&quot;中国科学院广州地球化学研究所&quot;,&quot;中国科学院广州能源研究所&quot;,&quot;中国科学院广州电子技术研究所&quot;,&quot;中国科学院南海海洋研究所&quot;,&quot;中国科学院华南植物园&quot;,&quot;中国科学院亚热带农业生态研究所&quot;,&quot;中国科学院测量与地球物理研究所&quot;,&quot;中国科学院武汉文献情报中心&quot;,&quot;中国科学院武汉植物园&quot;,&quot;中国科学院水生生物研究所&quot;,&quot;中国科学院武汉物理与数学研究所&quot;,&quot;中国科学院福建物质结构研究所&quot;,&quot;中国科学院合肥物质科学研究院&quot;,&quot;中国科学院宁波材料技术与工程研究所&quot;,&quot;中国科学院苏州纳米技术与纳米仿生研究所&quot;,&quot;中国科学院南京地理与湖泊研究所&quot;,&quot;中国科学院南京地质古生物研究所&quot;,&quot;中国科学院南京土壤研究所&quot;,&quot;中国科学院紫金山天文台&quot;,&quot;中国科学院上海技术物理研究所&quot;,&quot;中国科学院上海天文台&quot;,&quot;中国科学院上海光学精密机械研究所&quot;,&quot;中国科学院上海硅酸盐研究所&quot;,&quot;中国科学院上海微系统与信息技术研究所&quot;,&quot;中国科学院上海有机化学研究所&quot;,&quot;中国科学院上海生命科学研究院&quot;,&quot;中国科学院东北地理与农业生态研究所&quot;,&quot;中国科学院长春应用化学研究所&quot;,&quot;中国科学院长春光学精密机械与物理研究所&quot;,&quot;中国科学院金属研究所&quot;,&quot;中国科学院沈阳应用生态研究所&quot;,&quot;中国科学院沈阳自动化研究所&quot;,&quot;中国科学院沈阳计算技术研究所&quot;,&quot;中国科学院山西煤炭化学研究所&quot;,&quot;中国科学院空间应用工程与技术中心&quot;,&quot;中国科学院光电研究院&quot;,&quot;中国科学院青藏高原研究所&quot;,&quot;中国科学院软件研究所&quot;,&quot;中国科学院计算机网络信息中心&quot;,&quot;中国科学院科技战略咨询研究院&quot;,&quot;中国科学院自然科学史研究所&quot;,&quot;中国科学院微电子研究所&quot;,&quot;中国科学院工程热物理研究所&quot;,&quot;中国科学院自动化研究所&quot;,&quot;中国科学院半导体研究所&quot;,&quot;中国科学院电工研究所&quot;,&quot;中国科学院电子学研究所&quot;,&quot;中国科学院计算技术研究所&quot;,&quot;中国科学院古脊椎动物与古人类研究所&quot;,&quot;中国科学院大气物理研究所&quot;,&quot;中国科学院地质与地球物理研究所&quot;,&quot;中国科学院遗传与发育生物学研究所&quot;,&quot;中国科学院生物物理研究所&quot;,&quot;中国科学院微生物研究所&quot;,&quot;中国科学院植物研究所&quot;,&quot;中国科学院动物研究所&quot;,&quot;中国科学院理化技术研究所&quot;,&quot;中国科学院过程工程研究所&quot;,&quot;中国科学院国家天文台&quot;,&quot;中国科学院声学研究所&quot;,&quot;中国科学院力学研究所&quot;,&quot;中国科学院理论物理研究所&quot;,&quot;中国科学院物理研究所&quot;,&quot;中国科学院化学研究所&quot;,&quot;中国科学院苏州生物医学工程技术研究所&quot;,&quot;中国科学院上海高等研究院&quot;,&quot;中国科学院遥感与数字地球研究所&quot;,&quot;中国科学院地球化学研究所&quot;,&quot;国家纳米科学中心&quot;,&quot;河南理工大学&quot;,&quot;天水师范学院&quot;,&quot;广东技术师范大学&quot;,&quot;杭州电子科技大学&quot;,&quot;安庆师范大学&quot;,&quot;华东交通大学&quot;,&quot;上海出版印刷高等专科学校&quot;,&quot;南京体育学院&quot;,&quot;福州大学&quot;,&quot;中国科学院天津工业生物技术研究所&quot;,&quot;北京生命科学研究所&quot;,&quot;中国科学院高能物理研究所&quot;,&quot;中国科学院空间科学与应用研究中心&quot;,&quot;中国科学院北京基因组研究所&quot;,&quot;中国科学院大连化学物理研究所&quot;,&quot;中国科学院上海应用物理研究所&quot;,&quot;中国科学院海洋研究所&quot;,&quot;中国科学院武汉岩土力学研究所&quot;,&quot;中国科学院北京综合研究中心&quot;,&quot;中国科学院城市环境研究所&quot;,&quot;中国科学院烟台海岸带研究所&quot;,&quot;中国科学院青岛生物能源与过程研究所&quot;,&quot;中国科学院国家天文台乌鲁木齐天文站&quot;,&quot;中国科学院成都山地灾害与环境研究所&quot;,&quot;中国科学院水土保持研究所&quot;,&quot;中国科学院重庆绿色智能技术研究院&quot;,&quot;中国科学院广州化学研究所&quot;,&quot;桂林航天工业学院&quot;,&quot;山西工商学院&quot;,&quot;安徽财贸职业学院&quot;,&quot;内蒙古农业大学&quot;,&quot;青海大学&quot;,&quot;合肥职业技术学院&quot;,&quot;浙江交通职业技术学院&quot;,&quot;合肥幼儿师范高等专科学校&quot;,&quot;西安文理学院&quot;,&quot;中国政法大学&quot;]" />
                        </ul>`
	f := bytes.NewBufferString(raw)
	r := NewRewriteReader(f)
	r.SetTagRewriter("a", "href", nil)
	b, err := ioutil.ReadAll(r)
	fmt.Println(string(b), err)
}

func TestRewriteMetaContent(t *testing.T) {
	var raw = `<meta http-equiv="REFRESH" href="http://www.baidu.com" content="0; URL=https://x.cnki.net/search">
	<meta http-equiv="REFRESH" content="0; url=https://dx.doi.org/10.1108/09685220310468628"/> 
	<meta http-equiv="refresh" content="0;URL=&#39;https://dx.doi.org/10.1108/09685220310468628&#39;"/>    

  <a href="https://www.baidu.com">a</a>
  `
	f := bytes.NewBufferString(raw)
	r := NewRewriteReader(f)
	// urlRewrite := NewURLRewriter("http://dx.doi.org", "iyoerhe.com", "https", true, 1)
	urlRewrite := NewURLRewriterRelativePath("http://dx.doi.org", "iyoerhe.com", "https", true, 1)
	// r.SetTagRewriter(`meta[http-equiv="REFRESH"]`, "href", urlRewrite)
	r.SetTagRewriter(`meta[http-equiv="REFRESH"]`, `content/\d+;\s?url=["']?(.+)["']?/`, urlRewrite)
	r.SetTagRewriter(`meta[http-equiv="refresh"]`, `content/\d+;\s?URL=['"](.+)["']/`, urlRewrite)
	r.SetTagRewriter(`a`, "href", urlRewrite)
	b, err := ioutil.ReadAll(r)
	fmt.Println(string(b), err)
}

func TestRewriteTransform(t *testing.T) {
	var raw = "<div>Hello \xb3\xa3\xd3\xc3\x87\xf8\xd7\xd6\x98\xcb\x9c\xca\xd7\xd6\xf3\x77\xb1\xed</div>"
	f := bytes.NewBufferString(raw)
	r := NewRewriteReader(f, SetTransform(&Transform{
		Encoding: "gbk",
	}))
	urlRewrite := NewURLRewriter("http://x.cnki.net", "iyoerhe.com", "https", true, 1)
	r.SetTagRewriter(`meta[http-equiv="REFRESH"]`, `content/\d+; URL=(.+)/`, urlRewrite)
	r.SetTagRewriter(`a`, "href", urlRewrite)
	b, err := ioutil.ReadAll(r)
	if err != nil {
		t.Error(err)
	}
	if len(raw) != len(b) {
		t.Error("error", string(b))
	}
	// fmt.Println(string(raw))
}

func TestImageData(t *testing.T) {
	// resp, _ := http.Get("http://d.drcnet.com.cn/eDRCnet.common.web/DocDetail.aspx?docid=5764920&leafid=208&chnId=611")
	f, _ := os.Open("./testdata/d.html")
	// bs, _ := ioutil.ReadAll(resp.Body)
	defer f.Close()
	r := NewRewriteReader(f, SetTransform(&Transform{
		// Encoding: "gbk",
	}))
	var b []byte
	// r.SetTagRewriter(`img`, "src", urlRewrite)

	// fe := DetermineEncoding("text/html; charset=gb2312")
	// r := fe.NewDecodeReader(f)
	// b, _ := ioutil.ReadAll(r)
	// r = fe.NewEncodeWriter(bytes.NewReader(b))
	b, _ = ioutil.ReadAll(r)
	// fmt.Println(string(b), err)
	ioutil.WriteFile("./testdata/d3.html", b, os.ModePerm)
}
func TestDubBasetData(t *testing.T) {
	f, _ := os.Open("./testdata/hasBase.html")
	defer f.Close()
	r := NewRewriteReader(f)
	urw := NewURLRewriterRelativePath("https://abc.com/detail/dd", "wf", "https", true, 0)
	r.htmlRewriter = NewHtmlRewriter(nil)
	r.SetTagRewriter(`link`, "href", urw)
	r.SetTagRewriter(`script`, "src", urw)

	r.AddInsertAfter("head", `<base href="http://abc.com/detail/d"></base>`, true)
	var b []byte
	b, _ = ioutil.ReadAll(r)
	fmt.Println(string(b))
	// for key, val := range r.tagTable {
	// 	// fmt.Println(key)
	// 	for _, val := range val {
	// 		fmt.Println(val.matchs[0].rewrite)
	// 	}
	// }

}

func TestIpubXML(t *testing.T) {
	f, _ := os.Open("./testdata/ipub.xml")
	defer f.Close()
	r := NewRewriteReader(f)
	urw := NewURLRewriterRelativePath("http://abc.com/detail/dd", "wf", "https", true, 0)
	r.htmlRewriter = NewHtmlRewriter(nil)
	r.SetTagRewriter(`link`, "href", urw)

	r.AddInsertAfter("head", `<base href="http://abc.com/detail/d"></base>`, true)
	var b []byte
	b, _ = ioutil.ReadAll(r)
	fmt.Println(string(b))
	// for key, val := range r.tagTable {
	// 	// fmt.Println(key)
	// 	for _, val := range val {
	// 		fmt.Println(val.matchs[0].rewrite)
	// 	}
	// }

}
func TestBdCnkiNetData(t *testing.T) {
	f, _ := os.Open("./testdata/bd.cnki.net.html")
	defer f.Close()
	r := NewRewriteReader(f)
	urlRewrite := NewURLRewriter("http://x.cnki.net", "iyoerhe.com", "https", false, 1)
	r.SetTagRewriter(`script`, "src", urlRewrite)

	r.AddInsert("head", `<lin1eak></l1ink>`, true)
	r.AddInsert("head", `<li2nk></li2nk>`, true)
	var b []byte
	b, _ = ioutil.ReadAll(r)
	fmt.Println(string(b))
}
func TestInsertAfterBdCnkiNetData(t *testing.T) {
	f, _ := os.Open("./testdata/bd.cnki.net.html")
	defer f.Close()
	r := NewRewriteReader(f)
	r.AddInsert("head", `<linak></link>`, true)

	r.AddInsertAfter("head", `<liadnk></liacnk>`, true)
	var b []byte
	b, _ = ioutil.ReadAll(r)
	fmt.Println(string(b))
}
func TestChaoxingQikan(t *testing.T) {
	f, _ := os.Open("./testdata/chaoxing_qikan.html")
	defer f.Close()
	r := NewRewriteReader(f)
	urlRewrite := NewURLRewriter("http://x.cnki.net", "iyoerhe.com", "https", true, 1)
	r.SetTagRewriter(`img`, "src", urlRewrite)
	// r.AddInsert("head", `<link></link>`, true)
	var b []byte
	b, _ = ioutil.ReadAll(r)
	fmt.Println(string(b))
}
func TestWwebofknowledgeData(t *testing.T) {
	resp, err := http.Get("http://login.webofknowledge.com/error/Error?Error=IPError&PathInfo=%2F&RouterURL=http%3A%2F%2Fwww.webofknowledge.com%2F&Domain=.webofknowledge.com&Src=IP&Alias=WOK5")
	if err != nil {
		t.Fatal(err)
	}
	var f io.ReadCloser
	f = resp.Body
	defer f.Close()
	r := NewRewriteReader(f, SetTransform(&Transform{
		// Encoding: "gbk",
	}))
	r.SetJavascriptRewriter(NewJavaScriptRewrite())
	var b []byte
	b, _ = ioutil.ReadAll(r)
	ioutil.WriteFile("./testdata/webofknowledge.html", b, os.ModePerm)
}

// http://login.webofknowledge.com/error/Error?Error=IPError&PathInfo=%2F&RouterURL=http%3A%2F%2Fwww.webofknowledge.com%2F&Domain=.webofknowledge.com&Src=IP&Alias=WOK5
var htmlNoChange = `<!DOCTYPE html>
<html>
<head>
  <title></title>
  <head></head>
</head>
<body>
<script type="text/javascript">
    $(function(){
        function anxsGetCookie(name){
            var arr,reg=new RegExp("(^| )"+name+"=([^;]*)(;|$)");
</script>
</body>
</html>`

var basicHTMLRewriteIn = `<!DOCTYPE html>
<html>
<head>
  <title></title>
  <meta></meta>
</head>
<body background="background">
<div><a></a><div id=2><a/></div></div>
  <a href="/apples" nochange="leave/me/alone">link</a>
  <applet codebase="http://a.com/codebase" archive="http://appletarchive.com"></applet>
  <area href="http://a.com/path" />
  <audio src="http://a.com/audio/path" />
  <base href="/b"></base>
  <blockquote cite="http://a.com"></blockquote>
  <button formaction="/">
    <p>no touch</p>
  </button>
  <command icon="word">
    <del cite="citation"></del>
    <embed src="/b/v/c"></embed>
  </command>
  <iframe src="/word"></iframe>
  <image src="huh" xlink:href="im"></image>
  <img src="stuff" srcset="srcset"></img>
  <ins cite="/1234"></ins>
  <input src="/huh" formaction="/word"></input>
  <form action="/form/action">
    <frame src="/frame/src"></frame>
  </form>
  <link href="/link"></link>
  <script src="http://a.com:8080/stuff/static/js/script.js">afadfaf</script>
  <script src="/static/js/script.js"></script>
  <source src="/turn/left"></source>
  <video src="/yep" poster="poster"></video>
  <h3 href="http://a.com/stuff"></h3>
</body>
</html>`

var basicHTMLRewriteOut = `<!DOCTYPE html>
<html>
<head>
  <title></title>
  <meta></meta>
</head>
<body background="im_background">
  <a href="https://b.tv/apples" nochange="leave/me/alone">link</a>
  <applet codebase="oe_http://a.com/codebase" archive="oe_http://appletarchive.com"></applet>
  <area href="https://b.tv/path"/>
  <audio src="oe_http://a.com/audio/path"/>
  <base href="https://b.tv/b"></base>
  <blockquote cite="https://b.tv"></blockquote>
  <button formaction="https://b.tv/">
    <p>no touch</p>
  </button>
  <command icon="im_word">
    <del cite="https://b.tv/citation"></del>
    <embed src="oe_/b/v/c"></embed>
  </command>
  <iframe src="if_/word"></iframe>
  <image src="im_huh" xlink:href="im_im"></image>
  <img src="im_stuff" srcset="im_srcset"></img>
  <ins cite="https://b.tv/1234"></ins>
  <input src="im_/huh" formaction="https://b.tv/word"></input>
  <form action="https://b.tv/form/action">
    <frame src="fr_/frame/src"></frame>
  </form>
  <link href="oe_/link"></link>
  <script src="https://b.tv/stuff/static/js/script.js"></script>
  <script src="/static/js/script.js">adadfaf</script>
  <source src="oe_/turn/left"></source>
  <video src="oe_/yep" poster="im_poster"></video>
  <h3 href="http://a.com/stuff"></h3>
</body>
</html>`

func Test(t *testing.T) {
	var testdata = "asdf2019年阿斯顿发顺丰大叔分10点， 2月1日asdf, 5月6日, 2019-04-05"
	var stack []rune
	// var start int
	// var end int
	for _, ch := range testdata {
		if IsDig(ch) {
			stack = append(stack, ch)
			continue
		}
		// fmt.Println(string(stack))
		stack = []rune{}
	}
	r := regexp.MustCompile("(\\d{2,4}[年|-])?\\d{1,2}[月|-](\\d{1,2}日?)?")

	fmt.Println(r.FindAllIndex([]byte(testdata), -1))
	fmt.Println(string(stack))

}

func IsDig(char rune) bool {
	return char > 47 && char < 58
}

func Test_readRegexStrEnd(t *testing.T) {
	// str := "aaa/\\/aa\\//"
	// i := readRegexStrEnd(str)
	// fmt.Println(i, string(str[i]))
	// attr := "content/(script-src[^;]+;)/********/"
	// n, r, f, e := readAttrName(attr)
	// fmt.Println(n, r, f, e)
	// fmt.Println(string(f.Rewrite([]byte("default-src *; script-src &#39;unsafe-inline&#39; &#39;unsafe-eval&#39; &#39;self&#39; *.zhangyue.com *.ireader.com *.zhangyue01.com *.163yun.com *.163.com localhost *.126.net *.126.com *.netease.com *.qq.com *.gtimg.cn *.baidu.com *.bdstatic.com *.hicloud.com *.baidustatic.com; style-src * &#39;unsafe-inline&#39;; img-src * data: ; frame-src &#39;self&#39; *.zhangyue.com *.zhangyue01.com *.ireader.com *.alipay.com *.gtimg.cn *.qq.com *.baidu.com zhangyueireader: weixin:;"))))

	var raw = `
	<meta http-equiv="Content-Security-Policy" content="default-src *; script-src 'unsafe-inline' 'unsafe-eval' 'self' *.zhangyue.com *.ireader.com *.zhangyue01.com *.163yun.com *.163.com localhost *.126.net *.126.com *.netease.com *.qq.com *.gtimg.cn *.baidu.com *.bdstatic.com *.hicloud.com *.baidustatic.com; style-src * 'unsafe-inline'; img-src * data: ; frame-src 'self' *.zhangyue.com *.zhangyue01.com *.ireader.com *.alipay.com *.gtimg.cn *.qq.com *.baidu.com zhangyueireader: weixin:;"/>
  `
	f := bytes.NewBufferString(raw)
	r := NewRewriteReader(f)
	// urlRewrite := NewURLRewriter("http://dx.doi.org", "iyoerhe.com", "https", true, 1)
	urlRewrite := NewURLRewriterRelativePath("http://dx.doi.org", "iyoerhe.com", "http", true, 0)
	// r.SetTagRewriter(`meta[http-equiv="REFRESH"]`, "href", urlRewrite)
	r.SetTagRewriter(`meta[http-equiv="Content-Security-Policy"]`, `content/(script-src [^;]+)/script-src 'unsafe-inline' 'unsafe-eval' 'self' */`, urlRewrite)
	r.SetTagRewriter(`meta[http-equiv="Content-Security-Policy"]`, `content/(frame-src [^;]+)/frame-src 'self' */`, urlRewrite)
	// r.SetTagRewriter(`meta[http-equiv="Content-Security-Policy"]`, `href`, urlRewrite)
	// r.SetTagRewriter(`meta[http-equiv="refresh"]`, `href`, urlRewrite)
	// r.SetTagRewriter(`meta[http-equiv="refresh"]`, "href/(.*)/$1--ggg/", urlRewrite)
	b, err := ioutil.ReadAll(r)
	fmt.Println(string(b), err)
}

// <meta http-equiv="Content-Security-Policy" content="default-src *; script-src 'unsafe-inline' 'unsafe-eval' 'self' *.zhangyue.com *.ireader.com *.zhangyue01.com *.163yun.com *.163.com localhost *.126.net *.126.com *.netease.com *.qq.com *.gtimg.cn *.baidu.com *.bdstatic.com *.hicloud.com *.baidustatic.com; style-src * 'unsafe-inline'; img-src * data: ; frame-src 'self' *.zhangyue.com *.zhangyue01.com *.ireader.com *.alipay.com *.gtimg.cn *.qq.com *.baidu.com zhangyueireader: weixin:;"/>
