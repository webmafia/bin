package bin

import (
	"reflect"
	"unsafe"

	"github.com/webmafia/fast/binary"
)

func getStructDecoder(typ reflect.Type, opt *Options) (Decoder, error) {
	type field struct {
		dec    Decoder
		offset uintptr
	}

	numFields := typ.NumField()
	fields := make([]field, 0, numFields)

	for i := 0; i < numFields; i++ {
		f := typ.Field(i)

		if !opt.KeepUnexportedFields && !f.IsExported() {
			continue
		}

		dec, err := getDecoder(f.Type, opt)

		if err != nil {
			return nil, err
		}

		fields = append(fields, field{
			dec:    dec,
			offset: f.Offset,
		})
	}

	return func(r *binary.StreamReader, ptr unsafe.Pointer) (err error) {
		for i := range fields {
			if err = fields[i].dec(r, unsafe.Add(ptr, fields[i].offset)); err != nil {
				return
			}
		}

		return
	}, nil
}
