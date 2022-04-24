package lexical_analysis

const (
	LAMBDA      = byte('#')
	APPLICATION = byte(' ')
	ABSTRACTION = byte('.')

	LEFT_BRACKET  = byte('(')
	RIGHT_BRACKET = byte(')')

	EOF = 0
)
