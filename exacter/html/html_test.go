package html

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

func TestHTMLExacter_NewReader(t *testing.T) {
	fi, err := os.Open("./testdata/brief-cnki.aspx")
	if err != nil {
		t.Fatal(err)
	}
	defer fi.Close()
	tagsTable := NewTagsTable()
	tagsTable.AddTagQuery(`table[class="GridTableContent"]`, "")
	// fmt.Println(tagsTable)
	h, err := NewHTMLExacter(fi, tagsTable)
	if err != nil {
		t.Fatal(err)
	}
	h.On = (&Table{}).Parse

	_, err = ioutil.ReadAll(h)
	if err != nil {
		t.Fatal(err)
	}
	// fmt.Println(string(b))
}
func TestGetExactor(t *testing.T) {
	fi, err := os.Open("./testdata/detail-cnki-qikan.aspx")
	if err != nil {
		t.Fatal(err)
	}
	defer fi.Close()
	col1 := Command{
		Command:          "get",
		ExpressionString: "xpath=//h2",
		Value:            "title",
	}
	col2 := Command{
		Command:          "get",
		ExpressionString: `xpath=//*[@id="ChDivSummary"]`,
		Value:            "abs",
	}
	col3 := Command{
		Command:          "getAll",
		ExpressionString: `xpath=//*[@class="author"]/span`,
		Value:            "author",
	}
	col4 := Command{
		Command:          "get",
		ExpressionString: `xpath=//*[@id="catalog_ZTCLS"]/../text()`,
		Value:            "clcs",
	}

	col5 := Command{
		Command:          "getAll",
		ExpressionString: `xpath=//*[@id="catalog_KEYWORD"]/../a`,
		Value:            "keywords",
	}

	col6 := Command{
		Command:          "get",
		ExpressionString: `xpath=//*[@id="catalog_ZCDOI"]/../text()`,
		Value:            "doi",
	}
	col7 := Command{
		Command:          "getAll",
		ExpressionString: `xpath=//*[@class="orgn"]//a`,
		Value:            "orgn",
	}
	col8 := Command{
		Command:          "get",
		ExpressionString: `set=detail`,
		Value:            "action",
	}

	col10 := Command{
		Command:          "getAll",
		ExpressionString: `xpath=//*[@id="DownLoadParts"]/a/@href`,
		Value:            "download_url",
	}

	col11 := Command{
		Command:          "get",
		ExpressionString: `xpath=//div[@class="sourinfo"]/p[@class="title"]`,
		Value:            "source",
	}

	// col12 := Command{
	// 	Command:          "get",
	// 	ExpressionString: `xpath=^//*[@id="catalog_Ptitle"]`,
	// 	Value:            "document_type",
	// }
	command := Command{
		Command: "get",

		ExpressionString: `xpath=//*[@id="mainArea"]/div[@class="wxmain"]`,

		Value: "result",
		NestingCommands: Commands{
			col1, col2, col3, col4, col5, col6, col7, col8, col10, col11,
		},
	}

	command2 := Command{
		Command:          "get",
		ExpressionString: `xpath=//*`,
		// ExpressionString: `xpath=//a[contains(@onclick,"TurnPageToKnet('au'")]`,
		NestingCommands: Commands{
			col8, command,
		},
	}
	command3 := Command{
		Command:          "condition",
		ExpressionString: `xpath=//*`,
		// ExpressionString: `xpath=//a[contains(@onclick,"TurnPageToKnet('au'")]`,
		NestingCommands: Commands{
			command2,
		},
	}
	bs, err := json.Marshal(command3)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(bs))
	h, err := NewHTMLExacter(fi, nil)
	if err != nil {
		t.Fatal(err)
	}
	h.AddCommands(command3)
	result, err := h.Exec()
	if err != nil {
		t.Fatal(err)
	}
	printJSON(result)
	// fmt.Printf("%#v", result)
}

func printJSON(src interface{}) {
	bs, _ := json.Marshal(src)
	fmt.Println(string(bs))
}

func TestGetAllExactor(t *testing.T) {
	fi, err := os.Open("./testdata/brief-cnki.aspx")
	if err != nil {
		t.Fatal(err)
	}
	defer fi.Close()
	col1 := Command{
		Command:          "get",
		ExpressionString: "xpath=//a[@class='fz14']",
		Value:            "title",
	}

	col2 := Command{
		Command:          "get",
		ExpressionString: "xpath=//a[@class='fz14']/@href",
		Value:            "title_url",
	}

	col3 := Command{
		Command:          "get",
		ExpressionString: "xpath=//a[@class='briefDl_D']/@href",
		Value:            "download_url",
	}
	col4 := Command{
		Command:          "getAll",
		ExpressionString: "xpath=//a[@class='KnowledgeNetLink']",
		Value:            "author",
	}
	col5 := Command{
		Command:          "getAll",
		ExpressionString: "xpath=//td[4]/a",
		Value:            "source",
	}
	col6 := Command{
		Command:          "getAll",
		ExpressionString: "xpath=//td[5]",
		Value:            "date",
	}
	col7 := Command{
		Command:          "get",
		ExpressionString: "set=search",
		Value:            "action",
	}
	mapCommand := Command{
		Command:          "map",
		ExpressionString: `{{.download_url.URL.Query.Get "filename"}}`,
		Value:            "id",
	}
	command := Command{
		Command:          "getAll",
		ExpressionString: "xpath=//table[@class='GridTableContent']/tbody/tr[1]/following-sibling::tr",
		Value:            "results",
		NestingCommands: Commands{
			col1, col2, col3, col4, col5, col6, mapCommand,
		},
	}

	command2 := Command{
		Command:          "get",
		ExpressionString: "xpath=/html/body",

		NestingCommands: Commands{
			col7, command, //mapCommand,
		},
	}
	bs, err := json.Marshal(command2)
	if err != nil {
		t.Fatal(err)
	}
	var jcommand Command
	json.Unmarshal(bs, &jcommand)
	h, err := NewHTMLExacter(fi, nil)
	if err != nil {
		t.Fatal(err)
	}
	printJSON(jcommand)
	// fmt.Println(string(bs))
	h.AddCommands(jcommand)
	result, err := h.Exec()
	if err != nil {
		t.Fatal(err)
	}
	printJSON(result)
	r := result.([]interface{})
	fmt.Println(r)
	//commands

}

func TestConditionCommand(t *testing.T) {
	fi, err := os.Open("./testdata/detail-cnki-qikan.aspx")
	if err != nil {
		t.Fatal(err)
	}
	defer fi.Close()
	h, err := NewHTMLExacter(fi, nil)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("GET", "https://www.cnki.net/index.html?field=a&field=b", nil)
	if err != nil {
		t.Fatal(err)
	}
	h.AddRequestContext(req)

	command := NewCommand("condition", "when request.url.hostname:'www.cnki.net' ", "", NewCommand("get", "xpath=//h2", "title"))
	h.AddCommands(command)

	result, err := h.Exec()
	if err != nil {
		t.Fatal(err)
	}
	r := result.([]interface{})
	if !reflect.DeepEqual(r[0], true) {
		t.Error(r)
	}
}

func TestConditionParse(t *testing.T) {
	ctx := context.Background()
	command := Command{
		Command:          "condition",
		ExpressionString: "?domain:xxx",
	}
	_, expr, err := command.GetExpr()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(expr.Exec(ctx))
}

func TestXPathRootExpr(t *testing.T) {
	fi, err := os.Open("./testdata/detail-degree.do")
	if err != nil {
		t.Fatal(err)
	}
	defer fi.Close()
	title := Command{
		Command:          "get",
		ExpressionString: "xpath=^//*[@id=\"div_a\"]/div/div[1]/a[2]",
		Value:            "head_titlea",
	}
	head_title := Command{
		Command:          "get",
		ExpressionString: "xpath=^/html/head/title",
		Value:            "head_title",
	}
	resultC := Command{
		Value:            "results",
		Command:          "get",
		ExpressionString: "xpath=//*[@id=\"div_a\"]",
		NestingCommands: Commands{
			title, head_title,
		},
	}
	h, err := NewHTMLExacter(fi, nil)
	if err != nil {
		t.Fatal(err)
	}
	h.AddCommands(resultC)
	result, err := h.Exec()
	if err != nil {
		t.Fatal(err)
	}
	printJSON(result)
}

func TestWFSearchList(t *testing.T) {
	fi, err := os.Open("./testdata/searchList-wf.do")
	if err != nil {
		t.Fatal(err)
	}
	defer fi.Close()
	title := Command{
		Command:          "get",
		ExpressionString: `xpath=//div[@class="title"]/a`,
		Value:            "head_titlea",
	}
	head_title := Command{
		Command:          "get",
		ExpressionString: `xpath=//div[@class="title"]/a/@href`,
		Value:            "head_title",
	}
	resultC := Command{
		Value:            "results",
		Command:          "getAll",
		ExpressionString: `xpath=//div[contains(@class,"ResultBlock")]/div`,
		NestingCommands: Commands{
			title, head_title,
		},
	}
	h, err := NewHTMLExacter(fi, nil)
	if err != nil {
		t.Fatal(err)
	}
	h.AddCommands(resultC)
	result, err := h.Exec()
	if err != nil {
		t.Fatal(err)
	}
	printJSON(result)
}

func TestWFDetailJSON(t *testing.T) {
	fi, err := os.Open("./testdata/detail_wf_2.json")
	if err != nil {
		t.Fatal(err)
	}
	defer fi.Close()
	title := Command{
		Command: "map",
		// ExpressionString: `{{(index .detail 0).periodical.DOI}}`,
		ExpressionString: `map detail[0][type].Abstract[0]`,
		// ExpressionString: `{{ 0 | index (.detail.GetChild 0).periodical.Title  }}`,
		// ExpressionString: `{{ 1 | index (.detail.GetChild 0).periodical.Title }}`,
		// ExpressionString: `{{((.detail.GetChild 0).periodical.Title).GetChild 0}}`,
		Value: "title",
	}
	author := Command{
		Command: "map",
		// ExpressionString: `{{(index .detail 0).periodical.DOI}}`,
		ExpressionString: `map detail[0][type].PeriodicalTitle[0]`,
		// ExpressionString: `{{ 1 | index (.detail.GetChild 0).periodical.Title }}`,
		// ExpressionString: `{{((.detail.GetChild 0).periodical.Title).GetChild 0}}`,
		Value: "author",
	}
	titleen := Command{
		Command: "map",
		// ExpressionString: `{{(index .detail 0).periodical.DOI}}`,
		ExpressionString: `map detail[0][type].Title[0]`,
		// ExpressionString: `{{ 1 | index (.detail.GetChild 0).periodical.Title }}`,
		// ExpressionString: `{{((.detail.GetChild 0).periodical.Title).GetChild 0}}`,
		Value: "title_en",
	}
	set := Command{
		Command:          "get",
		ExpressionString: `set=periodical`,
		Value:            "type",
	}
	condition2 := Command{
		Command:          "condition",
		ExpressionString: `when request.url.path=~/\/Detail\/Periodical\/.+/`,
		NestingCommands: Commands{
			set,
		},
	}
	condition3 := Command{
		Command:          "condition",
		ExpressionString: `when request.url.hostname:'d.wanfangdata.com.cn2'`,
		NestingCommands: Commands{
			Command{
				Command:          "get",
				ExpressionString: `set=periodica2l`,
				Value:            "type",
			},
		},
	}
	mapdetail := Command{
		Command:          "map",
		ExpressionString: `{{.request.URL.Query.Get "searchKeywords"}}`,
		Value:            "mapdetail",
	}
	resultC := Command{
		Value:            "results",
		Command:          "get",
		ExpressionString: "json=",
		// ExpressionString: `xpath=//div[contains(@class,"ResultBlock")]/div`,
		NestingCommands: Commands{
			condition3, condition2, mapdetail,

			Command{
				Command:          "condition",
				ExpressionString: `when type:'periodicala'`,
				NestingCommands: Commands{
					Command{
						Command:          "map",
						ExpressionString: `{{.type}}/aaa`,
						Value:            "mapdetail",
					},
				},
			},
			title, titleen, author, //head_title,
		},
	}

	condition := Command{
		Command:          "condition",
		ExpressionString: "when request.url.hostname:'d.wanfangdata.com.cn' AND (request.url.path=~/\\/Detail\\/Periodical\\/.+/ OR request.url.path=~/\\/Detail\\/Thesis\\/.+/)",
		NestingCommands: Commands{
			resultC, //head_title,
		},
	}
	h, err := NewHTMLExacter(fi, nil)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", "http://d.wanfangdata.com.cn/Detail/Periodical/syyxzz202006016?searchKeywords=aa+b", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	resp.HeaderMap.Add("Content-Type", "application/json")
	h.AddRequestContext(req)
	// h.AddResponseContext(resp.Result())

	h.AddCommands(NewCommand("condition", "when request.url.hostname:'g.wanfangdata.com.cn'", ""), condition)
	result, err := h.Exec()

	printJSON(result)
	fmt.Println(err)
	if err != nil {
		t.Fatal(err)
	}
}

func TestTTTTT(t *testing.T) {
	fi, err := os.Open("./testdata/detail_wf_2.json")
	if err != nil {
		t.Fatal(err)
	}
	defer fi.Close()
	req, err := http.NewRequest("GET", "http://d.wanfangdata.com.cn/Detail/Periodical/syyxzz202006016", nil)
	if err != nil {
		t.Fatal(err)
	}

	result, err := newHTMLExactorWithScript(req, wfcommands, fi)
	printJSON(result)
	fmt.Println(err)
}

var wfcommands = `[{"command":"condition","id":"4d64a71d-b6e6-41ad-b3da-63cb3cadcf97","expression":"when request.url.hostname:\"www.wanfangdata.com.cn\" AND request.url.path:\"/details/detail.do\"","nesting_commands":[{"command":"get","id":"4b422159-530e-43df-9acf-9fb9d7092579","expression":"set=detail","value":"action"},{"command":"get","id":"bc216813-ad4e-4c12-b525-c6be05e73fa7","expression":"xpath=//*[@id=\"div_a\"]","nesting_commands":[{"command":"get","id":"1387c9ea-c464-4c52-964a-009d12f37be2","expression":"xpath=^/html/head/title","value":"title"},{"command":"get","id":"164f0565-b3f5-479b-9472-6ae2a80ba39b","expression":"xpath=//div[@class=\"English\"]","value":"title_en"},{"command":"get","id":"4c4f94e3-6355-4a30-a239-68d67b61805e","expression":"xpath=//*[@class=\"abstract\"]/textarea","value":"abs"},{"command":"getAll","id":"9929f7c5-ca2a-4675-b4dc-205e5dae9421","expression":"xpath=//ul[@class=\"info\"]/li/div[1 and text()=\"作者：\"]/following-sibling::*/a","value":"author"},{"command":"getAll","id":"148a4c28-4bfd-4c6f-ad46-f291eb10cf3e","expression":"xpath=//ul[@class=\"info\"]/li/div[1 and text()=\"关键词：\"]/following-sibling::*/a","value":"keywords"},{"command":"get","id":"40965706-cfd2-401d-80db-d0213e96f859","expression":"xpath=//ul[@class=\"info\"]/li/div[1 and text()=\"doi：\"]/following-sibling::*/a","value":"doi"},{"command":"get","id":"4445fb49-213d-4f0d-8e43-f640461da6ec","expression":"xpath=//ul[@class=\"info\"]/li/div[1 and text()=\"刊名：\"]/following-sibling::*/a","value":"source"}],"value":"results"}]},{"command":"condition","id":"e3575ec9-adc9-40e1-bd18-3a117a807e9e","expression":"when request.url.hostname:\"www.wanfangdata.com.cn\" AND request.url.path:\"/search/searchList.do\"","nesting_commands":[{"command":"get","id":"af883f0f-1bf0-4257-b4fd-5453b4df7c67","expression":"set=search","value":"action"},{"command":"getAll","id":"b70395d9-9ee7-43e1-99ef-486ee6833732","expression":"xpath=//div[@class=\"ResultBlock\"]/div","value":"results","nesting_commands":[{"command":"get","id":"d8f95080-13a4-46ff-b2aa-9718ad44b14c","expression":"xpath=//div[@class=\"title\"]/a","value":"title"},{"command":"get","id":"23979d0b-e7e3-4c66-8f4c-41f64620d7f6","expression":"xpath=//div[@class=\"title\"]/a/@href","value":"title_url"},{"command":"get","id":"9281ba67-7ae6-4366-b0eb-610fa11b134e","expression":"xpath=//div[@class=\"ResultCheck\"]/input/@docid","value":"id"},{"command":"get","id":"75fc0003-7544-49f5-9f73-b81164196b08","expression":"xpath=//div[@class=\"ResultCheck\"]/input/@doctype","value":"type"},{"command":"map","id":"d10cb9cd-2ba5-4d56-acd8-8495ec1c6935","expression":"www.wanfangdata.com.cn/{{.type}}/{{.id}}","value":"docid"},{"command":"map","id":"e5143df8-0022-4fc2-9ee7-ec62e933b10b","expression":"{{.type}}/{{.id}}","value":"id"}]}]},{"command":"condition","id":"4b108348-db85-4ece-8971-732e1e02adb9","expression":"when request.url.hostname:'d.wanfangdata.com.cn' AND (request.url.path=~/\\/Detail\\/Periodical\\/.+/ OR request.url.path=~/\\/Detail\\/Thesis\\/.+/)","nesting_commands":[{"command":"get","id":"2ff3d7b0-adbf-4718-83ac-97ce9d9c1ae6","expression":"set=detail","value":"action"},{"command":"condition","id":"67fd8459-545d-46a3-86e2-45e8db1a3d49","expression":"when request.url.path=~/\\/Detail\\/Periodical\\/.+/","nesting_commands":[{"command":"get","id":"de3b4c2e-f625-48e0-be9d-ef8cb769a6e0","expression":"set=periodical","value":"type"}]},{"command":"condition","id":"ce872a51-011e-498a-9d34-5e7235af09d8","expression":"when request.url.path=~/\\/Detail\\/Thesis\\/.+/","nesting_commands":[{"command":"get","id":"2cd9a335-9c4a-4d5f-9eb4-7f03c14ee920","expression":"set=thesis","value":"type"}]},{"command":"get","id":"d6ff7900-a1a1-49ad-ae37-cfb133398b82","expression":"json=","nesting_commands":[{"command":"condition","id":"f68c78ab-c397-44b8-867c-fac7d76bfda0","expression":"when request.url.path=~/\\/Detail\\/Periodical\\/.+/","value":"","nesting_commands":[{"command":"get","id":"20f96f22-e1a2-411c-ae0d-9ef8fab3e91a","expression":"set=periodical","value":"type"}]},{"command":"condition","id":"209eab2a-0e85-4e5e-95ed-5e1a57abf114","expression":"when request.url.path=~/\\/Detail\\/Thesis\\/.+/","nesting_commands":[{"command":"get","id":"71731e85-708d-44d4-ac76-1837cadfefa4","expression":"set=thesis","value":"type"}]},{"command":"map","id":"f5730426-ffad-4eb3-ac92-50ae3e62f9f1","expression":"map detail[0][type].Title[0]","value":"title"},{"command":"map","id":"de468b37-e5b2-470b-be1d-099c1e07b026","expression":"map detail[0][type].Title[1]","value":"title_en"},{"command":"map","id":"59a63fab-1e3a-413e-a64d-def882364c70","expression":"map detail[0][type].Abstract[0]","value":"abs"},{"command":"map","id":"2bd1b582-2ae2-4b77-aa1c-694c9c368842","expression":"map detail[0][type].Abstract[1]","value":"abs_en"},{"command":"map","id":"5246b108-d267-4cac-a3ab-d1ee868b08c2","expression":"map detail[0][type].Creator","value":"author"},{"command":"map","id":"584b5a7d-24cd-41b9-b0ab-696edd46ae19","expression":"map detail[0][type].OriginalOrganization","value":"organization"},{"command":"map","id":"828a778b-236e-4f36-a1e6-c8d484d9aaf4","expression":"map detail[0][type].DOI","value":"DOI"},{"command":"map","id":"54191b7e-96b0-4ebf-8e64-d8fd24580cd5","expression":"map detail[0][type].Keywords","value":"keywords"},{"command":"map","id":"e591487d-c3d2-4432-a3d2-55e7269ed1e9","expression":"map detail[0][type].PeriodicalTitle[0]","value":"periodical"},{"command":"map","id":"03411190-08ec-471d-adc7-8d77299f2978","expression":"map detail[0][type].PeriodicalTitle[1]","value":"periodical_en"},{"command":"map","id":"7d9ef501-339b-4e50-8edc-a649196c2a4b","expression":"map detail[0][type].PublishYear","value":"publishYear"},{"command":"map","id":"5c951f1a-9bbe-474e-9210-05798eb029d9","expression":"map detail[0][type].Volum","value":"volum"},{"command":"map","id":"4f4f44c5-975c-4326-998e-7ceae35dabe0","expression":"map detail[0][type].Issue","value":"issue"},{"command":"map","id":"ee0380dd-deef-439d-b6c2-226e6585f6f2","expression":"map detail[0][type].ClassCodeForSearch","value":"clc"},{"command":"map","id":"2d729763-5796-47b5-8807-8a55f9fa7c3a","expression":"map detail[0][type].ISSN","value":"issn"},{"command":"map","id":"d307517c-9a5b-4580-aaee-c811ecdec915","expression":"map detail[0][type].Page","value":"page"},{"command":"map","id":"02d35ce3-3c67-4522-bab7-fe2f988666fc","expression":"map detail[0][type].PageNo","value":"pageNum"},{"command":"map","id":"a5c46491-1ee4-44d4-8f9b-e5aa4aa6fc87","expression":"map detail[0][type].Language","value":"language"}],"value":"results"}]}]`
