package lisp

import (
	"fmt"
	"strconv"
	"strings"
)

type Builtin struct{}

var builtin = Builtin{}

var builtin_commands = map[string]string{
	"+":              "Add",
	"-":              "Sub",
	"*":              "Mul",
	">":              "Gt",
	"<":              "Lt",
	">=":             "Gte",
	"<=":             "Lte",
	"display":        "Display",
	"cons":           "Cons",
	"car":            "Car",
	"cdr":            "Cdr",
	"string-val":     "StringVal",
	"number-val":     "NumberVal",
	"string?":        "StringHuh",
	"string=?":       "StringEqualHuh",
	"string-length":  "StringLength",
	"string-append":  "StringAppend",
	"string-start=?": "StringStartEqualHuh",
	"string-end=?":   "StringEndEqualHuh",
	"string-match?":  "StringMatchHuh",
	"string-index":   "StringIndex",
	"string-first":   "StringFirst",
	"string-last":    "StringLast",
}

type BuiltinHandler func(vars ...Value) (Value, error)

var builtin_handlers = make(map[string]BuiltinHandler)

func SetHandler(name string, handler BuiltinHandler) {
	builtin_handlers[name] = handler
}

func (Builtin) Display(vars ...Value) (Value, error) {
	if len(vars) == 1 {
		fmt.Println(vars[0])
	} else {
		return badlyFormattedArguments(vars)
	}
	return Nil, nil
}

func (Builtin) Cons(vars ...Value) (Value, error) {
	if len(vars) == 2 {
		cons := Cons{&vars[0], &vars[1]}
		return Value{consValue, &cons}, nil
	} else {
		return badlyFormattedArguments(vars)
	}
}

func (Builtin) Car(vars ...Value) (Value, error) {
	if len(vars) == 1 && vars[0].typ == consValue {
		cons := vars[0].Cons()
		return *cons.car, nil
	} else {
		return badlyFormattedArguments(vars)
	}
}

func (Builtin) Cdr(vars ...Value) (Value, error) {
	if len(vars) == 1 && vars[0].typ == consValue {
		cons := vars[0].Cons()
		return *cons.cdr, nil
	} else {
		return badlyFormattedArguments(vars)
	}
}

func (Builtin) Add(vars ...Value) (Value, error) {
	sum := 0.0
	for _, v := range vars {
		if v.typ == numberValue {
			sum += v.Number()
		} else {
			return badlyFormattedArguments(vars)
		}
	}
	return Value{numberValue, sum}, nil
}

func (Builtin) Sub(vars ...Value) (Value, error) {
	if len(vars) == 0 || vars[0].typ != numberValue {
		return badlyFormattedArguments(vars)
	}
	sum := vars[0].Number()
	for _, v := range vars[1:] {
		if v.typ == numberValue {
			sum -= v.Number()
		} else {
			return badlyFormattedArguments(vars)
		}
	}
	return Value{numberValue, sum}, nil
}

func (Builtin) Mul(vars ...Value) (Value, error) {
	if len(vars) == 0 {
		return Value{numberValue, 1.0}, nil
	}
	if vars[0].typ != numberValue {
		return badlyFormattedArguments(vars)
	}
	sum := vars[0].Number()
	for _, v := range vars[1:] {
		if v.typ == numberValue {
			sum *= v.Number()
		} else {
			return badlyFormattedArguments(vars)
		}
	}
	return Value{numberValue, sum}, nil
}

func (Builtin) Gt(vars ...Value) (Value, error) {
	if len(vars) == 0 {
		return badlyFormattedArguments(vars)
	}
	for i := 1; i < len(vars); i++ {
		v1 := vars[i-1]
		v2 := vars[i]
		if v1.typ != numberValue || v2.typ != numberValue {
			return badlyFormattedArguments(vars)
		} else if !(v1.Number() > v2.Number()) {
			return False, nil
		}
	}
	return True, nil
}

func (Builtin) Lt(vars ...Value) (Value, error) {
	if len(vars) == 0 {
		return badlyFormattedArguments(vars)
	}
	for i := 1; i < len(vars); i++ {
		v1 := vars[i-1]
		v2 := vars[i]
		if v1.typ != numberValue || v2.typ != numberValue {
			return badlyFormattedArguments(vars)
		} else if !(v1.Number() < v2.Number()) {
			return False, nil
		}
	}
	return True, nil
}

func (Builtin) Gte(vars ...Value) (Value, error) {
	if len(vars) == 0 {
		return badlyFormattedArguments(vars)
	}
	for i := 1; i < len(vars); i++ {
		v1 := vars[i-1]
		v2 := vars[i]
		if v1.typ != numberValue || v2.typ != numberValue {
			return badlyFormattedArguments(vars)
		} else if !(v1.Number() >= v2.Number()) {
			return False, nil
		}
	}
	return True, nil
}

func (Builtin) Lte(vars ...Value) (Value, error) {
	if len(vars) == 0 {
		return badlyFormattedArguments(vars)
	}
	for i := 1; i < len(vars); i++ {
		v1 := vars[i-1]
		v2 := vars[i]
		if v1.typ != numberValue || v2.typ != numberValue {
			return badlyFormattedArguments(vars)
		} else if !(v1.Number() <= v2.Number()) {
			return False, nil
		}
	}
	return True, nil
}

func (Builtin) StringVal(vars ...Value) (Value, error) {
	if len(vars) != 1 {
		return badlyFormattedArguments(vars)
	}
	switch {
	case vars[0].typ == stringValue:
		return vars[0], nil

	case vars[0].typ == numberValue:
		return StringValue(strconv.FormatFloat(vars[0].Number(), 'f', 10, 64)), nil
	}
	return Nil, nil
}

func (Builtin) NumberVal(vars ...Value) (Value, error) {
	if len(vars) != 1 {
		return badlyFormattedArguments(vars)
	}
	switch {
	case vars[0].typ == stringValue:
		val, _ := strconv.ParseFloat(vars[0].String(), 64)
		return NumberValue(val), nil

	case vars[0].typ == numberValue:
		return vars[0], nil
	}
	return Nil, nil
}

func (Builtin) StringHuh(vars ...Value) (Value, error) {
	if len(vars) != 1 {
		return badlyFormattedArguments(vars)
	}
	if vars[0].typ == stringValue {
		return True, nil
	}
	return False, nil
}

func (Builtin) StringEqualHuh(vars ...Value) (Value, error) {
	if len(vars) < 2 {
		return badlyFormattedArguments(vars)
	}
	for i := 1; i < len(vars); i++ {
		v1 := vars[i-1]
		v2 := vars[i]
		if v1.typ != stringValue || v2.typ != stringValue {
			return badlyFormattedArguments(vars)
		} else if v1.String() != v2.String() {
			return False, nil
		}
	}
	return True, nil
}

func (Builtin) StringLength(vars ...Value) (Value, error) {
	if len(vars) != 1 || vars[0].typ != stringValue {
		return badlyFormattedArguments(vars)
	}
	return Value{numberValue, float64(len(vars[0].String()))}, nil
}

func (Builtin) StringAppend(vars ...Value) (Value, error) {
	result := ""
	for i := 0; i < len(vars); i++ {
		v := vars[i]
		if v.typ != stringValue {
			return badlyFormattedArguments(vars)
		} else {
			result = result + v.String()
		}
	}
	return Value{stringValue, result}, nil
}

func (Builtin) StringStartEqualHuh(vars ...Value) (Value, error) {
	if len(vars) != 2 {
		return badlyFormattedArguments(vars)
	}
	if strings.HasPrefix(vars[0].String(), vars[1].String()) {
		return True, nil
	}
	return False, nil
}

func (Builtin) StringEndEqualHuh(vars ...Value) (Value, error) {
	if len(vars) != 2 {
		return badlyFormattedArguments(vars)
	}
	if strings.HasSuffix(vars[0].String(), vars[1].String()) {
		return True, nil
	}
	return False, nil
}

func (Builtin) StringMatchHuh(vars ...Value) (Value, error) {
	if len(vars) != 2 {
		return badlyFormattedArguments(vars)
	}
	if strings.Contains(vars[0].String(), vars[1].String()) {
		return True, nil
	}
	return False, nil
}

func (Builtin) StringIndex(vars ...Value) (Value, error) {
	if len(vars) != 2 {
		return badlyFormattedArguments(vars)
	}

	return NumberValue(float64(strings.Index(vars[0].String(), vars[1].String()))), nil
}

func (Builtin) StringFirst(vars ...Value) (Value, error) {
	if len(vars) != 2 {
		return badlyFormattedArguments(vars)
	}

	n := int(vars[1].Number())
	val := vars[0].String()
	if n < 0 {
		return StringValue(val[0 : len(val)+n]), nil
	}

	return StringValue(val[0:n]), nil
}

func (Builtin) StringLast(vars ...Value) (Value, error) {
	if len(vars) != 2 {
		return badlyFormattedArguments(vars)
	}

	n := int(vars[1].Number())
	val := vars[0].String()
	if n < 0 {
		return StringValue(val[-n:]), nil
	}
	return StringValue(val[len(val)-n:]), nil
}

func badlyFormattedArguments(vars []Value) (Value, error) {
	return Nil, fmt.Errorf("Badly formatted arguments: %v", vars)
}
