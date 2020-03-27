package rewrite

type Config struct {
	DestUrl        string
	Defmod         Rewriter
	CookieRewriter Rewriter
	Rewriters      []RewriterType
	HeaderPrefix   string
	HeaderRules    map[string]RewriteRule
	RewritersMp    map[RewriterType]Rewriter
	Transform      *Transform
}

func DefaultConfig() *Config {
	return &Config{
		DestUrl:      "",
		Defmod:       NoopRewriter,
		HeaderPrefix: "X-Archive-Orig-",
		HeaderRules:  DefaultHeaderRewriters,
		// RewritersMp: map[string]Rewriter{
		// 	"Defmod": NoopRewriter,
		// 	// "CookieRewriter":
		// },
	}
}

func SetCookieRewriter(cookiesRW Rewriter) func(cfg *Config) {
	return func(cfg *Config) {
		cfg.CookieRewriter = cookiesRW
	}
}
func makeConfig(configs ...func(*Config)) *Config {
	cfg := DefaultConfig()
	for _, config := range configs {
		config(cfg)
	}
	return cfg
}

func makeRewriters(cfg *Config) map[RewriterType]Rewriter {
	rws := map[RewriterType]Rewriter{}
	for _, rwt := range cfg.Rewriters {
		switch rwt {
		case RwTypeCookie:
		}
	}

	return rws
}

// func SetTransform(transform Transform) func(cfg *Config) {
// 	return func(cfg *Config) {
// 		cfg.Transform = &transform
// 	}
// }
