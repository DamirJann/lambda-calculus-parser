package syntactical_analyzer

import (
	"context"
	"gotest.tools/assert"
	"math-parser/pkg/lexical_analysis"
	"math-parser/pkg/utils/logging"
	"testing"
)

func TestLexicalAnalyzer_Tokenize(t *testing.T) {
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
