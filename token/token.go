package token

type TokenType string

type Token struct{
	Type TokenType
	Literal string
}

const(
	ILLEGAL = "ILLEGAL"
	EOF = "EOF" //end of file

	IDENT = "IDENT" //add, foobar, x, y, etc
	INT = "INT" //123456

	ASSIGN = "=" 
	PLUS = "+"

	COMMA = "COMMA"
	SEMICOLON = "SEMICOLON"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET = "LET"
)