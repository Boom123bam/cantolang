package evaluator

import (
	"bytes"
	"cantolang/object"
	"fmt"
)

var Builtins = map[string]object.BuiltInFunction{
	"有幾長": func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return Errorf("wrong number of arguments", "expected 1 arg got %d", len(args))
		}
		switch arg := args[0].(type) {
		case *object.Array:
			return &object.Integer{Value: len(arg.Items)}
		case *object.String:
			return &object.Integer{Value: len(arg.Value)}
		}
		return Errorf("invalid argument type", "%s", args[0].Type())
	},
	"講": func(args ...object.Object) object.Object {
		if len(args) == 0 {
			return Errorf("wrong number of arguments", "expected 1 or more got 0")
		}
		buff := bytes.Buffer{}
		for _, arg := range args {
			buff.WriteString(arg.Inspect() + " ")
		}
		fmt.Print(buff.String())
		return object.NULL
	},
	"加上": func(args ...object.Object) object.Object {
		if len(args) < 2 {
			return Errorf("wrong number of arguments", "expected 2 or more got 0")
		}
		if args[0].Type() != object.ARRAY_OBJ {
			return Errorf("invalid argument type", "expected array got %s", args[0].Type())
		}
		oldArr := args[0].(*object.Array).Items
		newArr := make([]object.Object, len(oldArr))
		copy(newArr, oldArr)
		for _, arg := range args[1:] {
			newArr = append(newArr, arg)
		}
		return &object.Array{Items: newArr}
	},
}
