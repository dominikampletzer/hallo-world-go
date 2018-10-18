package registry

type Registry struct {
	services map[string]interface{}
}

func (reg *Registry) Get(str string) interface{} {
	result := reg.services[str]
	return result
}

func (reg *Registry) Add(str string, service interface{}) {
	reg.services[str] = service
}
func NewRegistry() *Registry {
	registry := new(Registry)
	registry.services = make(map[string]interface{})
	return registry
}
