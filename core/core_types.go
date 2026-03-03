package core

// Int Data /----

type cfInt map[string]int

func (h *handeler) Int(key string, val int) *cfInt {
	h.intData[key] = val
	return &(h.intData)
}
func (c *cfInt) Set(key string, newValue int) {
	(*c)[key] = newValue
}

// String Data /----

type cfString map[string]string

func (h *handeler) String(key, val string) *cfString {
	h.stringData[key] = val
	return &(h.stringData)
}
func (c *cfString) Set(key string, newValue string) {
	(*c)[key] = newValue
}

// List Data /----

type cflist map[string]any

func (h *handeler) List(key string, val any) *cflist {
	h.listData[key] = val
	return &(h.listData)
}
func (l *cflist) Set(key string, newValue any) {
	(*l)[key] = newValue
}
