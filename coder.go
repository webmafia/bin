package bin

import (
	"reflect"
	"unsafe"

	"github.com/webbmaffian/go-fast"
)

type coder struct {
	fields []field
}

func (c *coder) addField(kind reflect.Kind, offset uintptr) {
	switch kind {

	case reflect.Bool:
		c.fields = append(c.fields, boolField{offset: offset})

	case reflect.Int:
		c.fields = append(c.fields, intField{offset: offset})

	case reflect.Int8:
		c.fields = append(c.fields, int8Field{offset: offset})

	case reflect.Int16:
		c.fields = append(c.fields, int16Field{offset: offset})

	case reflect.Int32:
		c.fields = append(c.fields, int32Field{offset: offset})

	case reflect.Int64:
		c.fields = append(c.fields, int64Field{offset: offset})

	case reflect.Uint:
		c.fields = append(c.fields, uintField{offset: offset})

	case reflect.Uint8:
		c.fields = append(c.fields, uint8Field{offset: offset})

	case reflect.Uint16:
		c.fields = append(c.fields, uint16Field{offset: offset})

	case reflect.Uint32:
		c.fields = append(c.fields, uint32Field{offset: offset})

	case reflect.Uint64:
		c.fields = append(c.fields, uint64Field{offset: offset})

	case reflect.Float32:
		c.fields = append(c.fields, float32Field{offset: offset})

	case reflect.Float64:
		c.fields = append(c.fields, float64Field{offset: offset})

	case reflect.Array:
		panic("not implemented yet")

	case reflect.Map:
		panic("not implemented yet")

	case reflect.Slice:
		panic("not implemented yet")

	case reflect.String:
		c.fields = append(c.fields, stringField{offset: offset})

	}
}

func (c coder) encode(ptr unsafe.Pointer, b *fast.BinaryBuffer) {
	for i := range c.fields {
		c.fields[i].encode(ptr, b)
	}
}

func (c coder) decode(ptr unsafe.Pointer, b *fast.BinaryBufferReader) {
	for i := range c.fields {
		c.fields[i].decode(ptr, b)
	}
}
