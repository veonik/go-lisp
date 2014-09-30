package lisp

import "errors"

var executionLimit = 0
var currentDepth = 0
var halt = errors.New("Execution limit exceeded")

func SetExecutionLimit(n int) {
	executionLimit = n
}

func EvalString(line string) (Value, error) {
	currentDepth = 0
	expanded, err := NewTokens(line).Expand()
	if err != nil {
		return Nil, err
	}
	parsed, err := expanded.Parse()
	if err != nil {
		return Nil, err
	}
	evaled, err := parsed.Eval()
	if err != nil {
		return Nil, err
	}
	scope.Create("_", evaled)
	return evaled, nil
}
