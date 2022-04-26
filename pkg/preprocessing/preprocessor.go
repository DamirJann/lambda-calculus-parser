package preprocessing

import (
	"context"
	"math-parser/pkg/utils/logging"
)

const (
	SPACE string = " "
	EMPTY string = ""
)

func NewPreprocessing() Preprocessor {
	return preprocessor{
		logging: context.WithValue(context.Background(), "preprocessor", logging.NewBuiltinLogger()),
	}
}

type Preprocessor interface {
	Process(string) (output string)
}

type preprocessor struct {
	logging context.Context
}

func (preprocessor) Process(input string) (output string) {
	return output
}
