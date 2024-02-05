package main

import (
	"reflect"
	"testing"

	"github.com/webbmaffian/go-bin"
	"github.com/webbmaffian/go-fast"
)

func BenchmarkEncode(b *testing.B) {
	var key [32]byte

	typs, err := bin.NewTypes(key[:])

	if err != nil {
		panic(err)
	}

	err = typs.Register(reflect.TypeOf(Foo{}))

	if err != nil {
		b.Fatal(err)
	}

	buf := fast.NewBinaryBuffer(1024)
	f := Foo{
		Name: "my name is foo",
		ID:   123,
		Bar: Bar{
			Baz: []Outer{
				{
					A: 1,
					B: 2,
					C: Inner{
						D: 3,
					},
				},
			},
		},
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		bin.Encode(typs, buf, &f)
		buf.Reset()
	}
}

func BenchmarkEncode2(b *testing.B) {
	var key [32]byte

	type simple struct {
		A string
	}

	typs, err := bin.NewTypes(key[:])

	if err != nil {
		panic(err)
	}

	err = typs.Register(reflect.TypeOf(simple{}))

	if err != nil {
		b.Fatal(err)
	}

	buf := fast.NewBinaryBuffer(1024)
	s := simple{}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		bin.Encode(typs, buf, &s)
		buf.Reset()
	}
}

func BenchmarkGetEncoder(b *testing.B) {
	var key [32]byte

	type simple struct {
		A string
	}

	typs, err := bin.NewTypes(key[:])

	if err != nil {
		panic(err)
	}

	err = typs.Register(reflect.TypeOf(simple{}))

	if err != nil {
		b.Fatal(err)
	}

	// buf := fast.NewBinaryBuffer(1024)
	s := simple{}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = typs.GetCoder(reflect.TypeOf(s))
	}
}

func BenchmarkDecode(b *testing.B) {
	var key [32]byte

	typs, err := bin.NewTypes(key[:])

	if err != nil {
		panic(err)
	}

	err = typs.Register(reflect.TypeOf(Foo{}))

	if err != nil {
		b.Fatal(err)
	}

	buf := fast.NewBinaryBuffer(1024)
	f := Foo{
		Name: "my name is foo",
		ID:   123,
		Bar: Bar{
			Baz: []Outer{
				{
					A: 1,
					B: 2,
					C: Inner{
						D: 3,
					},
				},
			},
		},
	}

	bin.Encode(typs, buf, &f)

	var f2 Foo
	f2.Bar.Baz = make([]Outer, 0, 2)
	r := fast.NewBinaryBufferReader(buf)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		bin.Decode(typs, &r, &f2)
		r.Reset()
		f2.Bar.Baz = f2.Bar.Baz[:0]
	}
}
