// rewrite is a package for modifying the contents of html & other web-related content types.
// it's primarily used as a tool to maintain the functionality of a web resource within the context
// of an archive
package rewrite

import (
	"bytes"
	"errors"
	"io"
)

var ErrNotFinished = errors.New("not finished")

type BuildRewriteReader func(r io.Reader) io.Reader

// Rewriter takes an input byte slice of and returns an output
// slice of rewritten bytes, the length of input & output will
// not necessarily match, implementations *may* alter input bytes
type Rewriter interface {
	Rewrite(i []byte) (o []byte)
}

type RewriterStream interface {
	RewriteStream(f io.Reader, t io.Writer) error
}

// RewriterType enumerates rewriters that operate on different
// types of content
type RewriterType int

const (
	RwTypeUnknown RewriterType = iota
	RwTypeUrl
	RwTypeHeader
	RwTypeContent
	RwTypeCookie
	RwTypeHtml
	RwTypeJavascript
	RwTypeCss
	RwTypeScript
)

func (rwt RewriterType) String() string {
	return map[RewriterType]string{
		RwTypeUnknown:    "",
		RwTypeUrl:        "url",
		RwTypeHeader:     "header",
		RwTypeContent:    "content",
		RwTypeCookie:     "cookie",
		RwTypeHtml:       "html",
		RwTypeJavascript: "javascript",
		RwTypeCss:        "css",
		RwTypeScript:     "script",
	}[rwt]
}

var NoopRewriter = PrefixRewriter{}
var DefaultCookieRewriter = CookieRewriter{}

// PrefixRewriter adds a prefix if not present
type PrefixRewriter struct {
	Prefix []byte
}

func (prw PrefixRewriter) Rewrite(p []byte) []byte {
	if !bytes.HasPrefix(p, prw.Prefix) {
		return append(prw.Prefix, p...)
	}
	return p
}
