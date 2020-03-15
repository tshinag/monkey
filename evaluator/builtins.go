package evaluator

import "github.com/tshinag/monkey/object"

var builtins = map[string]*object.Builtin{
	"len":   &object.Builtin{Fn: fnLen},
	"first": &object.Builtin{Fn: fnFirst},
}

func fnLen(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}
	switch arg := args[0].(type) {
	case *object.String:
		return &object.Integer{Value: int64(len(arg.Value))}
	case *object.Array:
		return &object.Integer{Value: int64(len(arg.Elements))}
	default:
		return newError("argument to `len` not supported, got %s", args[0].Type())
	}
}

func fnFirst(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}
	switch arg := args[0].(type) {
	case *object.Array:
		if len(arg.Elements) > 0 {
			return arg.Elements[0]
		}
		return NULL
	default:
		return newError("argument to `first` must be ARRAY, got %s", args[0].Type())
	}
}
