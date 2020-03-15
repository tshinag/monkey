package evaluator

import (
	"fmt"

	"github.com/tshinag/monkey/ast"
	"github.com/tshinag/monkey/object"
)

var (
	// NULL is the instance of "null" literal
	NULL = &object.Null{}
	// TRUE is the instance of "true" literal
	TRUE = &object.Boolean{Value: true}
	// FALSE is the instance of "false" literal
	FALSE = &object.Boolean{Value: false}
)

// Eval evaluates ast.Node
func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {
	// statements
	case *ast.Program:
		return evalProgram(node, env)
	case *ast.BlockStatement:
		return evalBlockStatement(node, env)
	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)
	case *ast.ReturnStatement:
		val := Eval(node.ReturnValue, env)
		if isError(val) {
			return val
		}
		return &object.ReturnValue{Value: val}
	case *ast.LetStatement:
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}
		env.Set(node.Name.Value, val)
		return nil
	// expressions
	case *ast.CallExpression:
		return evalCallExpression(node, env)
	case *ast.Identifier:
		return evalIdentifier(node, env)
	case *ast.FunctionLiteral:
		params := node.Parameters
		body := node.Body
		return &object.Function{Parameters: params, Env: env, Body: body}
	case *ast.ArrayLiteral:
		return evalArrayLiteral(node, env)
	case *ast.IndexExpression:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}
		index := Eval(node.Index, env)
		if isError(index) {
			return index
		}
		return evalIndexExpression(left, index)
	case *ast.PrefixExpression:
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalPrefixExpression(node.Operator, right)
	case *ast.InfixExpression:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalInfixExpression(node.Operator, left, right)
	case *ast.IfExpression:
		return evalIfExpression(node, env)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.StringLiteral:
		return &object.String{Value: node.Value}
	case *ast.Boolean:
		return referenceBooleanObject(node.Value)
	}
	return nil
}

func evalProgram(program *ast.Program, env *object.Environment) (result object.Object) {
	for _, statement := range program.Statements {
		result = Eval(statement, env)
		switch result := result.(type) {
		case *object.ReturnValue:
			return result.Value
		case *object.Error:
			return result
		}
	}
	return result
}

func evalBlockStatement(block *ast.BlockStatement, env *object.Environment) (result object.Object) {
	for _, statement := range block.Statements {
		result = Eval(statement, env)
		switch result := result.(type) {
		case *object.ReturnValue:
			return result
		case *object.Error:
			return result
		}
	}
	return result
}

func evalCallExpression(node *ast.CallExpression, env *object.Environment) object.Object {
	function := Eval(node.Function, env)
	if isError(function) {
		return function
	}
	var args []object.Object
	for _, e := range node.Arguments {
		arg := Eval(e, env)
		if isError(arg) {
			return arg
		}
		args = append(args, arg)
	}
	return evalFunction(function, args)
}

func evalFunction(fn object.Object, args []object.Object) object.Object {
	switch fn := fn.(type) {
	case *object.Function:
		extendedEnv := extendFunctionEnv(fn, args)
		evaluated := Eval(fn.Body, extendedEnv)
		switch evaluated := evaluated.(type) {
		case *object.ReturnValue:
			return evaluated.Value
		}
		return evaluated
	case *object.Builtin:
		return fn.Fn(args...)
	default:
		return newError("not a function: %s", fn.Type())
	}
}

func evalArrayLiteral(node *ast.ArrayLiteral, env *object.Environment) object.Object {
	var elements []object.Object
	for _, e := range node.Elements {
		element := Eval(e, env)
		if isError(element) {
			return element
		}
		elements = append(elements, element)
	}
	return &object.Array{Elements: elements}
}

func evalIndexExpression(left, index object.Object) object.Object {
	l, ok := left.(*object.Array)
	if !ok {
		return newError("index operator not supported: %s", left.Type())
	}
	i, ok := index.(*object.Integer)
	if !ok {
		return newError("index operator not supported: %s", left.Type())
	}
	return evalArrayIndexExpression(l, i)
}

func evalArrayIndexExpression(array *object.Array, index *object.Integer) object.Object {
	idx := index.Value
	max := int64(len(array.Elements) - 1)
	if idx < 0 || idx > max {
		return NULL
	}
	return array.Elements[idx]
}

func evalIdentifier(node *ast.Identifier, env *object.Environment) object.Object {
	if val, ok := env.Get(node.Value); ok {
		return val
	}
	if builtin, ok := builtins[node.Value]; ok {
		return builtin
	}
	return newError("identifier not found: " + node.Value)
}

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusPrefixOperatorExpression(right)
	default:
		return newErrorUnknownPrefixOperator(operator, right)
	}
}

func evalBangOperatorExpression(right object.Object) object.Object {
	switch right := right.(type) {
	case *object.Boolean:
		return referenceBooleanObject(!right.Value)
	case *object.Null:
		return referenceBooleanObject(!false)
	default:
		return referenceBooleanObject(!true)
	}
}

func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	switch right := right.(type) {
	case *object.Integer:
		return &object.Integer{Value: -right.Value}
	default:
		return newErrorUnknownPrefixOperator("-", right)
	}
}

func evalInfixExpression(
	operator string,
	left, right object.Object,
) object.Object {
	if li, ok := left.(*object.Integer); ok {
		if ri, ok := right.(*object.Integer); ok {
			return evalIntegerInfixExpression(operator, li, ri)
		}
	}
	if li, ok := left.(*object.String); ok {
		if ri, ok := right.(*object.String); ok {
			return evalStringInfixExpression(operator, li, ri)
		}
	}
	return evalObjectInfixExpression(operator, left, right)
}

func evalObjectInfixExpression(
	operator string,
	left, right object.Object,
) object.Object {
	switch operator {
	case "==":
		return referenceBooleanObject(left == right)
	case "!=":
		return referenceBooleanObject(left != right)
	default:
		return newErrorInfixExpression(operator, left, right)
	}
}

func evalIntegerInfixExpression(
	operator string,
	left, right *object.Integer,
) object.Object {
	switch operator {
	case "+":
		return &object.Integer{Value: left.Value + right.Value}
	case "-":
		return &object.Integer{Value: left.Value - right.Value}
	case "*":
		return &object.Integer{Value: left.Value * right.Value}
	case "/":
		return &object.Integer{Value: left.Value / right.Value}
	case "<":
		return referenceBooleanObject(left.Value < right.Value)
	case ">":
		return referenceBooleanObject(left.Value > right.Value)
	case "==":
		return referenceBooleanObject(left.Value == right.Value)
	case "!=":
		return referenceBooleanObject(left.Value != right.Value)
	default:
		return newErrorInfixExpression(operator, left, right)
	}
}

func evalStringInfixExpression(
	operator string,
	left, right *object.String,
) object.Object {
	switch operator {
	case "+":
		return &object.String{Value: left.Value + right.Value}
	case "==":
		return referenceBooleanObject(left.Value == right.Value)
	case "!=":
		return referenceBooleanObject(left.Value != right.Value)
	default:
		return newErrorInfixExpression(operator, left, right)
	}
}

func evalIfExpression(ie *ast.IfExpression, env *object.Environment) object.Object {
	condition := Eval(ie.Condition, env)
	if isError(condition) {
		return condition
	}
	if isTruthy(condition) {
		return Eval(ie.Consequence, env)
	} else if ie.Alternative != nil {
		return Eval(ie.Alternative, env)
	} else {
		return NULL
	}
}

func extendFunctionEnv(fn *object.Function, args []object.Object) *object.Environment {
	env := object.NewEnclosedEnvironment(fn.Env)
	for i, param := range fn.Parameters {
		env.Set(param.Value, args[i])
	}
	return env
}

func isTruthy(obj object.Object) bool {
	switch obj := obj.(type) {
	case *object.Null:
		return false
	case *object.Boolean:
		return obj.Value
	default:
		return true
	}
}

func referenceBooleanObject(value bool) *object.Boolean {
	if value {
		return TRUE
	}
	return FALSE
}
func isError(obj object.Object) bool {
	if _, ok := obj.(*object.Error); ok {
		return true
	}
	return false
}

func newErrorInfixExpression(operator string, left, right object.Object) *object.Error {
	if left.Type() != right.Type() {
		return newErrorTypeMismatch(operator, left, right)
	}
	return newErrorUnknownInfixOperator(operator, left, right)
}

func newErrorTypeMismatch(operator string, left, right object.Object) *object.Error {
	return newError("type mismatch: %s %s %s", left.Type(), operator, right.Type())
}

func newErrorUnknownInfixOperator(operator string, left, right object.Object) *object.Error {
	return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
}

func newErrorUnknownPrefixOperator(operator string, right object.Object) *object.Error {
	return newError("unknown operator: %s%s", operator, right.Type())
}

func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}
