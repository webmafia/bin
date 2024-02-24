package bin

import (
	"io"
	"reflect"
)

func (t arrayType) typeHash(w io.Writer) {
	w.Write([]byte{
		byte(reflect.Array),
		byte(t.len),
		byte(t.len >> 8),
		byte(t.len >> 16),
		byte(t.len >> 24),
		byte(t.len >> 32),
		byte(t.len >> 40),
		byte(t.len >> 48),
		byte(t.len >> 56),
	})
	t.typ.typeHash(w)
}

func (boolType) typeHash(w io.Writer) {
	w.Write([]byte{byte(reflect.Bool)})
}

func (float32Type) typeHash(w io.Writer) {
	w.Write([]byte{byte(reflect.Float32)})
}

func (float64Type) typeHash(w io.Writer) {
	w.Write([]byte{byte(reflect.Float64)})
}

func (intType) typeHash(w io.Writer) {
	w.Write([]byte{byte(reflect.Int)})
}

func (int8Type) typeHash(w io.Writer) {
	w.Write([]byte{byte(reflect.Int8)})
}

func (int16Type) typeHash(w io.Writer) {
	w.Write([]byte{byte(reflect.Int16)})
}

func (int32Type) typeHash(w io.Writer) {
	w.Write([]byte{byte(reflect.Int32)})
}

func (int64Type) typeHash(w io.Writer) {
	w.Write([]byte{byte(reflect.Int64)})
}

func (uintType) typeHash(w io.Writer) {
	w.Write([]byte{byte(reflect.Uint)})
}

func (uint8Type) typeHash(w io.Writer) {
	w.Write([]byte{byte(reflect.Uint8)})
}

func (uint16Type) typeHash(w io.Writer) {
	w.Write([]byte{byte(reflect.Uint16)})
}

func (uint32Type) typeHash(w io.Writer) {
	w.Write([]byte{byte(reflect.Uint32)})
}

func (uint64Type) typeHash(w io.Writer) {
	w.Write([]byte{byte(reflect.Uint64)})
}

func (t sliceType) typeHash(w io.Writer) {
	w.Write([]byte{byte(reflect.Slice)})
	t.typ.typeHash(w)
}

func (stringType) typeHash(w io.Writer) {
	w.Write([]byte{byte(reflect.String)})
}

func (t structType) typeHash(w io.Writer) {
	for i := range t.fields {
		t.fields[i].typeHash(w)
	}
}

func (r raw) typeHash(w io.Writer) {
	w.Write([]byte{
		byte(reflect.Struct),
		byte(r.size),
		byte(r.size >> 8),
		byte(r.size >> 16),
		byte(r.size >> 24),
		byte(r.size >> 32),
		byte(r.size >> 40),
		byte(r.size >> 48),
		byte(r.size >> 56),
	})
}
