package main

import (
	"context"
	"flag"
	"fmt"
	"math-parser/pkg/lexical_analysis"
	syntactical_analyzer "math-parser/pkg/syntactical_analysis"
	"math-parser/pkg/utils/logging"
)

func main() {
	ctx := context.WithValue(context.Background(), "logger", logging.NewBuiltinLogger())

	//preprocessor := preprocessing.NewPreprocessing()
	automata := lexical_analysis.NewAutomata()
	lexicalAnalyzer := lexical_analysis.NewLexicalAnalyzer(ctx, automata)
	syntacticalAnalyzer := syntactical_analyzer.NewLL1PredictableParser(ctx)
	//compiler := compilation.NewCompiler(preprocessor, lexicalAnalyzer, syntacticalAnalyzer)

	var expr string
	flag.StringVar(&expr, "expr", "", "expression")
	flag.Parse()

	tk, err := lexicalAnalyzer.Tokenize(expr)
	if err != nil {
		fmt.Printf("error: %s", err)
	}

	_, err = syntacticalAnalyzer.Parse(tk)
	if err != nil {
		fmt.Printf("error: %s", err)
	}
}
