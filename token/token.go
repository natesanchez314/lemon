package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL 	= "ILLEGAL"
	EOF 		= "EOF"

	// Identifiers & literals
	ID 			= "ID"
	INT 		= "INT"
	STRING		= "STRING"
	
	// Operators
	ASSIGN 		= "="
	PLUS 		= "+"
	MINUS 		= "-"
	MULT	 	= "*"
	DIV 		= "/"
	BANG 		= "!"
	LT 			= "<"
	GT 			= ">"
	EQ 			= "=="
	NOT_EQ 		= "!="

	// Delimiters
	COMMA 		= ","
	COLON		= ":"
	SEMICOLON 	= ";"

	LPAREN 		= "("
	RPAREN 		= ")"
	LBRACE 		= "{"
	RBRACE 		= "}"
	LBRACKET	= "["
	RBRACKET	= "]"

	// Keywords
	FUNCTION 	= "FUNCTION"
	LET 		= "LET"
	TRUE 		= "TRUE"
	FALSE 		= "FALSE"
	IF 			= "IF"
	ELSE 		= "ELSE"
	RETURN 		= "RETURN"
	MACRO		= "MACRO"
)

var keywords = map[string]TokenType {
	"fn": 		FUNCTION,
	"let": 		LET,
	"true": 	TRUE,
	"false": 	FALSE,
	"if": 		IF,
	"else": 	ELSE,
	"return": 	RETURN,
	"macro":	MACRO,
}

func LookUpId(id string) TokenType {
	if tok, ok := keywords[id]; ok {
		return tok
	}
	return ID
}