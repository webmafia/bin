package bin

import (
	"fmt"
	"io"
	"reflect"
	"unsafe"

	"github.com/webmafia/fast/binary"
)

type Decoder func(r *binary.StreamReader, ptr unsafe.Pointer) error

func NewDecoder[T any](r io.Reader, v *T, options ...Options) (dec func(v *T) error, err error) {
	var opt Options

	if len(options) > 0 {
		opt = options[0]
	}

	rawEnc, err := getDecoder(reflect.TypeOf(*v), &opt)

	if err != nil {
		return
	}

	b := binary.NewStreamReader(r)

	return func(v *T) error {
		return rawEnc(b, unsafe.Pointer(v))
	}, nil
}

func getDecoder(typ reflect.Type, opt *Options) (dec Decoder, err error) {
	kind := typ.Kind()

	switch kind {

	case reflect.Bool:
		return boolDecoder, nil

	case reflect.Int:
		return intDecoder, nil

	case reflect.Int8:
		return int8Decoder, nil

	case reflect.Int16:
		return int16Decoder, nil

	case reflect.Int32:
		return int32Decoder, nil

	case reflect.Int64:
		return int64Decoder, nil

	case reflect.Uint:
		return uintDecoder, nil

	case reflect.Uint8:
		return uint8Decoder, nil

	case reflect.Uint16:
		return uint16Decoder, nil

	case reflect.Uint32:
		return uint32Decoder, nil

	case reflect.Uint64:
		return uint64Decoder, nil

	case reflect.Float32:
		return float32Decoder, nil

	case reflect.Float64:
		return float64Decoder, nil

	case reflect.Array:
		return getArrayDecoder(typ, opt)

	case reflect.Slice:
		return getSliceDecoder(typ, opt)

	case reflect.String:
		return stringDecoder, nil

	case reflect.Struct:
		return getStructDecoder(typ, opt)

	default:
		return nil, fmt.Errorf("unsupported type: %s", kind)
	}
}
