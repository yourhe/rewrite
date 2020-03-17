package rewrite

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestJavaScriptRewrite_Read(t *testing.T) {
	rw := NewJavaScriptRewrite()
	r := rw.NewReader(bytes.NewBufferString("a = window.location.hash"))
	bf := bytes.NewBuffer(nil)
	io.Copy(bf, r)
	fmt.Println(bf.String())
}
