package evaluator

import "github.com/tshinag/monkey/object"

var builtins = map[string]*object.Builtin{
	"len":   &object.Builtin{Fn: fnLen},
	"first": &object.Builtin{Fn: fnFirst},
	"last":  &object.Builtin{Fn: fnLast},
	"rest":  &object.Builtin{Fn: fnRest},
	"push":  &object.Builtin{Fn: fnPush},
}

func fnLen(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}
	switch arg := args[0].(type) {
	case *object.String:
		return fnLenString(arg)
	case *object.Array:
		return fnLenArray(arg)
	default:
		return newError("argument to `len` not supported, got %s", arg.Type())
	}
}

func fnLenString(str *object.String) object.Object {
	return &object.Integer{Value: int64(len(str.Value))}
}

func fnLenArray(arr *object.Array) object.Object {
	return &object.Integer{Value: int64(len(arr.Elements))}
}

func fnFirst(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}
	switch arg := args[0].(type) {
	case *object.Array:
		return fnFirstArray(arg)
	default:
		return newError("argument to `first` must be ARRAY, got %s", args[0].Type())
	}
}

func fnFirstArray(arr *object.Array) object.Object {
	if len(arr.Elements) > 0 {
		return arr.Elements[0]
	}
	return NULL
}

func fnLast(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}
	switch arg := args[0].(type) {
	case *object.Array:
		return fnLastArray(arg)
	default:
		return newError("argument to `last` must be ARRAY, got %s", args[0].Type())
	}
}

func fnLastArray(arr *object.Array) object.Object {
	length := len(arr.Elements)
	if length > 0 {
		return arr.Elements[length-1]
	}
	return NULL
}

func fnRest(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}
	switch arg := args[0].(type) {
	case *object.Array:
		return fnRestArray(arg)
	default:
		return newError("argument to `rest` must be ARRAY, got %s", args[0].Type())
	}
}

func fnRestArray(arr *object.Array) object.Object {
	length := len(arr.Elements)
	if length > 0 {
		newElements := make([]object.Object, length-1, length-1)
		copy(newElements, arr.Elements[1:length])
		return &object.Array{Elements: newElements}
	}
	return NULL
}

func fnPush(args ...object.Object) object.Object {
	if len(args) != 2 {
		return newError("wrong number of arguments. got=%d, want=2", len(args))
	}
	switch arg := args[0].(type) {
	case *object.Array:
		return fnPushArray(arg, args[1])
	default:
		return newError("argument to `push` must be ARRAY, got %s", args[0].Type())
	}
}

func fnPushArray(arr *object.Array, obj object.Object) object.Object {
	length := len(arr.Elements)
	newElements := make([]object.Object, length+1, length+1)
	copy(newElements, arr.Elements)
	newElements[length] = obj
	return &object.Array{Elements: newElements}
}
