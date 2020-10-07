package rewrite

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

const noChangeCss = `
  html, head, body{
    color: black;
    background: black;
    margin: 0;
    padding: 0;
  }
`

func TestCssRewriter(t *testing.T) {
	// urlrw := NewUrlRewriter("http://a.com", "https://b.tv")
	urlrw := NewURLRewriterRelativePath("https://libcdn.fifedu.com/common/css/public.css", "dr2am.cn", "https", true, 0)
	f, _ := os.Open("./testdata/css/sslibrary_base.css")
	defer f.Close()
	rw := NewNewCssRewriterReader(f, urlrw)
	// buf :=bytes.NewBuffer(nil)
	bs, err := ioutil.ReadAll(rw)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(bs))
}
