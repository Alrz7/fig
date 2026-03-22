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

// List Data /----

// type ConfObj interface {
// 	*cField | *cfInt | *cfString | *int | *string
// }

type Topic map[string]any

// func (h *Handeler) NewField(key string) *cField {   // this feature is not supported for Reading YET!.
//
//		var newField = cField{}
//		if elmnt, ok := (*h).Data[key]; ok {
//			newField = cField{key: elmnt}
//		} else {
//			(*h).Data[key] = &newField
//		}
//		return &newField
//	}
func (l *Topic) Set(key string, newValue any) {
	(*l)[key] = newValue
}

func (c *Topic) Pop(key string) any {
	tempval := (*c)[key]
	defer delete(*c, key)
	return tempval
}
