package lisp

import "fmt"
import "reflect"

type Builtin struct{}

var builtin = Builtin{}

var builtin_commands = map[string]string{
	"+":       "Add",
	"-":       "Sub",
	"*":       "Mul",
	">":       "Gt",
	"<":       "Lt",
	">=":      "Gte",
	"<=":      "Lte",
	"display": "Display",
}

func isBuiltin(c Value) bool {
	s := c.String()
	if _, ok := builtin_commands[s]; ok {
		return true
	}
	return false
}

func runBuiltin(expr Sexp) (val Value, err error) {
	cmd := builtin_commands[expr[0].String()]
	values := []reflect.Value{}
	for _, i := range expr[1:] {
		if value, err := evalValue(i); err != nil {
			return Nil, err
		} else {
			values = append(values, reflect.ValueOf(value))
		}
	}
	result := reflect.ValueOf(&builtin).MethodByName(cmd).Call(values)
	val = result[0].Interface().(Value)
	err, _ = result[1].Interface().(error)
	return
}

func (Builtin) Display(vars ...Value) (Value, error) {
	var interfaces []interface{}
	for _, v := range vars {
		interfaces = append(interfaces, v)
	}
	fmt.Println(interfaces...)
	return Nil, nil
}

func (Builtin) Add(vars ...Value) (Value, error) {
	var sum int
	for _, v := range vars {
		if v.IsA(NumberKind) {
			sum += v.Number()
		} else {
			return Nil, fmt.Errorf("Badly formatted arguments: %v", vars)
		}
	}
	return NewValue(sum), nil
}

func (Builtin) Sub(vars ...Value) (Value, error) {
	if !vars[0].IsA(NumberKind) {
		return Nil, fmt.Errorf("Badly formatted arguments: %v", vars)
	}
	sum := vars[0].Number()
	for _, v := range vars[1:] {
		if v.IsA(NumberKind) {
			sum -= v.Number()
		} else {
			return Nil, fmt.Errorf("Badly formatted arguments: %v", vars)
		}
	}
	return NewValue(sum), nil
}

func (Builtin) Mul(vars ...Value) (Value, error) {
	if !vars[0].IsA(NumberKind) {
		return Nil, fmt.Errorf("Badly formatted arguments: %v", vars)
	}
	sum := vars[0].Number()
	for _, v := range vars[1:] {
		if v.IsA(NumberKind) {
			sum *= v.Number()
		} else {
			return Nil, fmt.Errorf("Badly formatted arguments: %v", vars)
		}
	}
	return NewValue(sum), nil
}

func (Builtin) Gt(vars ...Value) (Value, error) {
	for i := 1; i < len(vars); i++ {
		v1 := vars[i-1]
		v2 := vars[i]
		if !v1.IsA(NumberKind) || !v2.IsA(NumberKind) {
			return Nil, fmt.Errorf("Badly formatted arguments: %v", vars)
		} else if !(v1.Number() > v2.Number()) {
			return NewValue(false), nil
		}
	}
	return NewValue(true), nil
}

func (Builtin) Lt(vars ...Value) (Value, error) {
	for i := 1; i < len(vars); i++ {
		v1 := vars[i-1]
		v2 := vars[i]
		if !v1.IsA(NumberKind) || !v2.IsA(NumberKind) {
			return Nil, fmt.Errorf("Badly formatted arguments: %v", vars)
		} else if !(v1.Number() < v2.Number()) {
			return NewValue(false), nil
		}
	}
	return NewValue(true), nil
}

func (Builtin) Gte(vars ...Value) (Value, error) {
	for i := 1; i < len(vars); i++ {
		v1 := vars[i-1]
		v2 := vars[i]
		if !v1.IsA(NumberKind) || !v2.IsA(NumberKind) {
			return Nil, fmt.Errorf("Badly formatted arguments: %v", vars)
		} else if !(v1.Number() >= v2.Number()) {
			return NewValue(false), nil
		}
	}
	return NewValue(true), nil
}

func (Builtin) Lte(vars ...Value) (Value, error) {
	for i := 1; i < len(vars); i++ {
		v1 := vars[i-1]
		v2 := vars[i]
		if !v1.IsA(NumberKind) || !v2.IsA(NumberKind) {
			return Nil, fmt.Errorf("Badly formatted arguments: %v", vars)
		} else if !(v1.Number() <= v2.Number()) {
			return NewValue(false), nil
		}
	}
	return NewValue(true), nil
}
