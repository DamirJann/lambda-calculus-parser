package syntactical_analyzer

import (
	"context"
	"gotest.tools/assert"
	"math-parser/pkg/lexical_analysis"
	"math-parser/pkg/utils/logging"
	"testing"
)

func TestLexicalAnalyzer_Parse(t *testing.T) {
	var tests = []struct {
		name     string
		scenario func(*testing.T)
	}{
		{
			name:     "Happy flow. Parse expression",
			scenario: happyFlowParseExpression,
		},
		{
			name:     "Happy flow. Parse expression with double abstraction",
			scenario: happyFlowParseExpressionWithDoubleAbstraction,
		},
		{
			name:     "Happy flow. Parse expression with simple application",
			scenario: happyFlowParseExpressionWithDoubleApplication,
		},
		{
			name:     "Happy flow. Parse expression with application and abstraction",
			scenario: happyFlowParseExpressionWithApplicationAndAbstraction,
		},
		{
			name:     "Happy flow. Parse expression with brackets",
			scenario: happyFlowParseExpressionWithBrackets,
		},
		{
			name:     "Happy flow. Parse expression with brackets1",
			scenario: happyFlowParseExpressionWithBrackets1,
		},
	}

	t.Parallel()
	for _, test := range tests {
		t.Run(test.name, test.scenario)
	}
}

func happyFlowParseExpression(t *testing.T) {
	// arrange
	ctx := context.WithValue(context.Background(), "logger", logging.NewBuiltinLogger())
	automata := lexical_analysis.NewAutomata()
	analyzer := lexical_analysis.NewLexicalAnalyzer(ctx, automata)
	parser := NewLL1PredictableParser(ctx)

	// act
	tk, _ := analyzer.Tokenize("λx.x")
	_, err := parser.Parse(tk)

	// assert
	assert.Equal(t, err, nil)
}

func happyFlowParseExpressionWithDoubleAbstraction(t *testing.T) {
	// arrange
	ctx := context.WithValue(context.Background(), "logger", logging.NewBuiltinLogger())
	automata := lexical_analysis.NewAutomata()
	analyzer := lexical_analysis.NewLexicalAnalyzer(ctx, automata)
	parser := NewLL1PredictableParser(ctx)

	// act
	tk, _ := analyzer.Tokenize("λx.λy.y")
	_, err := parser.Parse(tk)

	// assert
	assert.Equal(t, err, nil)
}

func happyFlowParseExpressionWithDoubleApplication(t *testing.T) {
	// arrange
	ctx := context.WithValue(context.Background(), "logger", logging.NewBuiltinLogger())
	automata := lexical_analysis.NewAutomata()
	analyzer := lexical_analysis.NewLexicalAnalyzer(ctx, automata)
	parser := NewLL1PredictableParser(ctx)

	// act
	tk, _ := analyzer.Tokenize("x_y")
	_, err := parser.Parse(tk)

	// assert
	assert.Equal(t, err, nil)
}

func happyFlowParseExpressionWithApplicationAndAbstraction(t *testing.T) {
	// arrange
	ctx := context.WithValue(context.Background(), "logger", logging.NewBuiltinLogger())
	automata := lexical_analysis.NewAutomata()
	analyzer := lexical_analysis.NewLexicalAnalyzer(ctx, automata)
	parser := NewLL1PredictableParser(ctx)

	// act
	tk, _ := analyzer.Tokenize("x_λy.x_y_z")
	_, err := parser.Parse(tk)

	// assert
	assert.Equal(t, err, nil)
}

func happyFlowParseExpressionWithBrackets(t *testing.T) {
	// arrange
	ctx := context.WithValue(context.Background(), "logger", logging.NewBuiltinLogger())
	automata := lexical_analysis.NewAutomata()
	analyzer := lexical_analysis.NewLexicalAnalyzer(ctx, automata)
	parser := NewLL1PredictableParser(ctx)

	// act
	tk, _ := analyzer.Tokenize("x_(λy.x)_y_z")
	_, err := parser.Parse(tk)

	// assert
	assert.Equal(t, err, nil)
}

func happyFlowParseExpressionWithBrackets1(t *testing.T) {
	// arrange
	ctx := context.WithValue(context.Background(), "logger", logging.NewBuiltinLogger())
	automata := lexical_analysis.NewAutomata()
	analyzer := lexical_analysis.NewLexicalAnalyzer(ctx, automata)
	parser := NewLL1PredictableParser(ctx)

	// act
	tk, _ := analyzer.Tokenize("x_(λy.(x))_y_(z_z)")
	_, err := parser.Parse(tk)

	// assert
	assert.Equal(t, err, nil)
}

func TestLexicalAnalyzer_Unparse(t *testing.T) {
	var tests = []struct {
		name     string
		scenario func(*testing.T)
	}{
		{
			name:     "Happy flow. Unparse expression",
			scenario: happyFlowUnparseExpression,
		},
		{
			name:     "Happy flow. Unparse expression with double abstraction",
			scenario: happyFlowUnparseExpressionWithDoubleAbstraction,
		},
		{
			name:     "Happy flow. Unparse expression with simple application",
			scenario: happyFlowUnparseExpressionWithDoubleApplication,
		},
		{
			name:     "Happy flow. Unparse expression with application and abstraction",
			scenario: happyFlowUnparseExpressionWithApplicationAndAbstraction,
		},
		{
			name:     "Happy flow. Unparse expression with brackets",
			scenario: happyFlowUnparseExpressionWithBrackets,
		},
		{
			name:     "Happy flow. Unparse expression with brackets1",
			scenario: happyFlowUnparseExpressionWithBrackets1,
		},
	}

	t.Parallel()
	for _, test := range tests {
		t.Run(test.name, test.scenario)
	}
}
func happyFlowUnparseExpression(t *testing.T) {
	// arrange
	ctx := context.WithValue(context.Background(), "logger", logging.NewBuiltinLogger())
	automata := lexical_analysis.NewAutomata()
	analyzer := lexical_analysis.NewLexicalAnalyzer(ctx, automata)
	parser := NewLL1PredictableParser(ctx)
	expression := "λx.x"

	// act
	tk, _ := analyzer.Tokenize(expression)
	ast, err := parser.Parse(tk)
	res, err := parser.Unparse(ast)
	// assert
	assert.Equal(t, err, nil)
	assert.Equal(t, res, expression)
}

func happyFlowUnparseExpressionWithDoubleAbstraction(t *testing.T) {
	// arrange
	ctx := context.WithValue(context.Background(), "logger", logging.NewBuiltinLogger())
	automata := lexical_analysis.NewAutomata()
	analyzer := lexical_analysis.NewLexicalAnalyzer(ctx, automata)
	parser := NewLL1PredictableParser(ctx)

	// act
	expression := "λx.λy.y"
	tk, _ := analyzer.Tokenize(expression)
	ast, err := parser.Parse(tk)
	res, err := parser.Unparse(ast)
	// assert
	assert.Equal(t, err, nil)
	assert.Equal(t, res, expression)
}

func happyFlowUnparseExpressionWithDoubleApplication(t *testing.T) {
	// arrange
	ctx := context.WithValue(context.Background(), "logger", logging.NewBuiltinLogger())
	automata := lexical_analysis.NewAutomata()
	analyzer := lexical_analysis.NewLexicalAnalyzer(ctx, automata)
	parser := NewLL1PredictableParser(ctx)

	// act
	expression := "x_y"
	tk, _ := analyzer.Tokenize(expression)
	ast, err := parser.Parse(tk)
	res, err := parser.Unparse(ast)
	// assert
	assert.Equal(t, err, nil)
	assert.Equal(t, res, expression)
}

func happyFlowUnparseExpressionWithApplicationAndAbstraction(t *testing.T) {
	// arrange
	ctx := context.WithValue(context.Background(), "logger", logging.NewBuiltinLogger())
	automata := lexical_analysis.NewAutomata()
	analyzer := lexical_analysis.NewLexicalAnalyzer(ctx, automata)
	parser := NewLL1PredictableParser(ctx)

	// act
	expression := "x_λy.x_y_z"
	tk, _ := analyzer.Tokenize(expression)
	ast, err := parser.Parse(tk)
	res, err := parser.Unparse(ast)
	// assert
	assert.Equal(t, err, nil)
	assert.Equal(t, res, expression)
}

func happyFlowUnparseExpressionWithBrackets(t *testing.T) {
	// arrange
	ctx := context.WithValue(context.Background(), "logger", logging.NewBuiltinLogger())
	automata := lexical_analysis.NewAutomata()
	analyzer := lexical_analysis.NewLexicalAnalyzer(ctx, automata)
	parser := NewLL1PredictableParser(ctx)

	// act
	expression := "x_(λy.x)_y_z"
	tk, _ := analyzer.Tokenize(expression)
	ast, err := parser.Parse(tk)
	res, err := parser.Unparse(ast)
	// assert
	assert.Equal(t, err, nil)
	assert.Equal(t, res, expression)
}

func happyFlowUnparseExpressionWithBrackets1(t *testing.T) {
	// arrange
	ctx := context.WithValue(context.Background(), "logger", logging.NewBuiltinLogger())
	automata := lexical_analysis.NewAutomata()
	analyzer := lexical_analysis.NewLexicalAnalyzer(ctx, automata)
	parser := NewLL1PredictableParser(ctx)

	// act
	expression := "x_(λy.(x))_y_(z_z)"
	tk, _ := analyzer.Tokenize(expression)
	ast, err := parser.Parse(tk)
	res, err := parser.Unparse(ast)
	// assert
	assert.Equal(t, err, nil)
	assert.Equal(t, res, expression)
}
