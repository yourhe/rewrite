package html

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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
	fi, err := os.Open("./testdata/detail-cnki.aspx")
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
	command := Command{
		Command: "get",

		ExpressionString: `xpath=//*[@id="mainArea"]/div[3]`,

		// Value: "detail",
		NestingCommands: Commands{
			col1, col2, col3, col4, col5,
		},
	}

	// command := Command{
	// 	Command:          "get",
	// 	ExpressionString: `xpath=/html/head/meta/@content`,
	// 	// ExpressionString: `xpath=//a[contains(@onclick,"TurnPageToKnet('au'")]`,
	// 	Value: "detail",
	// }

	h, err := NewHTMLExacter(fi, nil)
	if err != nil {
		t.Fatal(err)
	}
	h.AddCommands(command)
	result, err := h.Exec()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
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
	command := Command{
		Command:          "getAll",
		ExpressionString: "xpath=//table[@class='GridTableContent']/tbody/tr[1]/following-sibling::*",
		// Value:            "col1",
		NestingCommands: Commands{
			col1, col2, col3, col4, col5, col6,
		},
	}
	bs, err := json.Marshal(command)
	if err != nil {
		t.Fatal(err)
	}
	var jcommand Command
	json.Unmarshal(bs, &jcommand)
	h, err := NewHTMLExacter(fi, nil)
	if err != nil {
		t.Fatal(err)
	}
	h.AddCommands(jcommand)
	result, err := h.Exec()
	if err != nil {
		t.Fatal(err)
	}

	r := result.([]interface{})
	//commands
	for _, v := range r {

		// rows
		r := v.([]interface{})

		for _, v := range r {
			fmt.Println(v)
			// 	//column
			// 	r := v.(map[string]interface{})
			// 	for k, v := range r {
			// 		//column
			// 		fmt.Println(k, v.(string))

			// 	}
			// 	fmt.Println("*************")
		}
	}

}
