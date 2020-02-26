package evaluator

import (
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
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.PrefixExpression:
		right := Eval(node.Right)
		return evalPrefixExpression(node.Operator, right)
	case *ast.InfixExpression:
		left := Eval(node.Left)
		right := Eval(node.Right)
		return evalInfixExpression(node.Operator, left, right)
	// expressions
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.Boolean:
		return referenceBooleanObject(node.Value)
	}
	return nil
}

func evalStatements(statements []ast.Statement) object.Object {
	var result object.Object
	for _, statement := range statements {
		result = Eval(statement)
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
		return NULL
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
		return NULL
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
		return NULL
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
		return NULL
	}
}

func referenceBooleanObject(value bool) *object.Boolean {
	if value {
		return TRUE
	}
	return FALSE
}
