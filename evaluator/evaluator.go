package evaluator

import (
	"cantolang/ast"
	"cantolang/object"
	"cantolang/token"
	"fmt"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		object.ERROR.Message = ""
		object.ERROR.Description = ""
		return EvalProgram(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case ast.Expression:
		return EvalExpression(node)
	case *ast.BlockStatement:
		return EvalProgram(node.Statements)
	case *ast.AssignStatement:
		return EvalAssignStatement(node)
	case *ast.FunctionDefStatment:
		return EvalFunctionDefStatement(node)
	case *ast.ReturnStatement:
		return &object.ReturnValue{Value: Eval(node.Expression)}
	default:
		return object.NULL
	}
}

func EvalFunctionDefStatement(statement *ast.FunctionDefStatment) object.Object {
	function := &object.Function{Parameters: statement.Parameters, Body: statement.Body}
	object.ENV.Set(statement.Identifier, function)
	return function
}

func EvalAssignStatement(statement *ast.AssignStatement) object.Object {
	val := Eval(statement.Expression)
	object.ENV.Set(statement.Identifier, val)
	return val
}

func EvalProgram(statements []ast.Statement) object.Object {
	var result object.Object
	for _, statement := range statements {
		result = Eval(statement)
		if result.Type() == object.RETURN_OBJ {
			return result.(*object.ReturnValue).Value
		}
		if object.ERROR.Message != "" {
			return object.ERROR
		}
	}
	return result
}

func EvalExpression(expression ast.Expression) object.Object {
	if object.ERROR.Message != "" {
		return object.ERROR
	}
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
	case *ast.Identifier:
		val, ok := object.ENV.Get(expression.Token.TokenLiteral)
		if ok {
			return val
		}
		return object.NULL
	case *ast.FunctionCallExpression:
		obj, ok := object.ENV.Get(expression.Identifier.Token.TokenLiteral)
		if !ok {
			return errorf("undefined variable", "%s is used before assignment", expression.Identifier.Token.TokenLiteral)
		}
		function, ok := obj.(*object.Function)
		if !ok {
			return errorf("type error", "%s expected function type got %T", expression.Identifier.Token.TokenLiteral, obj)
		}
		for i, p := range function.Parameters {
			object.ENV.Set(p.Token.TokenLiteral, Eval(expression.Parameters[i]))
		}
		return Eval(function.Body)

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
	if object.ERROR.Message != "" {
		return object.ERROR
	}
	if left.Type() != right.Type() {
		return errorf("type mismatch", "%T (%+v) %s %T (%+v)", left, left, infix.TokenLiteral, right, right)
	}
	l, l_ok := left.(*object.Integer)
	r, r_ok := right.(*object.Integer)
	switch infix.TokenType {
	case token.ADD:
		if l_ok && r_ok {
			return &object.Integer{Value: l.Value + r.Value}
		}
		return errorf("invalid operation", "%T (%+v) %s %T (%+v)", left, left, infix.TokenLiteral, right, right)
	case token.MINUS:
		if l_ok && r_ok {
			return &object.Integer{Value: l.Value - r.Value}
		}
		return errorf("invalid operation", "%T (%+v) %s %T (%+v)", left, left, infix.TokenLiteral, right, right)
	case token.MULTIPLY:
		if l_ok && r_ok {
			return &object.Integer{Value: l.Value * r.Value}
		}
		return errorf("invalid operation", "%T (%+v) %s %T (%+v)", left, left, infix.TokenLiteral, right, right)
	case token.DIVIDE:
		if l_ok && r_ok {
			return &object.Integer{Value: l.Value / r.Value}
		}
		return errorf("invalid operation", "%T (%+v) %s %T (%+v)", left, left, infix.TokenLiteral, right, right)
	case token.LESS_THAN:
		if l_ok && r_ok {
			return getBoolObj(l.Value < r.Value)
		}
		return errorf("invalid comparison", "%T (%+v) %s %T (%+v)", left, left, infix.TokenLiteral, right, right)
	case token.GREATER_THAN:
		if l_ok && r_ok {
			return getBoolObj(l.Value > r.Value)
		}
		return errorf("invalid comparison", "%T (%+v) %s %T (%+v)", left, left, infix.TokenLiteral, right, right)
	case token.EQUAL_TO:
		if l_ok && r_ok {
			return getBoolObj(l.Value == r.Value)
		}
		if left.Type() == object.BOOL_OBJ && right.Type() == object.BOOL_OBJ {
			return getBoolObj(left == right)
		}
		return errorf("invalid comparison", "%T (%+v) %s %T (%+v)", left, left, infix.TokenLiteral, right, right)
	}
	// invalid infix
	return errorf("invalid infix", "%T (%+v) %s %T (%+v)", left, left, infix.TokenLiteral, right, right)
}

func evalPrefixExpression(tokenType string, right object.Object) object.Object {
	switch tokenType {
	case token.MINUS:
		rightInt, ok := right.(*object.Integer)
		if !ok {
			return errorf("invalid prefix", "%s %T (%+v)", tokenType, right, right)
		}
		return &object.Integer{Value: -rightInt.Value}
	case token.NOT:
		rightBool, ok := right.(*object.Boolean)
		if !ok {
			return errorf("invalid prefix", "%s %T (%+v)", tokenType, right, right)
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

func errorf(message, descrigtionFormat string, a ...interface{}) *object.Error {
	object.ERROR.Message = message
	object.ERROR.Description = fmt.Sprintf(descrigtionFormat, a...)
	return object.ERROR
}
