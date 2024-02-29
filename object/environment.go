package object

var ENV = New()

type Environment struct {
	Bindings map[string]Object
}

func New() *Environment {
	return &Environment{Bindings: make(map[string]Object)}
}

func (e *Environment) Set(varName string, value Object) {
	e.Bindings[varName] = value
}

func (e *Environment) Get(varName string) (Object, bool) {
	val, ok := e.Bindings[varName]
	return val, ok
}
