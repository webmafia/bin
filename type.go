package bin

import (
	"fmt"
	"io"
	"reflect"
	"unsafe"

	"github.com/webmafia/fast"
)

type Type interface {
	encodedSize(ptr unsafe.Pointer) int
	encode(ptr unsafe.Pointer, b fast.Writer)
	decode(ptr unsafe.Pointer, b *fast.BinaryBufferReader, nocopy bool) (err error)
	typeHash(io.Writer)
}

func getType(typ reflect.Type, offset uintptr, allowAllocations bool) (t Type, err error) {
	kind := typ.Kind()

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
		return getArrayType(typ, offset, allowAllocations)

	case reflect.Slice:
		return getSliceType(typ, offset, allowAllocations)

	case reflect.String:
		return stringType{offset: offset}, nil

	case reflect.Struct:
		return getStructType(typ, offset, allowAllocations)

	default:
		return nil, fmt.Errorf("unsupported type: %s", kind)
	}
}
