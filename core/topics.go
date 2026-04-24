package fig

// Int Data /----

type tpInt map[string]int

func (f *Topic) NewIntField(key string) *tpInt {
	newIntField := tpInt{}
	(*f)[key] = &newIntField
	return &newIntField

}
func (c *tpInt) Set(key string, newValue int) {
	(*c)[key] = newValue
}

func (c *tpInt) Pop(key string) int {
	tempval := (*c)[key]
	defer delete(*c, key)
	return tempval
}

// String Data /----

type tpString map[string]string

func (f *Topic) NewStringField(key, val string) *tpString {
	newStringField := tpString{}
	(*f)[key] = &newStringField
	return &newStringField
}
func (c *tpString) Set(key string, newValue string) {
	(*c)[key] = newValue
}

func (c *tpString) Pop(key string) string {
	tempval := (*c)[key]
	defer delete(*c, key)
	return tempval
}

// free Data /----

type Topic map[string]any

func (t *Topic) Set(key string, newValue any) {
	(*t)[key] = newValue
}

func (t *Topic) Pop(key string) any {
	tempval := (*t)[key]
	defer delete(*t, key)
	return tempval
}
