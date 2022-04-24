package entity

type Tag int

const (
	ABSTRACTION Tag = iota
	APPLICATION
	VARIABLE
	LAMBDA
	LEFT_BRACKET
	RIGHT_BRACKET
)

type Token struct {
	Tag   Tag
	Value interface{}
}

func NewLambdaToken(lexem string) *Token {
	return &Token{
		Tag:   LAMBDA,
		Value: lexem,
	}
}

func NewApplicationToken(lexem string) *Token {
	return &Token{
		Tag:   APPLICATION,
		Value: lexem,
	}
}

func NewAbstractionToken(lexem string) *Token {
	return &Token{
		Tag:   ABSTRACTION,
		Value: lexem,
	}
}

func NewBracketToken(lexem string) *Token {
	var tag Tag
	if lexem == string("(") {
		tag = LEFT_BRACKET
	} else {
		tag = RIGHT_BRACKET
	}
	return &Token{
		Tag:   tag,
		Value: lexem,
	}
}

func NewVariableToken(lexem string) *Token {
	return &Token{
		Tag:   VARIABLE,
		Value: lexem,
	}
}
