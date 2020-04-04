package html

import (
	"bytes"
	"context"
	"io"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"gopkg.in/xmlpath.v2"
)

type HTMLExacter struct {
	tokenizer *html.Tokenizer
	buf, buf1 *bytes.Buffer
	matcher   []int
	eof       bool
	On        func([]byte) error
	tagstable TagsTable
	ctx       context.Context
	commands  Commands
}

func NewHTMLExacter(r io.Reader, tagsTable TagsTable) (*HTMLExacter, error) {
	h := &HTMLExacter{
		tagstable: tagsTable,
		ctx:       context.Background(),
	}
	err := h.parseXPath(r)
	if err != nil {
		return nil, err
	}
	// h.NewReader(r)
	return h, nil
}

func (h *HTMLExacter) parseXPath(r io.Reader) error {
	root, err := xmlpath.ParseHTML(r)
	if err != nil {
		return err
	}
	h.ctx = context.WithValue(h.ctx, RootNodeKey, root)
	// fmt.Println(ctx)

	return nil
}

func (h *HTMLExacter) AddCommands(command ...Command) {
	h.commands = append(h.commands, command...)
}

func (h *HTMLExacter) Exec() (interface{}, error) {
	return h.commands.Exec(h.ctx)

}

func (h *HTMLExacter) NewReader(r io.Reader) io.Reader {
	if r == nil {
		return h
	}
	h.Reset()
	h.tokenizer = html.NewTokenizer(r)
	return h
}

func (h *HTMLExacter) Reset() {
	h.buf = bytes.NewBuffer(nil)
	h.buf1 = bytes.NewBuffer(nil)
}

func (r *HTMLExacter) read(p []byte) (n int, err error) {
	return r.buf.Read(p)
}

func (r *HTMLExacter) write(p []byte) (n int, err error) {
	return r.buf.Write(p)
}

func (r *HTMLExacter) Read(p []byte) (n int, err error) {
	tokenizer := r.tokenizer

	tt := tokenizer.Next()

	// var raw []byte
	switch tt {
	case html.ErrorToken:

		r.eof = true
		return r.read(p)
	case html.StartTagToken:
		r.write(tokenizer.Raw())
		raw, matched, dataAtom := r.IsMatch(tokenizer)
		if matched {
			r.buf1.Write(raw)
			data := r.readCloseNode(dataAtom)
			if r.On != nil {
				r.On(data)
			}
		}

	case html.SelfClosingTagToken:
		r.write(tokenizer.Raw())
		raw, matched, _ := r.IsMatch(tokenizer)
		if matched && r.On != nil {
			r.On(raw)
		}

	default:
		r.write(tokenizer.Raw())
	}

	return r.read(p)
}

func (r *HTMLExacter) readCloseNode(dataAtom atom.Atom) []byte {
	defer r.buf1.Reset()
	left := 1
	tokenizer := r.tokenizer

loop:
	tt := tokenizer.Next()
	raw := make([]byte, len(tokenizer.Raw()))
	copy(raw, tokenizer.Raw())
	// r.write(raw)
	r.buf1.Write(raw)
	token := tokenizer.Token()
	matchAtom := token.DataAtom == dataAtom
	switch tt {
	case html.ErrorToken:
		r.eof = true
		return r.buf1.Bytes()
	case html.StartTagToken, html.SelfClosingTagToken:
		if matchAtom {
			left++
		}
	case html.EndTagToken:

		if matchAtom {
			left--
		}

		if left == 0 {
			return r.buf1.Bytes()
		}
		// default:
		// 	r.buf1.Write(raw)
	}
	goto loop
}

func (r *HTMLExacter) IsMatch(t *html.Tokenizer) ([]byte, bool, atom.Atom) {

	// var pattam = atom.Table
	raw := make([]byte, len(t.Raw()))
	copy(raw, t.Raw())
	token := t.Token()
	// var isMatch bool
	isMatch, _ := r.tagstable.queryTags(&token)
	return raw, isMatch, token.DataAtom
}

// func (r *HTMLExacter) Match(t *html.Tokenizer, tp html.TokenType) ([]byte, bool) {

// 	// var pattam = atom.Table
// 	raw := make([]byte, len(t.Raw()))
// 	copy(raw, t.Raw())
// 	token := t.Token()

// 	// var isMatch bool
// 	isMatch, dataAtom := r.tagstable.queryTags(&token)
// 	return raw, isMatch
// 	// fmt.Println(isMatch, dataAtom)
// 	// switch token.DataAtom {
// 	// case pattam:
// 	// 	isMatch = true
// 	// default:
// 	// }

// 	if !isMatch && len(r.matcher) == 0 {
// 		return
// 	}
// 	// fmt.Println(string(raw))
// 	start := r.buf1.Len()
// 	last := len(r.matcher) - 1
// 	switch tp {
// 	case html.StartTagToken:
// 		r.buf1.Write(raw)
// 		if isMatch {
// 			r.matcher = append(r.matcher, start)
// 		}
// 	case html.EndTagToken:
// 		r.buf1.Write(raw)
// 		if last > -1 && (token.DataAtom == dataAtom) {
// 			start = r.matcher[last]
// 			end := r.buf1.Len() + 1
// 			r.matcher[last] = -1
// 			r.matcher = r.matcher[:last]
// 			if r.On != nil {
// 				r.On(r.buf1.Bytes()[start:end])
// 			}
// 			// fmt.Println(string(r.buf1.Bytes()[start:end]))
// 			// fmt.Println("*****table****")
// 		}
// 	default:
// 		r.buf1.Write(raw)
// 	}

// }
