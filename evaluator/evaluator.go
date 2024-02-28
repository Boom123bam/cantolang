package evaluator

import (
	"cantolang/ast"
	"cantolang/object"
	"cantolang/token"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return EvalProgram(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case ast.Expression:
		return EvalExpression(node)
	case *ast.BlockStatement:
		return EvalProgram(node.Statements)
	default:
		return object.NULL
	}
}

func EvalProgram(statements []ast.Statement) object.Object {
	var result object.Object
	for _, statement := range statements {
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
	case *ast.InfixExpression:
		left := Eval(expression.Left)
		right := Eval(expression.Right)
		return evalInfixExpression(left, right, expression.Infix)
	case *ast.IfExpression:
		condition := Eval(expression.Condition)
		if isTruthy(condition) {
			return Eval(expression.Consequence)
		}
		return Eval(expression.Alternative)

	default:
		return object.NULL
	}
}

func isTruthy(obj object.Object) bool {
	switch obj := obj.(type) {
	case *object.Boolean:
		if !obj.Value {
			return false
		}
	case *object.Null:
		return false
	}
	return true
}

func evalInfixExpression(left object.Object, right object.Object, infix token.Token) object.Object {
	// + - * / 係 細過 大過
	l, l_ok := left.(*object.Integer)
	r, r_ok := right.(*object.Integer)
	switch infix.TokenType {
	case token.ADD:
		if l_ok && r_ok {
			return &object.Integer{Value: l.Value + r.Value}
		}
	case token.MINUS:
		if l_ok && r_ok {
			return &object.Integer{Value: l.Value - r.Value}
		}
	case token.MULTIPLY:
		if l_ok && r_ok {
			return &object.Integer{Value: l.Value * r.Value}
		}
	case token.DIVIDE:
		if l_ok && r_ok {
			return &object.Integer{Value: l.Value / r.Value}
		}
	case token.LESS_THAN:
		if l_ok && r_ok {
			return getBoolObj(l.Value < r.Value)
		}
	case token.GREATER_THAN:
		if l_ok && r_ok {
			return getBoolObj(l.Value > r.Value)
		}
	case token.EQUAL_TO:
		if l_ok && r_ok {
			return getBoolObj(l.Value == r.Value)
		}
		lBool, ok := left.(*object.Boolean)
		if !ok {
			return object.NULL
		}
		rBool, ok := right.(*object.Boolean)
		if !ok {
			return object.NULL
		}
		return getBoolObj(rBool == lBool)
	}
	// invalid infix
	return object.NULL
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
