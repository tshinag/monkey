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
	// expressions
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.Boolean:
		if node.Value {
			return TRUE
		}
		return FALSE
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
