package lexical_analysis

import (
	"bytes"
	"errors"
	"io"
	"math-parser/pkg/entity"
)

type automata struct {
	input *bytes.Buffer
	lexem string
}

type Automata interface {
	extractToken(input *bytes.Buffer) (*entity.Token, error)
}

func NewAutomata() Automata {
	return &automata{}
}

func (a *automata) Peek() (rune, error) {
	r, _, err := a.input.ReadRune()
	return r, err
}

func (a *automata) Lookahead() (rune, error) {
	b, _, err := a.input.ReadRune()
	if err == nil {
		err = a.input.UnreadRune()
	}
	if err == io.EOF {
		return EOF, nil
	}
	return b, err
}

func (a *automata) Unread() error {
	return a.input.UnreadRune()
}

func (a *automata) extractToken(input *bytes.Buffer) (*entity.Token, error) {
	a.input = input
	a.lexem = ""

	lookahead, err := a.Lookahead()
	if lookahead == EOF {
		return nil, io.EOF
	}
	if err != nil {
		return nil, err
	}

	return a.s1()
}

func (a *automata) s1() (*entity.Token, error) {
	peek, err := a.Peek()

	if err != nil {
		return nil, err
	}

	if nextState := a.s1TransitTo(peek); nextState != nil {
		a.lexem += string(peek)
		return nextState()
	} else {
		return nil, errors.New("error in S1 state")
	}

}

func (a *automata) s2() (*entity.Token, error) {
	return entity.NewAbstractionToken(a.lexem), nil
}

func (a *automata) s3() (*entity.Token, error) {
	return entity.NewApplicationToken(a.lexem), nil
}
func (a *automata) s4() (*entity.Token, error) {
	return entity.NewLambdaToken(a.lexem), nil
}

func (a *automata) s5() (*entity.Token, error) {
	return entity.NewVariableToken(a.lexem), nil
}

func (a *automata) s6() (*entity.Token, error) {
	return entity.NewBracketToken(a.lexem), nil
}

func (a *automata) s1TransitTo(lookahead rune) func() (*entity.Token, error) {
	res, ok := map[rune]func() (*entity.Token, error){
		ABSTRACTION:   a.s2,
		APPLICATION:   a.s3,
		LAMBDA:        a.s4,
		LEFT_BRACKET:  a.s6,
		RIGHT_BRACKET: a.s6,
	}[lookahead]
	if ok {
		return res
	} else {
		if lookahead >= 'a' && lookahead <= 'z' {
			return a.s5
		}
	}
	return nil
}
