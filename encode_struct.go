package bin

import (
	"reflect"
	"unsafe"

	"github.com/webmafia/fast/binary"
)

func getStructEncoder(typ reflect.Type, opt *Options) (Encoder, error) {
	type field struct {
		enc    Encoder
		offset uintptr
	}

	numFields := typ.NumField()
	fields := make([]field, 0, numFields)

	for i := 0; i < numFields; i++ {
		f := typ.Field(i)

		if !opt.KeepUnexportedFields && !f.IsExported() {
			continue
		}

		enc, err := getEncoder(f.Type, opt)

		if err != nil {
			return nil, err
		}

		fields = append(fields, field{
			enc:    enc,
			offset: f.Offset,
		})
	}

	return func(w *binary.StreamWriter, ptr unsafe.Pointer) (err error) {
		for i := range fields {
			if err = fields[i].enc(w, unsafe.Add(ptr, fields[i].offset)); err != nil {
				return
			}
		}

		return
	}, nil
}
