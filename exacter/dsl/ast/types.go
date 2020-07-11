package ast

import (
	"regexp"
	"time"
)

type TokenType int

// Indent string for formatted TICKscripts
const indentStep = "    "

type ValueType uint8

const (
	InvalidType ValueType = iota
	TFloat
	TInt
	TString
	TBool
	TRegex
	TTime
	TDuration
	TLambda
	TList
	TStar
	TMissing
)

//String representation of an TokenType
func (v ValueType) String() string {
	switch v {
	case TFloat:
		return "float"
	case TInt:
		return "int"
	case TString:
		return "string"
	case TBool:
		return "boolean"
	case TRegex:
		return "regex"
	case TTime:
		return "time"
	case TDuration:
		return "duration"
	case TLambda:
		return "lambda"
	case TList:
		return "list"
	case TStar:
		return "star"
	case TMissing:
		return "missing"
	}

	return "invalid type"
}
func TypeOf(v interface{}) ValueType {
	switch v.(type) {
	case float64:
		return TFloat
	case int64:
		return TInt
	case string:
		return TString
	case bool:
		return TBool
	case *regexp.Regexp:
		return TRegex
	case time.Time:
		return TTime
	case time.Duration:
		return TDuration
	// case *LambdaNode:
	// 	return TLambda
	case []interface{}:
		return TList
	case []string:
		return TList
	// case *StarNode:
	// 	return TStar
	// case *Missing:
	// 	return TMissing
	default:
		return InvalidType
	}
}

//String representation of an TokenType
func (t TokenType) String() string {

	// return "operatorStr[t]"
	return operatorStr[t]
}

const (
	TokenError TokenType = iota
	TokenNot
	TokenPlus
	TokenMinus
	TokenMult
	TokenDiv
	TokenMod
	TokenEqual
	TokenMatch
	TokenNotEqual
	TokenLess
	TokenGreater
	TokenLessEqual
	TokenGreaterEqual
	TokenRegexEqual
	TokenRegexNotEqual
	TokenAnd
	TokenOr
)

var stringTokenMap = map[string]TokenType{
	">":   TokenGreater,
	">=":  TokenGreaterEqual,
	"<":   TokenLess,
	"<=":  TokenLessEqual,
	"==":  TokenEqual,
	"!=":  TokenNotEqual,
	":":   TokenMatch,
	"AND": TokenAnd,
	"OR":  TokenOr,
	"/":   TokenDiv,
	"-":   TokenMinus,
	"+":   TokenPlus,
	"*":   TokenMult,
	"=~":  TokenRegexEqual,
	"!~":  TokenRegexNotEqual,

	// TokenNot:           "!",
	// TokenPlus:          "+",
	// TokenMinus:         "-",
	// TokenMult:          "*",
	// TokenDiv:           "/",
	// TokenMod:           "%",
}

var operatorStr = [...]string{
	TokenNot:           "!",
	TokenPlus:          "+",
	TokenMinus:         "-",
	TokenMult:          "*",
	TokenDiv:           "/",
	TokenMod:           "%",
	TokenEqual:         "==",
	TokenMatch:         ":",
	TokenNotEqual:      "!=",
	TokenLess:          "<",
	TokenGreater:       ">",
	TokenLessEqual:     "<=",
	TokenGreaterEqual:  ">=",
	TokenRegexEqual:    "=~",
	TokenRegexNotEqual: "!~",
	TokenAnd:           "AND",
	TokenOr:            "OR",
}

func IsMathOperator(typ TokenType) bool {
	// return typ > begin_tok_operator_math && typ < end_tok_operator_math
	return true

}

// True if token type is an operator used in comparisons.
func IsCompOperator(typ TokenType) bool {
	// return typ > begin_tok_operator_comp && typ < end_tok_operator_comp
	return true

}

func IsLogicalOperator(typ TokenType) bool {
	// return typ > begin_tok_operator_logic && typ < end_tok_operator_logic
	return true

}

type Missing struct{}

var MissingValue = Missing{}
