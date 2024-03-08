package evaluator

import (
	"cantolang/ast"
	"cantolang/object"
	"cantolang/token"
	"fmt"
)

func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		object.ERROR.Message = ""
		object.ERROR.Description = ""
		return EvalStatements(node.Statements, env, true)
	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)
	case ast.Expression:
		return EvalExpression(node, env)
	case *ast.AssignStatement:
		return EvalAssignStatement(node, env)
	case *ast.FunctionDefStatment:
		return EvalFunctionDefStatement(node, env)
	case *ast.ReturnStatement:
		return &object.ReturnValue{Value: Eval(node.Expression, env)}
	case *ast.IncrementDecrementStatement:
		val, ok := env.Get(node.Identifier)
		if !ok {
			return Errorf("undefined variable", "%s is used before assignment", node.Identifier)
		}
		switch val := val.(type) {
		case *object.Integer:
			if node.IsIncrement {
				env.Set(node.Identifier, &object.Integer{Value: val.Value + 1})
			} else {
				env.Set(node.Identifier, &object.Integer{Value: val.Value - 1})
			}
			return object.NULL
		default:
			return Errorf("type error", "cannot increment %s", val.Type())
		}
	default:
		return object.NULL
	}
}

func EvalFunctionDefStatement(statement *ast.FunctionDefStatment, env *object.Environment) object.Object {
	function := &object.Function{Parameters: statement.Parameters, Body: statement.Body}
	env.Set(statement.Identifier, function)
	return function
}

func EvalAssignStatement(statement *ast.AssignStatement, env *object.Environment) object.Object {
	val := Eval(statement.Expression, env)
	env.Set(statement.Identifier, val)
	return val
}

func EvalStatements(statements []ast.Statement, env *object.Environment, unwrapReturn bool) object.Object {
	var result object.Object
	for _, statement := range statements {
		result = Eval(statement, env)
		if result.Type() == object.RETURN_OBJ {
			if unwrapReturn {
				return result.(*object.ReturnValue).Value
			}
			return result
		}
		if object.ERROR.Message != "" {
			return object.ERROR
		}
	}
	return result
}

func EvalExpression(expression ast.Expression, env *object.Environment) object.Object {
	if object.ERROR.Message != "" {
		return object.ERROR
	}
	switch expression := expression.(type) {
	case *ast.IntegerLiteral:
		return &object.Integer{Value: expression.Value}
	case *ast.StringLiteral:
		return &object.String{Value: expression.Value}
	case *ast.ArrayLiteral:
		arr := &object.Array{}
		for _, item := range expression.Items {
			res := Eval(item, env)
			if object.ERROR.Message != "" {
				return object.ERROR
			}
			arr.Items = append(arr.Items, res)
		}
		return arr
	case *ast.IndexExpression:
		idxObj := EvalExpression(expression.Index, env)
		if object.ERROR.Message != "" {
			return object.ERROR
		}
		idx, ok := idxObj.(*object.Integer)
		if !ok {
			return Errorf("type error", "index must be number")
		}
		left := Eval(expression.Left, env)
		if object.ERROR.Message != "" {
			return object.ERROR
		}
		switch left := left.(type) {
		case *object.Array:
			if idx.Value < 0 || idx.Value >= len(left.Items) {
				return Errorf("index error", "list index out of range")
			}
			return left.Items[idx.Value]
		case *object.String:
			if idx.Value < 0 || idx.Value >= len(left.Value) {
				return Errorf("index error", "string index out of range")
			}
			return &object.String{Value: string([]rune(left.Value)[idx.Value])}

		default:
			return Errorf("type error", "cannot index %s", left.Type())
		}
	case *ast.Boolean:
		return getBoolObj(expression.Value)
	case *ast.PrefixExpression:
		right := Eval(expression.Right, env)
		return evalPrefixExpression(expression.PrefixToken.TokenType, right)
	case *ast.InfixExpression:
		left := Eval(expression.Left, env)
		right := Eval(expression.Right, env)
		return evalInfixExpression(left, right, expression.Infix)
	case *ast.IfExpression:
		condition := Eval(expression.Condition, env)
		if isTruthy(condition) {
			return EvalStatements(expression.Consequence.Statements, env, false)
		}
		return EvalStatements(expression.Alternative.Statements, env, false)
	case *ast.WhileLoop:
		condition := Eval(expression.Condition, env)
		var result object.Object
		result = object.NULL
		for isTruthy(condition) {
			result = EvalStatements(expression.Body.Statements, env, false)
			if result.Type() == object.RETURN_OBJ {
				return result
			}
			condition = Eval(expression.Condition, env)
		}
		return result
	case *ast.Identifier:
		val, ok := env.Get(expression.Token.TokenLiteral)
		if ok {
			return val
		}
		return object.NULL
	case *ast.FunctionCallExpression:
		if obj, ok := env.Get(expression.Identifier.Token.TokenLiteral); ok {
			function, ok := obj.(*object.Function)
			if !ok {
				return Errorf("type error", "%s expected function type got %T", expression.Identifier.Token.TokenLiteral, obj)
			}
			childEnv := object.NewEnvironment(env)
			for i, p := range function.Parameters {
				childEnv.Set(p.Token.TokenLiteral, Eval(expression.Parameters[i], childEnv))
			}
			return EvalStatements(function.Body.Statements, childEnv, true)
		} else if builtin, ok := Builtins[expression.Identifier.Token.TokenLiteral]; ok {
			params := []object.Object{}
			for _, param := range expression.Parameters {
				res := Eval(param, env)
				if res.Type() == object.ERROR_OBJ {
					return res
				}
				params = append(params, res)
			}
			return builtin(params...)
		}
		return Errorf("undefined variable", "%s is used before assignment", expression.Identifier.Token.TokenLiteral)

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
		return Errorf("type mismatch", "%T (%+v) %s %T (%+v)", left, left, infix.TokenLiteral, right, right)
	}
	lInt, l_ok := left.(*object.Integer)
	rInt, r_ok := right.(*object.Integer)
	switch infix.TokenType {
	case token.ADD:
		if l_ok && r_ok {
			return &object.Integer{Value: lInt.Value + rInt.Value}
		}
		if left.Type() == object.STRING_OBJ {
			return &object.String{Value: left.(*object.String).Value + right.(*object.String).Value}
		}
		return Errorf("invalid operation", "%T (%+v) %s %T (%+v)", left, left, infix.TokenLiteral, right, right)
	case token.MINUS:
		if l_ok && r_ok {
			return &object.Integer{Value: lInt.Value - rInt.Value}
		}
		return Errorf("invalid operation", "%T (%+v) %s %T (%+v)", left, left, infix.TokenLiteral, right, right)
	case token.MULTIPLY:
		if l_ok && r_ok {
			return &object.Integer{Value: lInt.Value * rInt.Value}
		}
		return Errorf("invalid operation", "%T (%+v) %s %T (%+v)", left, left, infix.TokenLiteral, right, right)
	case token.DIVIDE:
		if l_ok && r_ok {
			return &object.Integer{Value: lInt.Value / rInt.Value}
		}
		return Errorf("invalid operation", "%T (%+v) %s %T (%+v)", left, left, infix.TokenLiteral, right, right)
	case token.LESS_THAN:
		if l_ok && r_ok {
			return getBoolObj(lInt.Value < rInt.Value)
		}
		return Errorf("invalid comparison", "%T (%+v) %s %T (%+v)", left, left, infix.TokenLiteral, right, right)
	case token.GREATER_THAN:
		if l_ok && r_ok {
			return getBoolObj(lInt.Value > rInt.Value)
		}
		return Errorf("invalid comparison", "%T (%+v) %s %T (%+v)", left, left, infix.TokenLiteral, right, right)
	case token.EQUAL_TO:
		if l_ok && r_ok {
			return getBoolObj(lInt.Value == rInt.Value)
		}
		if left.Type() == object.STRING_OBJ {
			return getBoolObj(left.(*object.String).Value == right.(*object.String).Value)
		}
		if left.Type() == object.BOOL_OBJ {
			return getBoolObj(left == right)
		}
		return Errorf("invalid comparison", "%T (%+v) %s %T (%+v)", left, left, infix.TokenLiteral, right, right)
	}
	// invalid infix
	return Errorf("invalid infix", "%T (%+v) %s %T (%+v)", left, left, infix.TokenLiteral, right, right)
}

func evalPrefixExpression(tokenType string, right object.Object) object.Object {
	switch tokenType {
	case token.MINUS:
		rightInt, ok := right.(*object.Integer)
		if !ok {
			return Errorf("invalid prefix", "%s %T (%+v)", tokenType, right, right)
		}
		return &object.Integer{Value: -rightInt.Value}
	case token.NOT:
		rightBool, ok := right.(*object.Boolean)
		if !ok {
			return Errorf("invalid prefix", "%s %T (%+v)", tokenType, right, right)
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

func Errorf(message, descrigtionFormat string, a ...interface{}) *object.Error {
	object.ERROR.Message = message
	object.ERROR.Description = fmt.Sprintf(descrigtionFormat, a...)
	return object.ERROR
}
