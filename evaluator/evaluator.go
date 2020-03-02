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
func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	// statements
	case *ast.Program:
		return evalProgram(node)
	case *ast.BlockStatement:
		return evalBlockStatement(node)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.ReturnStatement:
		val := Eval(node.ReturnValue)
		if isError(val) {
			return val
		}
		return &object.ReturnValue{Value: val}
	// expressions
	case *ast.PrefixExpression:
		right := Eval(node.Right)
		if isError(right) {
			return right
		}
		return evalPrefixExpression(node.Operator, right)
	case *ast.InfixExpression:
		left := Eval(node.Left)
		if isError(left) {
			return left
		}
		right := Eval(node.Right)
		if isError(right) {
			return right
		}
		return evalInfixExpression(node.Operator, left, right)
	case *ast.IfExpression:
		return evalIfExpression(node)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.Boolean:
		return referenceBooleanObject(node.Value)
	}
	return nil
}

func evalProgram(program *ast.Program) object.Object {
	var result object.Object
	for _, statement := range program.Statements {
		result = Eval(statement)
		switch result := result.(type) {
		case *object.ReturnValue:
			return result.Value
		case *object.Error:
			return result
		}
	}
	return result
}

func evalBlockStatement(block *ast.BlockStatement) object.Object {
	var result object.Object
	for _, statement := range block.Statements {
		result = Eval(statement)
		switch result := result.(type) {
		case *object.ReturnValue:
			return result
		case *object.Error:
			return result
		}
	}
	return result
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

func evalIfExpression(ie *ast.IfExpression) object.Object {
	condition := Eval(ie.Condition)
	if isError(condition) {
		return condition
	}
	if isTruthy(condition) {
		return Eval(ie.Consequence)
	} else if ie.Alternative != nil {
		return Eval(ie.Alternative)
	} else {
		return NULL
	}
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
