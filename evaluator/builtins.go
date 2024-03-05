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
		fmt.Println(buff.String())
		return object.NULL
	},
}
