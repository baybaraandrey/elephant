package object

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

func (e *Environment) ToHash() *Hash {
	pairs := map[HashKey]HashPair{}
	for k, v := range e.store {
		hashKey := &String{Value: k}

		hashed := hashKey.HashKey()
		pairs[hashed] = HashPair{Key: hashKey, Value: v}
	}
	hash := &Hash{Pairs: pairs}

	return hash
}

func (e *Environment) Type() ObjectType { return HASH_OBJ }
func (e *Environment) Inspect() string {
	return e.ToHash().Inspect()
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
