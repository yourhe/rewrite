package html

import (
	"regexp"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type TagsTable map[string][]matcher

type match struct {
	key   string
	value *string
	regex *regexp.Regexp
}

type matcher struct {
	matchs []match
}

func NewTagsTable() TagsTable {
	return make(map[string][]matcher)
}
func readAttrName(src string) (name string, regex *regexp.Regexp) {
	name = src
	i := strings.Index(src, "/")
	if i > -1 {
		name = src[:i]
		j := strings.LastIndex(src, "/")
		rgStr := src[i+1 : j]
		regex = regexp.MustCompile(rgStr)
	}
	return
}

func (r TagsTable) AddTagQuery(query, attrName string) TagsTable {
	if query == "" {
		return r
	}
	matchs := []match{}
	attrName, regex := readAttrName(attrName)
	// if r.tagTable == nil {
	// 	r.tagTable = map[string][]matcher{}
	// }
	if attrName != "" {
		matchs = append(matchs, match{key: attrName, regex: regex})
	}
	var tagName = query
	i := strings.Index(query, "[")
	if i > -1 {
		j := strings.LastIndex(query, "]")
		if j > -1 {
			tagName = query[:i]
			expr := query[i+1 : j]
			parts := strings.Split(expr, "=")
			key := strings.TrimSpace(parts[0])
			var val *string
			if len(parts) > 1 {
				v := strings.TrimSpace(parts[1])
				if v[0] == '"' || v[0] == '\'' {
					v = v[1 : len(v)-1]
				}
				val = &v
			}
			matchs = append(matchs, match{
				key:   key,
				value: val,
			})
		}

	}

	r[tagName] = append(r[tagName], matcher{
		matchs: matchs,
	})

	// r.tagTable[tagName] = &AttrsRewriter{
	// 	AttrRewriter: AttrRewriter{
	// 		attrName: matcher{key: attrName, rewrite: rw},
	// 	},
	// }
	// r.tagTable[tagName].AttrRewriter
	return r
}
func (tt TagsTable) queryTags(token *html.Token) (matched bool, atomData atom.Atom) {
	tag := tt[token.Data]
	if tag == nil {
		return false, 0
	}

	for _, tagMatchs := range tag {
		var matched = make([]bool, len(tagMatchs.matchs))
		for i, matcher := range tagMatchs.matchs {
			for _, attr := range token.Attr {

				if attr.Key != matcher.key {
					continue
				}

				if matcher.value != nil {
					if attr.Val != *matcher.value {
						continue
					}
				}
				// if matcher.rewrite != nil {

				// 	rewriter = &tagMatchs.matchs[i]

				// 	idx = j
				// }
				matched[i] = true
				var e = true
				for _, b := range matched {
					e = e && b
				}
				if e {
					return true, token.DataAtom
				}
			}

		}
	}
	return false, 0
}
