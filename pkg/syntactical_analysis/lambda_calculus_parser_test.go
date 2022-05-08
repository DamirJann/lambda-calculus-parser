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
	tk, _ := analyzer.Tokenize("(λy.x_z)_y")
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
	assert.Equal(t, res, "(λx.x)")
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
	assert.Equal(t, res, "(λx.(λy.y))")
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
	assert.Equal(t, res, "(x_y)")
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
	assert.Equal(t, res, "(x_(λy.(x_(y_z))))")
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
	assert.Equal(t, res, "(x_((λy.x)_(y_z)))")
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
	assert.Equal(t, res, "(x_((λy.x)_(y_(z_z))))")
}

func TestLexicalAnalyzer_BetaReduction(t *testing.T) {
	var tests = []struct {
		name     string
		scenario func(*testing.T)
	}{
		{
			name:     "Happy flow. Beta reduction1",
			scenario: happyFlowBetaReduction1,
		},
		{
			name:     "Happy flow. Beta reduction2",
			scenario: happyFlowBetaReduction2,
		},
		{
			name:     "Happy flow. Beta reduction3",
			scenario: happyFlowBetaReduction3,
		},
		{
			name:     "Happy flow. Beta reduction4",
			scenario: happyFlowBetaReduction4,
		},
		{
			name:     "Happy flow. Beta reduction5",
			scenario: happyFlowBetaReduction5,
		},
		{
			name:     "Happy flow. Beta reduction6",
			scenario: happyFlowBetaReduction6,
		},
		{
			name:     "Happy flow. Beta reduction7",
			scenario: happyFlowBetaReduction7,
		},
		{
			name:     "Happy flow. Beta reduction8",
			scenario: happyFlowBetaReduction8,
		},
		{
			name:     "Happy flow. Beta reduction9",
			scenario: happyFlowBetaReduction9,
		},
		{
			name:     "Happy flow. Beta reduction10",
			scenario: happyFlowBetaReduction10,
		},
		{
			name:     "Happy flow. Beta reduction11",
			scenario: happyFlowBetaReduction11,
		},
		{
			name:     "Happy flow. Beta reduction12",
			scenario: happyFlowBetaReduction12,
		},
	}

	t.Parallel()
	for _, test := range tests {
		t.Run(test.name, test.scenario)
	}
}

func happyFlowBetaReduction1(t *testing.T) {
	// arrange
	ctx := context.WithValue(context.Background(), "logger", logging.NewBuiltinLogger())
	automata := lexical_analysis.NewAutomata()
	analyzer := lexical_analysis.NewLexicalAnalyzer(ctx, automata)
	parser := NewLL1PredictableParser(ctx)

	// act
	expression := "(λy.y)_x"
	tk, _ := analyzer.Tokenize(expression)
	ast, err := parser.Parse(tk)
	ast, err = parser.BetaReduce(ast)
	res, err := parser.Unparse(ast)

	// assert
	assert.Equal(t, err, nil)
	assert.Equal(t, res, "x")
}

func happyFlowBetaReduction2(t *testing.T) {
	// arrange
	ctx := context.WithValue(context.Background(), "logger", logging.NewBuiltinLogger())
	automata := lexical_analysis.NewAutomata()
	analyzer := lexical_analysis.NewLexicalAnalyzer(ctx, automata)
	parser := NewLL1PredictableParser(ctx)

	// act
	expression := "(λy.y)_(λz.z)"
	tk, _ := analyzer.Tokenize(expression)
	ast, err := parser.Parse(tk)
	ast, err = parser.BetaReduce(ast)
	res, err := parser.Unparse(ast)
	// assert
	assert.Equal(t, err, nil)
	assert.Equal(t, res, "(λz.z)")
}

func happyFlowBetaReduction3(t *testing.T) {
	// arrange
	ctx := context.WithValue(context.Background(), "logger", logging.NewBuiltinLogger())
	automata := lexical_analysis.NewAutomata()
	analyzer := lexical_analysis.NewLexicalAnalyzer(ctx, automata)
	parser := NewLL1PredictableParser(ctx)

	// act
	expression := "(λy.y)_(λz.z)_t"
	tk, _ := analyzer.Tokenize(expression)
	ast, err := parser.Parse(tk)
	ast, err = parser.BetaReduce(ast)
	res, err := parser.Unparse(ast)
	// assert
	assert.Equal(t, err, nil)
	assert.Equal(t, res, "t")
}

func happyFlowBetaReduction4(t *testing.T) {
	// arrange
	ctx := context.WithValue(context.Background(), "logger", logging.NewBuiltinLogger())
	automata := lexical_analysis.NewAutomata()
	analyzer := lexical_analysis.NewLexicalAnalyzer(ctx, automata)
	parser := NewLL1PredictableParser(ctx)

	// act
	expression := "(λy.y_z_y)_t"
	tk, _ := analyzer.Tokenize(expression)
	ast, err := parser.Parse(tk)
	ast, err = parser.BetaReduce(ast)
	res, err := parser.Unparse(ast)
	// assert
	assert.Equal(t, err, nil)
	assert.Equal(t, res, "(t_(z_t))")
}

func happyFlowBetaReduction5(t *testing.T) {
	// arrange
	ctx := context.WithValue(context.Background(), "logger", logging.NewBuiltinLogger())
	automata := lexical_analysis.NewAutomata()
	analyzer := lexical_analysis.NewLexicalAnalyzer(ctx, automata)
	parser := NewLL1PredictableParser(ctx)

	// act
	expression := "(λy.z)_(λy.y)_x"
	tk, _ := analyzer.Tokenize(expression)
	ast, err := parser.Parse(tk)
	ast, err = parser.BetaReduce(ast)
	res, err := parser.Unparse(ast)
	// assert
	assert.Equal(t, err, nil)
	assert.Equal(t, res, "z")
}

func happyFlowBetaReduction6(t *testing.T) {
	// arrange
	ctx := context.WithValue(context.Background(), "logger", logging.NewBuiltinLogger())
	automata := lexical_analysis.NewAutomata()
	analyzer := lexical_analysis.NewLexicalAnalyzer(ctx, automata)
	parser := NewLL1PredictableParser(ctx)

	// act
	expression := "(λy.y)_(λa.a)_(λb.b)"
	tk, _ := analyzer.Tokenize(expression)
	ast, err := parser.Parse(tk)
	ast, err = parser.BetaReduce(ast)
	res, err := parser.Unparse(ast)
	// assert
	assert.Equal(t, err, nil)
	assert.Equal(t, res, "(λb.b)")
}

func happyFlowBetaReduction7(t *testing.T) {
	// arrange
	ctx := context.WithValue(context.Background(), "logger", logging.NewBuiltinLogger())
	automata := lexical_analysis.NewAutomata()
	analyzer := lexical_analysis.NewLexicalAnalyzer(ctx, automata)
	parser := NewLL1PredictableParser(ctx)

	// act
	expression := "(λy.y)_(λt.t_z)"
	tk, _ := analyzer.Tokenize(expression)
	ast, err := parser.Parse(tk)
	ast, err = parser.BetaReduce(ast)
	res, err := parser.Unparse(ast)
	// assert
	assert.Equal(t, err, nil)
	assert.Equal(t, res, "(λt.(t_z))")
}

func happyFlowBetaReduction8(t *testing.T) {
	// arrange
	ctx := context.WithValue(context.Background(), "logger", logging.NewBuiltinLogger())
	automata := lexical_analysis.NewAutomata()
	analyzer := lexical_analysis.NewLexicalAnalyzer(ctx, automata)
	parser := NewLL1PredictableParser(ctx)

	// act
	expression := "(λy.y)_x_y"
	tk, _ := analyzer.Tokenize(expression)
	ast, err := parser.Parse(tk)
	ast, err = parser.BetaReduce(ast)
	res, err := parser.Unparse(ast)
	// assert
	assert.Equal(t, err, nil)
	assert.Equal(t, res, "(x_y)")
}

func happyFlowBetaReduction9(t *testing.T) {
	// arrange
	ctx := context.WithValue(context.Background(), "logger", logging.NewBuiltinLogger())
	automata := lexical_analysis.NewAutomata()
	analyzer := lexical_analysis.NewLexicalAnalyzer(ctx, automata)
	parser := NewLL1PredictableParser(ctx)

	// act
	expression := "(λy.y)_λt.t_x"
	tk, _ := analyzer.Tokenize(expression)
	ast, err := parser.Parse(tk)
	ast, err = parser.BetaReduce(ast)
	res, err := parser.Unparse(ast)
	// assert
	assert.Equal(t, err, nil)
	assert.Equal(t, res, "(λt.(t_x))")
}

func happyFlowBetaReduction10(t *testing.T) {
	// arrange
	ctx := context.WithValue(context.Background(), "logger", logging.NewBuiltinLogger())
	automata := lexical_analysis.NewAutomata()
	analyzer := lexical_analysis.NewLexicalAnalyzer(ctx, automata)
	parser := NewLL1PredictableParser(ctx)

	// act
	expression := "(λy.y_k)_t"
	tk, _ := analyzer.Tokenize(expression)
	ast, err := parser.Parse(tk)
	ast, err = parser.BetaReduce(ast)
	res, err := parser.Unparse(ast)
	// assert
	assert.Equal(t, err, nil)
	assert.Equal(t, res, "(t_k)")
}

func happyFlowBetaReduction11(t *testing.T) {
	// arrange
	ctx := context.WithValue(context.Background(), "logger", logging.NewBuiltinLogger())
	automata := lexical_analysis.NewAutomata()
	analyzer := lexical_analysis.NewLexicalAnalyzer(ctx, automata)
	parser := NewLL1PredictableParser(ctx)

	// act
	expression := "(λy.y_k)_λz.z"
	tk, _ := analyzer.Tokenize(expression)
	ast, err := parser.Parse(tk)
	ast, err = parser.BetaReduce(ast)
	res, err := parser.Unparse(ast)
	// assert
	assert.Equal(t, err, nil)
	assert.Equal(t, res, "((λz.z)_k)")
}

func happyFlowBetaReduction12(t *testing.T) {
	// arrange
	ctx := context.WithValue(context.Background(), "logger", logging.NewBuiltinLogger())
	automata := lexical_analysis.NewAutomata()
	analyzer := lexical_analysis.NewLexicalAnalyzer(ctx, automata)
	parser := NewLL1PredictableParser(ctx)

	// act
	expression := "(λy.y_y_r)_((λy.y_z)_z)"
	tk, _ := analyzer.Tokenize(expression)
	ast, err := parser.Parse(tk)
	ast, err = parser.BetaReduce(ast)
	res, err := parser.Unparse(ast)
	// assert
	assert.Equal(t, err, nil)
	assert.Equal(t, res, "((z_z)_((z_z)_r))")
}
