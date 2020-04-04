package html

import (
	"bytes"
	"fmt"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type Table struct {
	rows [][]Element
}

type Element struct {
	Name   string
	Value  string
	Link   string
	Values []string
	Links  []string
	Childs []Element
	Attrs  map[string]string
}

func (t *Table) Parse(p []byte) error {
	// inconsistentNode := &html.Node{
	// 	Type:     html.ElementNode,
	// 	DataAtom: atom.Table,
	// 	Data:     "table",
	// }
	tokenizer := html.NewTokenizer(bytes.NewReader(p))

	var f func(*html.Tokenizer)
	f = func(tokenizer *html.Tokenizer) {
		tp := tokenizer.Next()
		switch tp {

		case html.ErrorToken:
			fmt.Println(t.rows)
			return
		case html.StartTagToken, html.SelfClosingTagToken:
			// token := tokenizer.Token()
			// fmt.Println(token.DataAtom)
			t.ReadTable(tokenizer)
		}
		f(tokenizer)
	}
	f(tokenizer)
	// for n := node.FirstChild; n != nil; n = node.NextSibling {
	// 	fmt.Printf("%#v", n)
	// }
	return nil
}
func (t *Table) ReadTable(tokenizer *html.Tokenizer) {
	token := tokenizer.Token()
	switch token.DataAtom {
	case atom.Table:
		//init table
	case atom.Tr:
		t.rows = append(t.rows, []Element{})
	case atom.Td, atom.Th:
		current := len(t.rows) - 1
		col := t.ReadColumn(tokenizer)
		t.rows[current] = append(t.rows[current], col)
	}
}
func (t *Table) ReadRow(p []byte) {

}
func (t *Table) ReadColumn(tokenizer *html.Tokenizer) Element {
	element := Element{}
	content := []string{}

	var isA bool
	// var isImg bool
loop:
	tp := tokenizer.Next()
	token := tokenizer.Token()

	if token.DataAtom == atom.Td && tp == html.EndTagToken {
		if content != nil && len(content) > 0 {
			element.Values = append(element.Values, strings.Join(content, ""))
		}
		return element
	}

	if tp == html.TextToken && isA {
		text := (string)(token.Data)
		t := strings.TrimSpace(text)
		content = append(content, t)
	}
	switch token.DataAtom {
	case atom.A:
		if tp == html.EndTagToken {
			element.Values = append(element.Values, strings.Join(content, ""))
			content = nil
			isA = false
		} else {
			isA = true
		}

		for _, attr := range token.Attr {
			if attr.Key == "href" {
				element.Links = append(element.Links, attr.Val)
			}

		}
	default:

	}
	goto loop

}

// func (t *Table) toEndTag(dataAtom atom.Atom, tokenizer *html.Tokenizer) error {
// 	left := 1
// 	for {
// 		tp := tokenizer.Next()
// 		token := tokenizer.Token()
// 		switch tp {
// 		case html.StartTagToken:
// 			if token.DataAtom == dataAtom {
// 				left++
// 			}
// 		case html.EndTagToken:
// 			if token.DataAtom == dataAtom {
// 				left--
// 				if left == 0 {
// 					return nil
// 				}
// 			}

// 		}
// 	}

// }

// func (t *Table) ParseDoc(p []byte) error {
// 	node, err := xmlpath.Parse(bytes.NewReader(p))
// 	if err != nil {
// 		fmt.Println(err)
// 		return err
// 	}
// 	path := xmlpath.MustCompile("")
// 	path.Iter(node)
// }
func (t *Table) ParseDoc(p []byte) error {
	// inconsistentNode := &html.Node{
	// 	Type:     html.ElementNode,
	// 	DataAtom: atom.Table,
	// 	Data:     "table",
	// }

	doc, err := html.Parse(bytes.NewReader(p))
	fmt.Printf("%#v", doc)
	if err != nil {
		fmt.Println(err)
		return err
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		switch n.Type {
		case html.TextNode:
			fmt.Println(n.Data)
		case html.ElementNode:
			if n.Data == "a" {
				for _, a := range n.Attr {
					if a.Key == "href" {
						fmt.Println(a.Val)
						break
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	// for n := node.FirstChild; n != nil; n = node.NextSibling {
	// 	fmt.Printf("%#v", n)
	// }
	return nil
}
