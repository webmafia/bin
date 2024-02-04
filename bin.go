package bin

import (
	"reflect"
	"unsafe"

	"github.com/webbmaffian/go-fast"
)

func Encode[T any](types *Types, b *fast.BinaryBuffer, v *T) (err error) {
	c, err := types.GetCoder(reflect.TypeOf(*v))

	if err != nil {
		return
	}

	c.encode(unsafe.Pointer(v), b)

	return
}

func Decode[T any](types *Types, b *fast.BinaryBufferReader, v *T) (err error) {
	c, err := types.GetCoder(reflect.TypeOf(*v))

	if err != nil {
		return
	}

	c.decode(unsafe.Pointer(v), b)

	return
}
