package html

import (
	"fmt"
	"testing"

	"gopkg.in/xmlpath.v2"
)

func TestXPathExpr_ExecReturnIter(t *testing.T) {
	path, err := xmlpath.Compile(`//div[@class="ResultBlock"]/*[contains(@class,"ResultList")]/self::*[2]`)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(path)
}
