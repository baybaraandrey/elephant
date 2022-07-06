package object

import (
	"bytes"
	"fmt"
)

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer

	return env
}

type Environment struct {
	store map[string]Object
	outer *Environment
}

func (e *Environment) String() string {
	var out bytes.Buffer
	out.WriteString("{\n")
	for k, v := range e.store {
		out.WriteString(fmt.Sprintf("'%s': '%s',\n", k, v.Inspect()))
	}
	out.WriteString("}\n")
	return out.String()
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
