package bin

import (
	"unsafe"

	"github.com/webbmaffian/go-fast"
)

type boolField struct {
	offset uintptr
}

func (f boolField) encode(ptr unsafe.Pointer, b *fast.BinaryBuffer) {
	b.WriteBool(*(*bool)(unsafe.Add(ptr, f.offset)))
}

func (f boolField) decode(ptr unsafe.Pointer, b *fast.BinaryBufferReader) {
	*(*bool)(unsafe.Add(ptr, f.offset)) = b.ReadBool()
}

type intField struct {
	offset uintptr
}

func (f intField) encode(ptr unsafe.Pointer, b *fast.BinaryBuffer) {
	b.WriteInt(*(*int)(unsafe.Add(ptr, f.offset)))
}

func (f intField) decode(ptr unsafe.Pointer, b *fast.BinaryBufferReader) {
	*(*int)(unsafe.Add(ptr, f.offset)) = b.ReadInt()
}

type int8Field struct {
	offset uintptr
}

func (f int8Field) encode(ptr unsafe.Pointer, b *fast.BinaryBuffer) {
	b.WriteInt8(*(*int8)(unsafe.Add(ptr, f.offset)))
}

func (f int8Field) decode(ptr unsafe.Pointer, b *fast.BinaryBufferReader) {
	*(*int8)(unsafe.Add(ptr, f.offset)) = b.ReadInt8()
}

type int16Field struct {
	offset uintptr
}

func (f int16Field) encode(ptr unsafe.Pointer, b *fast.BinaryBuffer) {
	b.WriteInt16(*(*int16)(unsafe.Add(ptr, f.offset)))
}

func (f int16Field) decode(ptr unsafe.Pointer, b *fast.BinaryBufferReader) {
	*(*int16)(unsafe.Add(ptr, f.offset)) = b.ReadInt16()
}

type int32Field struct {
	offset uintptr
}

func (f int32Field) encode(ptr unsafe.Pointer, b *fast.BinaryBuffer) {
	b.WriteInt32(*(*int32)(unsafe.Add(ptr, f.offset)))
}

func (f int32Field) decode(ptr unsafe.Pointer, b *fast.BinaryBufferReader) {
	*(*int32)(unsafe.Add(ptr, f.offset)) = b.ReadInt32()
}

type int64Field struct {
	offset uintptr
}

func (f int64Field) encode(ptr unsafe.Pointer, b *fast.BinaryBuffer) {
	b.WriteInt64(*(*int64)(unsafe.Add(ptr, f.offset)))
}

func (f int64Field) decode(ptr unsafe.Pointer, b *fast.BinaryBufferReader) {
	*(*int64)(unsafe.Add(ptr, f.offset)) = b.ReadInt64()
}

type uintField struct {
	offset uintptr
}

func (f uintField) encode(ptr unsafe.Pointer, b *fast.BinaryBuffer) {
	b.WriteUint(*(*uint)(unsafe.Add(ptr, f.offset)))
}

func (f uintField) decode(ptr unsafe.Pointer, b *fast.BinaryBufferReader) {
	*(*uint)(unsafe.Add(ptr, f.offset)) = b.ReadUint()
}

type uint8Field struct {
	offset uintptr
}

func (f uint8Field) encode(ptr unsafe.Pointer, b *fast.BinaryBuffer) {
	b.WriteUint8(*(*uint8)(unsafe.Add(ptr, f.offset)))
}

func (f uint8Field) decode(ptr unsafe.Pointer, b *fast.BinaryBufferReader) {
	*(*uint8)(unsafe.Add(ptr, f.offset)) = b.ReadUint8()
}

type uint16Field struct {
	offset uintptr
}

func (f uint16Field) encode(ptr unsafe.Pointer, b *fast.BinaryBuffer) {
	b.WriteUint16(*(*uint16)(unsafe.Add(ptr, f.offset)))
}

func (f uint16Field) decode(ptr unsafe.Pointer, b *fast.BinaryBufferReader) {
	*(*uint16)(unsafe.Add(ptr, f.offset)) = b.ReadUint16()
}

type uint32Field struct {
	offset uintptr
}

func (f uint32Field) encode(ptr unsafe.Pointer, b *fast.BinaryBuffer) {
	b.WriteUint32(*(*uint32)(unsafe.Add(ptr, f.offset)))
}

func (f uint32Field) decode(ptr unsafe.Pointer, b *fast.BinaryBufferReader) {
	*(*uint32)(unsafe.Add(ptr, f.offset)) = b.ReadUint32()
}

type uint64Field struct {
	offset uintptr
}

func (f uint64Field) encode(ptr unsafe.Pointer, b *fast.BinaryBuffer) {
	b.WriteUint64(*(*uint64)(unsafe.Add(ptr, f.offset)))
}

func (f uint64Field) decode(ptr unsafe.Pointer, b *fast.BinaryBufferReader) {
	*(*uint64)(unsafe.Add(ptr, f.offset)) = b.ReadUint64()
}

type float32Field struct {
	offset uintptr
}

func (f float32Field) encode(ptr unsafe.Pointer, b *fast.BinaryBuffer) {
	b.WriteFloat32(*(*float32)(unsafe.Add(ptr, f.offset)))
}

func (f float32Field) decode(ptr unsafe.Pointer, b *fast.BinaryBufferReader) {
	*(*float32)(unsafe.Add(ptr, f.offset)) = b.ReadFloat32()
}

type float64Field struct {
	offset uintptr
}

func (f float64Field) encode(ptr unsafe.Pointer, b *fast.BinaryBuffer) {
	b.WriteFloat64(*(*float64)(unsafe.Add(ptr, f.offset)))
}

func (f float64Field) decode(ptr unsafe.Pointer, b *fast.BinaryBufferReader) {
	*(*float64)(unsafe.Add(ptr, f.offset)) = b.ReadFloat64()
}
