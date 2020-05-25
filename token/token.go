package token

type Type string

type Token struct {
	Type    Type
	Literal string
}

const (
	Illegal = "ILLEGAL"
	EOF     = "EOF"

	Ident = "IDENT"
	Int   = "INT"

	Assign   = "="
	Plus     = "+"
	Minus    = "-"
	Bang     = "!"
	Asterisk = "*"
	Slash    = "/"

	LT = "<"
	GT = ">"

	EQ    = "=="
	NotEQ = "!="

	Comma     = ","
	Semicolon = ";"

	Lparen = "("
	Rparen = ")"
	Lbrace = "{"
	Rbrace = "}"

	Function = "FUNCTION"
	Let      = "LET"
	True     = "TRUE"
	False    = "FALSE"
	If       = "IF"
	Else     = "ELSE"
	Return   = "RETURN"
)

var keywords = map[string]Type{
	"fn":     Function,
	"let":    Let,
	"true":   True,
	"false":  False,
	"if":     If,
	"else":   Else,
	"return": Return,
}

func LookUpIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return Ident
}
