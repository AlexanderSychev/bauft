package parser

type TokenType byte

const (
	TokenTypeSymbol = iota
	TokenTypeString
	TokenTypeNumber
	TokenTypeBoolean
	TokenTypeComma
	TokenTypeColon
	TokenTypeSemicolon
)

type Token struct {
	row    int
	column int
	tp     TokenType
}

func Tokenize(src string) []Token {
	result := make([]Token, 0, 1024)

	return result
}
