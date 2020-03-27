package rewrite

import (
	"io"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

type Transform struct {
	Encoding string
	// transform transform.Transformer
	encoding encoding.Encoding
}

func (t *Transform) NewDecodeReader(r io.Reader) io.Reader {
	if t.encoding != nil {
		return transform.NewReader(r, t.encoding.NewDecoder())
	}
	if t.Encoding == "" {
		return r
	}
	e, _ := charset.Lookup(t.Encoding)
	if e != nil {
		t.encoding = e
		return transform.NewReader(r, e.NewDecoder())
	}
	return r
}

func (t *Transform) NewEncodeReader(r io.Reader) io.Reader {
	if t.encoding != nil {
		return transform.NewReader(r, t.encoding.NewEncoder())
	}
	if t.Encoding == "" {
		return r
	}
	e, _ := charset.Lookup(t.Encoding)
	if e != nil {
		t.encoding = e
		return transform.NewReader(r, e.NewEncoder())
	}
	return r
}

func DetermineEncoding(contentType string) *Transform {
	e, name, certain := charset.DetermineEncoding(nil, contentType)
	if certain {
		return &Transform{
			Encoding: name,
			encoding: e,
		}
	}
	return nil
}
