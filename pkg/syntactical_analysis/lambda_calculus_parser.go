package syntactical_analyzer

import (
	"context"
	"errors"
	"fmt"
	"math-parser/pkg/entity"
	"math-parser/pkg/utils/logging"
)

const (
	TERM          = "Λ"
	APPLICATION   = "_"
	ABSTRACTION   = "."
	LAMBDA        = "λ"
	EPSILON       = "ε"
	TERMS         = "Λs"
	FACTOR        = "FACTOR"
	NUMBER        = "NUMBER"
	TERMINAL      = "TERMINAL"
	LEFT_BRACKET  = "("
	RIGHT_BRACKET = ")"
)

func NewLL1PredictableParser(ctx context.Context) LL1PredictableParser {
	return &lL1PredictableParser{
		logging: ctx.Value("logger").(logging.Logger),
	}
}

type LL1PredictableParser interface {
	Parse([]entity.Token) (entity.Ast, error)
}

type lL1PredictableParser struct {
	logging logging.Logger
	buffer  entity.TokenBuffer
}

func (l *lL1PredictableParser) Parse(t []entity.Token) (entity.Ast, error) {
	l.bufferInit(t)

	if root, err := l.parse(entity.TERM); err == nil {
		ast := entity.NewAst(root)
		l.logging.Debugf("computed ast: \n%v", ast.Visualize())

		return ast, nil
	} else {
		return nil, err
	}
}

func (l *lL1PredictableParser) bufferInit(t []entity.Token) {
	l.buffer = entity.NewTokenBuffer(append(t, entity.Token{Tag: entity.EPSILON}))
}

func (l *lL1PredictableParser) parse(nonTerminalTag entity.Tag) (entity.Node, error) {
	rules := map[entity.Tag]map[entity.Tag][]entity.Tag{
		entity.TERM: {
			entity.VARIABLE:     {entity.VARIABLE, entity.TERMS},
			entity.LAMBDA:       {entity.LAMBDA, entity.VARIABLE, entity.ABSTRACTION, entity.TERM, entity.TERMS},
			entity.LEFT_BRACKET: {entity.LEFT_BRACKET, entity.TERM, entity.RIGHT_BRACKET, entity.TERMS},
		},
		entity.TERMS: {
			entity.APPLICATION: {entity.APPLICATION, entity.TERM},
			entity.EPSILON:     {entity.EPSILON},
		},
	}
	if rule, ok := rules[nonTerminalTag]; !ok {
		return nil, errors.New("can't define rule")
	} else {
		res := l.NewNodeFromNonTerminal(nonTerminalTag)
		if prod, ok := rule[l.buffer.Lookahead().Tag]; ok {
			for _, t := range prod {
				var child entity.Node
				if t == entity.EPSILON {
					child = l.NewNodeFromTerminal(*entity.NewEpsilonToken())
				} else if entity.IsTerminal(t) {
					l.buffer.NextToken()
					child = l.NewNodeFromTerminal(*l.buffer.Current())
					if child.Token().Tag != t {
						return nil, errors.New(fmt.Sprintf("expected %d instead of %v", t, child))
					}
				} else {
					var err error
					if child, err = l.parse(t); err != nil {
						return nil, err
					}
				}
				res.AddChildToEnd(child)
			}
		}
		return res, nil
	}
}

func (l lL1PredictableParser) NewNodeFromNonTerminal(t entity.Tag) entity.Node {
	switch t {
	case entity.TERM:
		{
			return entity.NewNode(TERM, entity.Token{
				Tag:   t,
				Value: TERM,
			})
		}
	case entity.TERMS:
		{
			return entity.NewNode(TERMS, entity.Token{
				Tag:   t,
				Value: TERMS,
			})
		}
	default:
		{
			return nil
		}
	}

}

func (l lL1PredictableParser) NewNodeFromTerminal(t entity.Token) entity.Node {
	switch t.Tag {
	case entity.APPLICATION:
		{
			return entity.NewNode(APPLICATION, t)
		}
	case entity.ABSTRACTION:
		{
			return entity.NewNode(ABSTRACTION, t)
		}
	case entity.LAMBDA:
		{
			return entity.NewNode(LAMBDA, t)
		}
	case entity.VARIABLE:
		{
			return entity.NewNode(fmt.Sprintf("%s", t.Value), t)
		}
	case entity.LEFT_BRACKET:
		{
			return entity.NewNode(LEFT_BRACKET, t)
		}
	case entity.RIGHT_BRACKET:
		{
			return entity.NewNode(RIGHT_BRACKET, t)
		}
	case entity.EPSILON:
		{
			return entity.NewNode(EPSILON, t)
		}
	default:
		{
			l.logging.Debugf("unknown token tag %s", t)
			return nil
		}
	}
}
