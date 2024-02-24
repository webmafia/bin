package bin

import (
	"reflect"
	"unsafe"

	"github.com/webbmaffian/go-fast"
)

type structType struct {
	fields []Type
}

func getStructType(typ reflect.Type, offset uintptr, hasher func(reflect.Kind)) (Type, error) {
	num := typ.NumField()
	t := structType{
		fields: make([]Type, 0, num),
	}

	for i := 0; i < num; i++ {
		f := typ.Field(i)

		if !f.IsExported() {
			continue
		}

		subtyp, err := getType(f.Type, offset+f.Offset, hasher)

		if err != nil {
			return nil, err
		}

		t.fields = append(t.fields, subtyp)
	}

	return t, nil
}

func (c structType) encodedSize(ptr unsafe.Pointer) (s int) {
	for i := range c.fields {
		s += c.fields[i].encodedSize(ptr)
	}

	return
}

func (c structType) encode(ptr unsafe.Pointer, b *fast.BinaryBuffer) {
	for i := range c.fields {
		c.fields[i].encode(ptr, b)
	}
}

func (c structType) decode(ptr unsafe.Pointer, b *fast.BinaryBufferReader) (err error) {
	for i := range c.fields {
		if err = c.fields[i].decode(ptr, b); err != nil {
			return
		}
	}

	return
}
