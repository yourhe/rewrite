package rewrite

import (
	"regexp"
	"sync"

	"gitlab.iyorhe.com/wfgz/reverseproxy/errors"
	rproto "gitlab.iyorhe.com/wfgz/reverseproxy/proto/rewrite"

	log "github.com/sirupsen/logrus"
)

var regexpMaps *RegexpMaps

type JavaScript struct {
	rwurls Rewriter
	rm     *RegexpMaps
	rules  *rproto.Rules
	// rw     Rewriter
}

func NewJavaScriptRewriter(mod Rewriter, rules *rproto.Rules) *JavaScript {
	rwjs := JavaScript{
		rwurls: mod,
		rm: &RegexpMaps{
			sync.Map{},
		},
		rules: rules,
	}
	// rwjs.rw = rwjs.ReplaceAllByRules()
	return &rwjs
}

func (js *JavaScript) Rewrite(p []byte) []byte {
	return js.ReplaceAllByRules()(p)
}

type ReplaceFunc func([]byte) []byte

func Noop(p []byte) []byte {
	return p
}

//ReplaceStringFunc 字符串替换
func ReplaceStringFunc(rg *regexp.Regexp, r []byte) ReplaceFunc {
	return func(s []byte) []byte {
		return rg.ReplaceAll(s, r)
	}
}

//ReplaceURLFunc url替换
func ReplaceURLFunc(rg *regexp.Regexp, urls Rewriter) ReplaceFunc {
	return func(s []byte) []byte {
		return rg.ReplaceAllFunc(s, urls.Rewrite)
	}
}

//ReplaceAllByRules 输入rules并替换
func (js *JavaScript) ReplaceAllByRules() func(p []byte) []byte {
	_rules := js.rules
	if _rules == nil || len(_rules.Rules) == 0 {
		log.WithError(errors.RulesNotFound).Error("BuildUrlsRewriterByRules")
		return Noop
	}
	var fns = []ReplaceFunc{}
	for _, r := range _rules.Rules {
		if r.Type != rproto.RewriteType_RESPONSE || r.GetRwa() != rproto.RewriteAction_JS {
			continue
		}
		if fn := js.MacthReplaceFunc(r); fn != nil {
			fns = append(fns, fn)
		}

	}
	return ReplaceAllFunc(fns...)
}

//MacthReplaceFunc 输入rule，并返回匹配的ReplaceFunc
// rule.Replace == "" 返回ReplaceURLFunc
// rule.Replace ！= "" 返回ReplaceStringFunc
func (js *JavaScript) MacthReplaceFunc(r *rproto.Rule) ReplaceFunc {
	key := r.Regex
	var rg *regexp.Regexp
	var err error
	rg, err = js.rm.LoadOrStoreRegexp(key)
	if err != nil || rg == nil {
		return Noop
	}
	if r.Replace == "" {
		return ReplaceURLFunc(rg, js.rwurls)
	}

	return ReplaceStringFunc(rg, []byte(r.Replace))
}

//ReplaceAllFunc 输入ReplaceFunc数组，并执行ReplaceFunc
func ReplaceAllFunc(fns ...ReplaceFunc) func(p []byte) []byte {
	return func(p []byte) []byte {
		for _, fn := range fns {
			p = fn(p)
		}
		return p
	}

}

// func BuildRegexRewriterByRequestContext(r *http.Request) *RegexpMaps {
// 	var _rules = pctx.GetRulesByContextKey(r)
// 	if _rules == nil {
// 		log.WithError(errors.RulesNotFound).Error(r.URL)
// 		return nil
// 	}
// 	domain := fmt.Sprintf("%s%s", pctx.GetDomainByContextKey(r), pctx.GetServerPortByContextKey(r))
// 	return BuildRegexRewriterByRules(_rules, domain)
// }

// func BuildRegexRewriterByRules(_rules *rproto.Rules, domain string) (rm *RegexpMaps) {
// 	if _rules == nil {
// 		log.WithError(errors.RulesNotFound).Error("BuildUrlsRewriterByRules")
// 		// return nil
// 		return nil
// 	}
// 	rm = &RegexpMaps{
// 		sync.Map{},
// 	}
// 	// var urlsOpts = []Option{}
// 	for _, r := range _rules.Rules {

// 		if r.Type != rproto.RewriteType_RESPONSE || r.GetRwa() != rproto.RewriteAction_JS {
// 			continue
// 		}
// 		key := r.Regex
// 		var rg *regexp.Regexp
// 		var err error
// 		if regexpMaps != nil {
// 			rg, err = regexpMaps.LoadOrStoreRegexp(key)
// 			if err != nil {
// 				continue
// 			}
// 			if rg != nil {
// 				rm.Store(key, rg)
// 			}
// 		} else {
// 			_, err = rm.LoadOrStoreRegexp(key)
// 			if err != nil {
// 				continue
// 			}
// 		}
// 	}
// 	return
// }

type RegexpMaps struct {
	sync.Map
}

func (rm RegexpMaps) LoadOrStoreRegexp(key string) (rg *regexp.Regexp, err error) {
	if rgi, ok := rm.Load(key); !ok {
		rg, err = regexp.Compile(key)
		if err != nil {
			return
		}
		rm.Store(key, rg)
	} else {
		rg = rgi.(*regexp.Regexp)
	}
	return
}

func init() {
	(&sync.Once{}).Do(func() {
		regexpMaps = &RegexpMaps{
			sync.Map{},
		}
	})
}
