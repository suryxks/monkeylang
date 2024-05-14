package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var kewords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookUpIndent(indent string) TokenType {
	if tok, ok := kewords[indent]; ok {
		return tok
	}
	return IDENT
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//Identifiers + literals
	IDENT = "IDENT"
	INT   = "INT"

	//operators

	ASSIGN = "="
	PLUS   = "+"

	//Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
