package evaluator

import (
	"github.com/tshinag/monkey/ast"
	"github.com/tshinag/monkey/object"
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
