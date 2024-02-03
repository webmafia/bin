package main

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/viant/xunsafe"
)

func BenchmarkReflectTypeOf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = reflect.TypeOf(Foo{})
	}
}

func BenchmarkReflectFieldByName(b *testing.B) {
	fooType := reflect.TypeOf(Foo{})
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = xunsafe.FieldByName(fooType, "ID")
	}
}

func BenchmarkReflectFieldValue(b *testing.B) {
	fooType := reflect.TypeOf(Foo{})
	fooID := xunsafe.FieldByName(fooType, "ID")
	foo := Foo{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = fooID.Int(unsafe.Pointer(&foo))
	}
}

func BenchmarkReflectFieldInterfaceValue(b *testing.B) {
	fooType := reflect.TypeOf(Foo{})
	field := xunsafe.FieldByName(fooType, "Name")
	foo := Foo{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = field.Value(unsafe.Pointer(&foo))
	}
}
