package evaluator

import (
	"bytes"
	"fmt"
	"os"

	"github.com/baybaraandrey/elephant/object"
)

func init() {
	registerBuiltin("len", blen)
	registerBuiltin("first", bfirst)
	registerBuiltin("print", bprint)
	registerBuiltin("exit", bexit)
	registerBuiltin("locals", blocals)
	registerBuiltin("type", btype)
}

func registerBuiltin(
	name string,
	f func(env *object.Environment, args ...object.Object) object.Object,
) {
	builtins[name] = &object.Builtin{Fn: f, Env: nil}
}

var builtins = map[string]*object.Builtin{}

func blen(env *object.Environment, args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1",
			len(args))
	}

	switch arg := args[0].(type) {
	case *object.String:
		return &object.Integer{Value: int64(len(arg.Value))}
	case *object.Array:
		return &object.Integer{Value: int64(len(arg.Elements))}
	default:
		return newError("argument to `len` not supported, got %s",
			args[0].Type())
	}

}

func btype(env *object.Environment, args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1",
			len(args))
	}
	return &object.ObjType{Value: string(args[0].Type())}
}

func blocals(env *object.Environment, args ...object.Object) object.Object {
	if env != nil {
		return env.ToHash()
	}
	return NULL
}

func bfirst(env *object.Environment, args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1",
			len(args))
	}

	if args[0].Type() != object.ARRAY_OBJ {
		return newError("argument to 'first' must be ARRAY, got %s",
			args[0].Type())
	}

	arr := args[0].(*object.Array)
	if len(arr.Elements) > 0 {
		return arr.Elements[0]
	}

	return NULL
}

func bprint(env *object.Environment, args ...object.Object) object.Object {
	var out bytes.Buffer
	for index, arg := range args {
		out.WriteString(arg.Inspect())
		if index+1 == len(args) {
			continue
		}
		out.WriteString(" ")
	}
	out.WriteString("\n")
	fmt.Print(out.String())
	return NULL
}

func bexit(env *object.Environment, args ...object.Object) object.Object {
	switch len(args) {
	case 0:
		os.Exit(0)
	case 1:
		if args[0].Type() != object.INTEGER_OBJ {
			return newError("argument to 'exit' must be INT, got %s",
				args[0].Type())

		}
		status := args[0].(*object.Integer).Value
		os.Exit(int(status))
	default:
		return newError("wrong number of arguments. got=%d, want=1",
			len(args))
	}
	return NULL
}
