package entity

type Tag int

const (
	// Terminals
	ABSTRACTION Tag = iota
	APPLICATION
	VARIABLE
	LAMBDA
	LEFT_BRACKET
	RIGHT_BRACKET

	// Non-Terminals
	TERM
	TERMS
	EPSILON
)

func IsTerminal(t Tag) bool {
	return map[Tag]bool{
		ABSTRACTION:   true,
		APPLICATION:   true,
		VARIABLE:      true,
		LAMBDA:        true,
		LEFT_BRACKET:  true,
		RIGHT_BRACKET: true,
	}[t]
}

func NewEpsilonToken() *Token {
	return &Token{
		Tag:   EPSILON,
		Value: "ε",
	}
}

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
	if lexem == "(" {
		tag = LEFT_BRACKET
	} else {
		tag = RIGHT_BRACKET
	}
	return &Token{
		Tag:   tag,
		Value: LEFT_BRACKET,
	}
}

func NewVariableToken(lexem string) *Token {
	return &Token{
		Tag:   VARIABLE,
		Value: lexem,
	}
}
