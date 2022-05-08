package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"math-parser/pkg/lexical_analysis"
	syntactical_analyzer "math-parser/pkg/syntactical_analysis"
	"math-parser/pkg/utils/logging"
)

func main() {
	ctx := context.WithValue(context.Background(), "logger", logging.NewBuiltinLogger())

	automata := lexical_analysis.NewAutomata()
	lexicalAnalyzer := lexical_analysis.NewLexicalAnalyzer(ctx, automata)
	syntacticalAnalyzer := syntactical_analyzer.NewLL1PredictableParser(ctx)

	var expr string
	flag.StringVar(&expr, "expr", "", "expression")

	var red string
	flag.StringVar(&red, "red", "", "reduction")
	flag.Parse()

	tk, err := lexicalAnalyzer.Tokenize(expr)
	if err != nil {
		fmt.Printf("error: %s", err)
		return
	}

	ast, err := syntacticalAnalyzer.Parse(tk)
	if err != nil {
		fmt.Printf("error: %s", err)
		return
	}

	if res, err := syntacticalAnalyzer.Unparse(ast); err != nil {
		fmt.Printf("Unparsed to %s", res)
		return
	}

	switch red {
	case "alpha":
		{
			err = errors.New("not implemented")
		}
	case "beta":
		{
			ast, err = syntacticalAnalyzer.BetaReduce(ast)
			if err != nil {
				break
			}
			res, err := syntacticalAnalyzer.Unparse(ast)
			if err != nil {
				break
			}
			fmt.Printf("Unparsed after beta-reduction to %s", res)

		}
	}

	if err != nil {
		fmt.Printf("Error during comand: %v", err)
		return
	}
}
