package rewrite

import (
	"bufio"
	"bytes"
	"fmt"
	"io"

	"gitlab.iyorhe.com/dr2am/jsast"
)

type JavaScriptRewrite struct {
	err       error
	buf       *bufio.Reader
	walker    io.Reader
	transform Transform
}

func NewJavaScriptRewrite() *JavaScriptRewrite {
	return &JavaScriptRewrite{}
}

func (jr *JavaScriptRewrite) NewReader(r io.Reader) io.Reader {
	// bs, err := ioutil.ReadAll(r)
	// if err != nil {
	// 	jr.err = err
	// }
	// jr.buf = bs
	jr.buf = bufio.NewReader(r)

	// jr.program, jr.err = parser.ParseFile(nil, "", jr.buf, 0)
	// _, err := io.Copy(jr.buf, r)
	// fmt.Println(string(jr.buf.String()))
	// fmt.Println(err)
	// if err != nil {
	// 	jr.err = err
	// }
	jr.walker, jr.err = jsast.NewWalker(jr.buf)
	return jr
}
func (jr *JavaScriptRewrite) Read(p []byte) (n int, err error) {
	if jr.err != nil {
		fmt.Println(jr.err)
		// n, _ := jr.buf.Read(p)
		return jr.buf.Read(p)
	}
	// jsast.
	return jr.walker.Read(p)
}

func (jr *JavaScriptRewrite) Rewrite(i []byte) (o []byte) {
	jsr := NewJavaScriptRewrite()
	jsr.NewReader(bytes.NewBuffer(i))
	w := bytes.NewBuffer(nil)

	_, err := io.Copy(w, jsr)
	if err != nil {
		return i
	}
	return w.Bytes()
}
