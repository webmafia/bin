package bin

import (
	"reflect"
	"unsafe"

	"github.com/webbmaffian/go-fast"
)

type field struct {
	kind   reflect.Kind
	offset uintptr
}

func (f field) encode(ptr unsafe.Pointer, b *fast.BinaryBuffer) {
	ptr = unsafe.Add(ptr, f.offset)

	switch f.kind {

	case reflect.Bool:
		b.WriteBool(*(*bool)(ptr))

	case reflect.Int:
		b.WriteInt(*(*int)(ptr))

	case reflect.Int8:
		b.WriteInt8(*(*int8)(ptr))

	case reflect.Int16:
		b.WriteInt16(*(*int16)(ptr))

	case reflect.Int32:
		b.WriteInt32(*(*int32)(ptr))

	case reflect.Int64:
		b.WriteInt64(*(*int64)(ptr))

	case reflect.Uint:
		b.WriteUint(*(*uint)(ptr))

	case reflect.Uint8:
		b.WriteUint8(*(*uint8)(ptr))

	case reflect.Uint16:
		b.WriteUint16(*(*uint16)(ptr))

	case reflect.Uint32:
		b.WriteUint32(*(*uint32)(ptr))

	case reflect.Uint64:
		b.WriteUint64(*(*uint64)(ptr))

	case reflect.Float32:
		b.WriteFloat32(*(*float32)(ptr))

	case reflect.Float64:
		b.WriteFloat64(*(*float64)(ptr))

	case reflect.Array:
		panic("not implemented yet")

	case reflect.Map:
		panic("not implemented yet")

	case reflect.Slice:
		panic("not implemented yet")

	case reflect.String:
		panic("not implemented yet")

	}
}
