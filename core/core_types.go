package core

// Int Data /----

type cfInt map[string]int

func (h *Handeler) Int(key string, val int) *cfInt {
	h.IntData[key] = val
	return &(h.IntData)
}
func (c *cfInt) Set(key string, newValue int) {
	(*c)[key] = newValue
}

// String Data /----

type cfString map[string]string

func (h *Handeler) String(key, val string) *cfString {
	h.StringData[key] = val
	return &(h.StringData)
}
func (c *cfString) Set(key string, newValue string) {
	(*c)[key] = newValue
}

// List Data /----

type cflist map[string]any

func (h *Handeler) List(key string, val any) *cflist {
	h.ListData[key] = val
	return &(h.ListData)
}
func (l *cflist) Set(key string, newValue any) {
	(*l)[key] = newValue
}
