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

func (c *cfInt) Pop(key string) int {
	tempval := (*c)[key]
	defer delete(*c, key)
	return tempval
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

func (c *cfString) Pop(key string) string {
	tempval := (*c)[key]
	defer delete(*c, key)
	return tempval
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

func (c *cflist) Pop(key string) any {
	tempval := (*c)[key]
	defer delete(*c, key)
	return tempval
}
