package fig

// Int Data /----

type cfInt map[string]int

func (f *cField) NewIntField(key string) *cfInt {
	newIntField := cfInt{}
	(*f)[key] = &newIntField
	return &newIntField

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

func (f *cField) NewStringField(key, val string) *cfString {
	newStringField := cfString{}
	(*f)[key] = &newStringField
	return &newStringField
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

// type ConfObj interface {
// 	*cField | *cfInt | *cfString | *int | *string
// }

type cField map[string]any

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
func (l *cField) Set(key string, newValue any) {
	(*l)[key] = newValue
}

func (c *cField) Pop(key string) any {
	tempval := (*c)[key]
	defer delete(*c, key)
	return tempval
}
