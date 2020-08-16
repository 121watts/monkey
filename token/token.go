package token

const (
	// ILLEGAL token
	ILLEGAL = "ILLEGAL"
	// EOF end of file
	EOF = "EOF"

	// IDENT indenfier
	IDENT = "IDENT"
	// INT integer
	INT = "INT"

	// ASSIGN operator
	ASSIGN = "="
	// PLUS operator
	PLUS = "+"
	// MINUS operator
	MINUS = "-"
	// BANG operator
	BANG = "!"
	// ASTERISK operator
	ASTERISK = "*"
	// SLASH operator
	SLASH = "/"

	// LT less than
	LT = "<"
	// GT greater than
	GT = ">"

	// COMMA delimiter
	COMMA = ","
	// SEMICOLON delimiter
	SEMICOLON = ";"

	// LPAREN left paren
	LPAREN = "("
	// RPAREN right paren
	RPAREN = ")"
	// LBRACE left curly brace
	LBRACE = "{"
	// RBRACE left curly brace
	RBRACE = "}"

	// FUNCTION keyword
	FUNCTION = "FUNCTION"
	// LET keyword
	LET = "LET"
	// TRUE keyword
	TRUE = "TRUE"
	// FALSE keyword
	FALSE = "FALSE"
	// IF keyword
	IF = "IF"
	// ELSE keyword
	ELSE = "ELSE"
	// RETURN keyword
	RETURN = "RETURN"

	// EQ equal
	EQ = "=="
	// NOT_EQ not equal
	NOT_EQ = "!="
)

// TokenType is the type of token we are representing in our lexer
type TokenType string

// Token is the struct that holds both the type and the literal value
type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdent checkes the keywords table to see whether the given
// identifier is, in fact, a keyword
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
