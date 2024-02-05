package bin

import (
	"fmt"
	"reflect"
	"unsafe"

	"github.com/webbmaffian/go-fast"
)

type Type interface {
	encode(ptr unsafe.Pointer, b *fast.BinaryBuffer)
	decode(ptr unsafe.Pointer, b *fast.BinaryBufferReader) (err error)
}

func getType(typ reflect.Type, offset uintptr, hasher func(reflect.Kind)) (t Type, err error) {
	kind := typ.Kind()

	hasher(kind)

	switch kind {

	case reflect.Bool:
		return boolType{offset: offset}, nil

	case reflect.Int:
		return intType{offset: offset}, nil

	case reflect.Int8:
		return int8Type{offset: offset}, nil

	case reflect.Int16:
		return int16Type{offset: offset}, nil

	case reflect.Int32:
		return int32Type{offset: offset}, nil

	case reflect.Int64:
		return int64Type{offset: offset}, nil

	case reflect.Uint:
		return uintType{offset: offset}, nil

	case reflect.Uint8:
		return uint8Type{offset: offset}, nil

	case reflect.Uint16:
		return uint16Type{offset: offset}, nil

	case reflect.Uint32:
		return uint32Type{offset: offset}, nil

	case reflect.Uint64:
		return uint64Type{offset: offset}, nil

	case reflect.Float32:
		return float32Type{offset: offset}, nil

	case reflect.Float64:
		return float64Type{offset: offset}, nil

	case reflect.Array:
		return getArrayType(typ, offset, hasher)

	case reflect.Slice:
		return getSliceType(typ, offset, hasher)

	case reflect.String:
		return stringType{offset: offset}, nil

	case reflect.Struct:
		return getStructType(typ, offset, hasher)

	default:
		return nil, fmt.Errorf("unsupported type: %s", kind)
	}
}
