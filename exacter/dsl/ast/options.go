package ast

type Option func(Node)

// type DeclarationOption func(*DeclarationNode)

//SetPosition option
func SetPosition(pos, line, char int) Option {
	return func(prog Node) {
		prog.Set(pos, line, char)
	}
}

//SetOperator option
func SetOperator(op TokenType) Option {
	return func(prog Node) {
		if so, ok := prog.(interface {
			SetOperator(op TokenType)
		}); ok {
			so.SetOperator(op)
		}
	}
}

func SetParens(parent bool) Option {
	return func(prog Node) {
		if so, ok := prog.(interface {
			SetParens(bool)
		}); ok {
			so.SetParens(parent)
		}
	}
}
