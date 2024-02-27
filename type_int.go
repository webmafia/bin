package bin

import (
	"unsafe"

	"github.com/webbmaffian/go-fast"
)

type intType struct {
	offset uintptr
}

func (intType) encodedSize(_ unsafe.Pointer) int {
	return 8
}

func (f intType) encode(ptr unsafe.Pointer, b *fast.BinaryBuffer) {
	b.WriteInt(*(*int)(unsafe.Add(ptr, f.offset)))
}

func (f intType) decode(ptr unsafe.Pointer, b *fast.BinaryBufferReader, _ bool) error {
	*(*int)(unsafe.Add(ptr, f.offset)) = b.ReadInt()
	return nil
}

type int8Type struct {
	offset uintptr
}

func (int8Type) encodedSize(_ unsafe.Pointer) int {
	return 1
}

func (f int8Type) encode(ptr unsafe.Pointer, b *fast.BinaryBuffer) {
	b.WriteInt8(*(*int8)(unsafe.Add(ptr, f.offset)))
}

func (f int8Type) decode(ptr unsafe.Pointer, b *fast.BinaryBufferReader, _ bool) error {
	*(*int8)(unsafe.Add(ptr, f.offset)) = b.ReadInt8()
	return nil
}

type int16Type struct {
	offset uintptr
}

func (int16Type) encodedSize(_ unsafe.Pointer) int {
	return 2
}

func (f int16Type) encode(ptr unsafe.Pointer, b *fast.BinaryBuffer) {
	b.WriteInt16(*(*int16)(unsafe.Add(ptr, f.offset)))
}

func (f int16Type) decode(ptr unsafe.Pointer, b *fast.BinaryBufferReader, _ bool) error {
	*(*int16)(unsafe.Add(ptr, f.offset)) = b.ReadInt16()
	return nil
}

type int32Type struct {
	offset uintptr
}

func (int32Type) encodedSize(_ unsafe.Pointer) int {
	return 4
}

func (f int32Type) encode(ptr unsafe.Pointer, b *fast.BinaryBuffer) {
	b.WriteInt32(*(*int32)(unsafe.Add(ptr, f.offset)))
}

func (f int32Type) decode(ptr unsafe.Pointer, b *fast.BinaryBufferReader, _ bool) error {
	*(*int32)(unsafe.Add(ptr, f.offset)) = b.ReadInt32()
	return nil
}

type int64Type struct {
	offset uintptr
}

func (int64Type) encodedSize(_ unsafe.Pointer) int {
	return 8
}

func (f int64Type) encode(ptr unsafe.Pointer, b *fast.BinaryBuffer) {
	b.WriteInt64(*(*int64)(unsafe.Add(ptr, f.offset)))
}

func (f int64Type) decode(ptr unsafe.Pointer, b *fast.BinaryBufferReader, _ bool) error {
	*(*int64)(unsafe.Add(ptr, f.offset)) = b.ReadInt64()
	return nil
}

type uintType struct {
	offset uintptr
}

func (uintType) encodedSize(_ unsafe.Pointer) int {
	return 8
}

func (f uintType) encode(ptr unsafe.Pointer, b *fast.BinaryBuffer) {
	b.WriteUint(*(*uint)(unsafe.Add(ptr, f.offset)))
}

func (f uintType) decode(ptr unsafe.Pointer, b *fast.BinaryBufferReader, _ bool) error {
	*(*uint)(unsafe.Add(ptr, f.offset)) = b.ReadUint()
	return nil
}

type uint8Type struct {
	offset uintptr
}

func (uint8Type) encodedSize(_ unsafe.Pointer) int {
	return 1
}

func (f uint8Type) encode(ptr unsafe.Pointer, b *fast.BinaryBuffer) {
	b.WriteUint8(*(*uint8)(unsafe.Add(ptr, f.offset)))
}

func (f uint8Type) decode(ptr unsafe.Pointer, b *fast.BinaryBufferReader, _ bool) error {
	*(*uint8)(unsafe.Add(ptr, f.offset)) = b.ReadUint8()
	return nil
}

type uint16Type struct {
	offset uintptr
}

func (uint16Type) encodedSize(_ unsafe.Pointer) int {
	return 2
}

func (f uint16Type) encode(ptr unsafe.Pointer, b *fast.BinaryBuffer) {
	b.WriteUint16(*(*uint16)(unsafe.Add(ptr, f.offset)))
}

func (f uint16Type) decode(ptr unsafe.Pointer, b *fast.BinaryBufferReader, _ bool) error {
	*(*uint16)(unsafe.Add(ptr, f.offset)) = b.ReadUint16()
	return nil
}

type uint32Type struct {
	offset uintptr
}

func (uint32Type) encodedSize(_ unsafe.Pointer) int {
	return 4
}

func (f uint32Type) encode(ptr unsafe.Pointer, b *fast.BinaryBuffer) {
	b.WriteUint32(*(*uint32)(unsafe.Add(ptr, f.offset)))
}

func (f uint32Type) decode(ptr unsafe.Pointer, b *fast.BinaryBufferReader, _ bool) error {
	*(*uint32)(unsafe.Add(ptr, f.offset)) = b.ReadUint32()
	return nil
}

type uint64Type struct {
	offset uintptr
}

func (uint64Type) encodedSize(_ unsafe.Pointer) int {
	return 8
}

func (f uint64Type) encode(ptr unsafe.Pointer, b *fast.BinaryBuffer) {
	b.WriteUint64(*(*uint64)(unsafe.Add(ptr, f.offset)))
}

func (f uint64Type) decode(ptr unsafe.Pointer, b *fast.BinaryBufferReader, _ bool) error {
	*(*uint64)(unsafe.Add(ptr, f.offset)) = b.ReadUint64()
	return nil
}
