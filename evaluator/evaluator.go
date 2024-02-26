package evaluator

import (
	"cantolang/ast"
	"cantolang/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return EvalProgram(node)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case ast.Expression:
		return EvalExpression(node)
	default:
		return object.NULL
	}
}

func EvalProgram(program *ast.Program) object.Object {
	var result object.Object
	for _, statement := range program.Statements {
		result = Eval(statement)
	}
	return result
}

func EvalExpression(expression ast.Expression) object.Object {
	switch expression := expression.(type) {
	case *ast.IntegerLiteral:
		return &object.Integer{Value: expression.Value}
	case *ast.Boolean:
		if expression.Value {
			return object.TRUE
		}
		return object.FALSE
	default:
		return object.NULL
	}
}
