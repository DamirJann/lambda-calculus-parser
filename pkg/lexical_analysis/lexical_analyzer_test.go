package lexical_analysis

import (
	"context"
	"gotest.tools/assert"
	"math-parser/pkg/entity"
	"math-parser/pkg/utils/logging"
	"testing"
)

func TestLexicalAnalyzer_Tokenize(t *testing.T) {
	var tests = []struct {
		name     string
		scenario func(*testing.T)
	}{
		{
			name:     "Happy flow. Process with basic operations",
			scenario: happyFlowTokenizeWithBasicOperations,
		},
	}

	t.Parallel()
	for _, test := range tests {
		t.Run(test.name, test.scenario)
	}
}

func happyFlowTokenizeWithBasicOperations(t *testing.T) {
	// arrange
	ctx := context.WithValue(context.Background(), "logger", logging.NewBuiltinLogger())
	automata := NewAutomata()
	lexicalAnalyzer := NewLexicalAnalyzer(ctx, automata)
	expression := "(λf.(λx.f(xx))_(λx.f(xx)))"

	// act
	ts, err := lexicalAnalyzer.Tokenize(expression)

	// assert
	assert.Equal(t, err, nil)
	assert.Equal(t, len(ts), 26)
	assert.Equal(t, ts[0].Tag, entity.LEFT_BRACKET)
	assert.Equal(t, ts[1].Tag, entity.LAMBDA)
	assert.Equal(t, ts[2].Tag, entity.VARIABLE)
	assert.Equal(t, ts[3].Tag, entity.ABSTRACTION)
	assert.Equal(t, ts[4].Tag, entity.LEFT_BRACKET)
	assert.Equal(t, ts[5].Tag, entity.LAMBDA)
	assert.Equal(t, ts[6].Tag, entity.VARIABLE)
	assert.Equal(t, ts[7].Tag, entity.ABSTRACTION)
	assert.Equal(t, ts[8].Tag, entity.VARIABLE)
	assert.Equal(t, ts[9].Tag, entity.LEFT_BRACKET)
	assert.Equal(t, ts[10].Tag, entity.VARIABLE)
	assert.Equal(t, ts[11].Tag, entity.VARIABLE)
	assert.Equal(t, ts[12].Tag, entity.RIGHT_BRACKET)
	assert.Equal(t, ts[13].Tag, entity.RIGHT_BRACKET)
	assert.Equal(t, ts[14].Tag, entity.APPLICATION)
	assert.Equal(t, ts[15].Tag, entity.LEFT_BRACKET)
	assert.Equal(t, ts[16].Tag, entity.LAMBDA)
	assert.Equal(t, ts[17].Tag, entity.VARIABLE)
	assert.Equal(t, ts[18].Tag, entity.ABSTRACTION)
	assert.Equal(t, ts[19].Tag, entity.VARIABLE)
	assert.Equal(t, ts[20].Tag, entity.LEFT_BRACKET)
	assert.Equal(t, ts[21].Tag, entity.VARIABLE)
	assert.Equal(t, ts[22].Tag, entity.VARIABLE)
	assert.Equal(t, ts[23].Tag, entity.RIGHT_BRACKET)
	assert.Equal(t, ts[24].Tag, entity.RIGHT_BRACKET)
	assert.Equal(t, ts[25].Tag, entity.RIGHT_BRACKET)
}
