package evaluator

import "cantolang/object"

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
}
