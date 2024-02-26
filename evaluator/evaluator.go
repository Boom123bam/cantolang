package evaluator

import (
	"cantolang/ast"
	"cantolang/object"
	"cantolang/token"
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
		return getBoolObj(expression.Value)
	case *ast.PrefixExpression:
		right := Eval(expression.Right)
		return evalPrefixExpression(expression.PrefixToken.TokenType, right)
	default:
		return object.NULL
	}
}

func evalPrefixExpression(tokenType string, right object.Object) object.Object {
	switch tokenType {
	case token.MINUS:
		rightInt, ok := right.(*object.Integer)
		if !ok {
			return object.NULL
		}
		return &object.Integer{Value: -rightInt.Value}
	case token.NOT:
		rightBool, ok := right.(*object.Boolean)
		if !ok {
			return object.NULL
		}
		return getBoolObj(!rightBool.Value)
	default:
		return object.NULL
	}
}

func getBoolObj(b bool) *object.Boolean {
	if b {
		return object.TRUE
	}
	return object.FALSE
}
