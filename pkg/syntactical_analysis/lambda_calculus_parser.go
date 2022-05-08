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
	Unparse(entity.Ast) (string, error)
	BetaReduce(entity.Ast) (entity.Ast, error)
}

type lL1PredictableParser struct {
	logging logging.Logger
	buffer  entity.TokenBuffer
}

func (l *lL1PredictableParser) Parse(t []entity.Token) (entity.Ast, error) {
	l.bufferInit(t)

	root, err := l.parse(entity.TERM)
	if err != nil {
		return nil, err
	}
	ast := entity.NewAst(root)
	l.logging.Debugf("computed ast: \n%v", ast.Visualize())

	l.simplify(ast.Root())
	if err != nil {
		return nil, err
	}
	l.logging.Debugf("simplified ast: \n%v", ast.Visualize())

	return ast, nil
}

func (l *lL1PredictableParser) simplify(node entity.Node) {
	for i := len(node.Child()) - 1; i >= 0; i-- {
		l.simplify(node.Child()[i])

		if len(node.Child()[i].Child()) == 0 && !entity.IsTerminal(node.Child()[i].Token().Tag) {
			node.Delete(i)
		} else {
			switch node.Child()[i].Token().Tag {
			case entity.EPSILON, entity.LEFT_BRACKET, entity.RIGHT_BRACKET:
				{
					node.Delete(i)
				}
			case entity.TERMS:
				{
					node.Replace(i, node.Child()[i].Child()...)
				}
			}

		}
	}

}

func (l *lL1PredictableParser) BetaReduce(ast entity.Ast) (entity.Ast, error) {
	err := l.betaReduce(ast.Root())
	if err != nil {
		return nil, err
	}

	l.logging.Debugf("ast after beta-reduction:\n%s", ast.Visualize())
	return ast, nil
}

func (l *lL1PredictableParser) betaReduce(n entity.Node) error {
	for _, child := range n.Child() {
		if err := l.betaReduce(child); err != nil {
			return err
		}
	}

	if len(n.Child()) == 3 {
		lo := n.Child()[0]
		ro := n.Child()[2]

		if len(lo.Child()) == 0 || lo.Child()[0].Token().Tag != entity.LAMBDA {
			return nil
		}

		localVar := lo.Child()[1]
		lo.Delete(2)
		lo.Delete(1)
		lo.Delete(0)
		l.apply(lo, ro, localVar)

		n.Delete(2)
		n.Delete(1)
	}
	return nil

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

func (l *lL1PredictableParser) Unparse(ast entity.Ast) (string, error) {
	res, err := l.unparse(ast.Root())
	l.logging.Debugf(`unparsed to "%s"`, res)
	return res, err
}

func (l *lL1PredictableParser) unparse(node entity.Node) (string, error) {
	if entity.IsTerminal(node.Token().Tag) {
		return node.Label(), nil
	}

	var res string
	for _, child := range node.Child() {
		parseChild, err := l.unparse(child)
		if err != nil {
			return res, err
		}
		res += parseChild
	}

	if node.Token().Tag == entity.TERM && len(node.Child()) > 2 {
		res = "(" + res + ")"
	}

	return res, nil
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

func (l *lL1PredictableParser) apply(lop entity.Node, rop entity.Node, localVar entity.Node) {
	for i, child := range lop.Child() {
		if *child.Token() == *localVar.Token() {
			lop.Replace(i, rop)
		} else if child.Token().Tag == entity.LAMBDA {
			return
		} else {
			l.apply(child, rop, localVar)
		}
	}
}
