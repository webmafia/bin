package bin

import "reflect"

type coder struct {
	fields []field
}

func (c *coder) addField(kind reflect.Kind, offset uintptr) {
	c.fields = append(c.fields, field{
		kind:   kind,
		offset: offset,
	})
}
