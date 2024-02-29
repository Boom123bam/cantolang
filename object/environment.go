package object

type Environment struct {
	Parent   *Environment
	Bindings map[string]Object
}

func NewEnvironment(parent *Environment) *Environment {
	return &Environment{Parent: parent, Bindings: make(map[string]Object)}
}

func (e *Environment) Set(varName string, value Object) {
	e.Bindings[varName] = value
}

func (e *Environment) Get(varName string) (Object, bool) {
	val, ok := e.Bindings[varName]
	if !ok && e.Parent != nil {
		val, ok = e.Parent.Get(varName)
	}
	return val, ok
}
