package token

// Type is the type of token
type Type string

// Token is the expression of token
type Token struct {
	Type    Type
	Literal string
}

const (
	// ILLEGAL means that the token is illegal
	ILLEGAL = "ILLEGAL"
	// EOF means EOF
	EOF = "EOF"
	// IDENT means identifier
	IDENT = "IDENT" // add, foobar, x, y, ...
	// INT means integer
	INT = "INT" // 1343456

	// ASSIGN means assignment token
	ASSIGN = "="
	// PLUS means plus token
	PLUS = "+"
	// MINUS means minus token
	MINUS = "-"
	// BANG means bang token
	BANG = "!"
	// ASTERISK means asterisk token
	ASTERISK = "*"
	// SLASH means slash token
	SLASH = "/"
	// COMMA means comma token
	COMMA = ","
	// SEMICOLON means semicolon token
	SEMICOLON = ";"

	// LT means less-than token
	LT = "<"
	// GT means greater-than token
	GT = ">"

	// LPAREN means left paren token
	LPAREN = "("
	// RPAREN means right paren token
	RPAREN = ")"
	// LBRACE means left brace token
	LBRACE = "{"
	// RBRACE means right brance token
	RBRACE = "}"

	// FUNCTION means function token
	FUNCTION = "FUNCTION"
	// LET means let token
	LET = "LET"
	// TRUE means true token
	TRUE = "TRUE"
	// FALSE means false token
	FALSE = "FALSE"
	// IF means if token
	IF = "IF"
	// ELSE means else token
	ELSE = "ELSE"
	// RETURN means return token
	RETURN = "RETURN"
)

var keywords = map[string]Type{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// New initializes Token
func New(t Type, l string) Token {
	return Token{t, l}
}

// NewChar initializes Token
func NewChar(t Type, char byte) Token {
	return New(t, string(char))
}

// NewIdent initializes Token
func NewIdent(ident string) Token {
	t := lookupIdent(ident)
	return New(t, ident)
}

// lookupIdent checks keywords, then returns Type
func lookupIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
