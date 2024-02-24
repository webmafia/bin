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

func Encode2[T any](types *AtomicTypes, b *fast.BinaryBuffer, v *T) (err error) {
	c, err := types.GetCoder(v)

	if err != nil {
		return
	}

	c.encode(unsafe.Pointer(v), b)

	return
}

func EncodedSize[T any](types *Types, v *T) (s int, err error) {
	c, err := types.GetCoder(reflect.TypeOf(*v))

	if err != nil {
		return
	}

	s = c.EncodedSize(unsafe.Pointer(v))

	return
}

func Decode[T any](types *Types, b *fast.BinaryBufferReader, v *T) (err error) {
	c, err := types.GetCoder(reflect.TypeOf(*v))

	if err != nil {
		return
	}

	return c.decode(unsafe.Pointer(v), b)
}
