package bin

import (
	"fmt"
	"io"
	"reflect"
	"unsafe"

	"github.com/webmafia/fast/binary"
)

type Encoder func(w *binary.StreamWriter, ptr unsafe.Pointer) error

func NewEncoder[T any](w io.Writer, v *T, options ...Options) (enc func(v *T) error, err error) {
	var opt Options

	if len(options) > 0 {
		opt = options[0]
	}

	rawEnc, err := getEncoder(reflect.TypeOf(*v), &opt)

	if err != nil {
		return
	}

	b := binary.NewStreamWriter(w)

	return func(v *T) error {
		return rawEnc(b, unsafe.Pointer(v))
	}, nil
}

func getEncoder(typ reflect.Type, opt *Options) (enc Encoder, err error) {
	kind := typ.Kind()

	switch kind {

	case reflect.Bool:
		return boolEncoder, nil

	case reflect.Int:
		return intEncoder, nil

	case reflect.Int8:
		return int8Encoder, nil

	case reflect.Int16:
		return int16Encoder, nil

	case reflect.Int32:
		return int32Encoder, nil

	case reflect.Int64:
		return int64Encoder, nil

	case reflect.Uint:
		return uintEncoder, nil

	case reflect.Uint8:
		return uint8Encoder, nil

	case reflect.Uint16:
		return uint16Encoder, nil

	case reflect.Uint32:
		return uint32Encoder, nil

	case reflect.Uint64:
		return uint64Encoder, nil

	case reflect.Float32:
		return float32Encoder, nil

	case reflect.Float64:
		return float64Encoder, nil

	case reflect.Array:
		return getArrayEncoder(typ, opt)

	case reflect.Slice:
		return getSliceEncoder(typ, opt)

	case reflect.String:
		return stringEncoder, nil

	case reflect.Struct:
		return getStructEncoder(typ, opt)

	default:
		return nil, fmt.Errorf("unsupported type: %s", kind)
	}
}
