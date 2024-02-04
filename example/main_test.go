package main

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/viant/xunsafe"
	bin "github.com/webbmaffian/go-binary"
	"github.com/webbmaffian/go-fast"
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

func BenchmarkEncode(b *testing.B) {
	var key [32]byte

	typs, err := bin.NewTypes(key[:])

	if err != nil {
		panic(err)
	}

	_, err = typs.Register(reflect.TypeOf(Foo{}))

	if err != nil {
		b.Fatal(err)
	}

	buf := fast.NewBinaryBuffer(1024)
	f := Foo{
		Name: "my name is foo",
		ID:   123,
		Bar: Bar{
			Baz: 456,
		},
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		bin.Encode(typs, buf, &f)
		buf.Reset()
	}
}

func BenchmarkDecode(b *testing.B) {
	var key [32]byte

	typs, err := bin.NewTypes(key[:])

	if err != nil {
		panic(err)
	}

	_, err = typs.Register(reflect.TypeOf(Foo{}))

	if err != nil {
		b.Fatal(err)
	}

	buf := fast.NewBinaryBuffer(1024)
	f := Foo{
		Name: "my name is foo",
		ID:   123,
		Bar: Bar{
			Baz: 456,
		},
	}

	bin.Encode(typs, buf, &f)

	var f2 Foo
	r := fast.NewBinaryBufferReader(buf)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		bin.Decode(typs, &r, &f2)
		r.Reset()
	}
}
