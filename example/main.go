package main

import (
	"log"
	"reflect"

	bin "github.com/webbmaffian/go-binary"
	"github.com/webbmaffian/go-fast"
)

type Foo struct {
	Name string
	ID   int
	Bar  Bar
	mjau string
}

type Bar struct {
	Baz []Outer
}

type Outer struct {
	A uint64
	B uint64
	C Inner
}

type Inner struct {
	D uint64
}

// func main() {
// 	outer := reflect.TypeOf(Outer{})
// 	inner := reflect.TypeOf(Inner{})
// 	a := xunsafe.FieldByName(outer, "A")
// 	b := xunsafe.FieldByName(outer, "B")
// 	c := xunsafe.FieldByName(outer, "C")
// 	d := xunsafe.FieldByName(inner, "D")
// 	d.Offset = c.Offset

// 	_, _, _, _ = a, b, c, d

// 	val := Outer{
// 		A: 1,
// 		B: 2,
// 		C: Inner{
// 			D: 3,
// 		},
// 	}

// 	log.Println(d.Uint64(unsafe.Pointer(&val)))
// }

func main() {
	var key [32]byte

	typs, err := bin.NewTypes(key[:])

	if err != nil {
		panic(err)
	}

	err = typs.Register(reflect.TypeOf(Foo{}))

	if err != nil {
		panic(err)
	}

	b := fast.NewBinaryBuffer(1024)
	f := Foo{
		Name: "my name is foo",
		ID:   123,
		mjau: "mjau",
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

	if err = bin.Encode(typs, b, &f); err != nil {
		panic(err)
	}

	log.Println(b.String())
	log.Println(b.Bytes())

	var f2 Foo
	f2.Bar.Baz = make([]Outer, 0, 2)
	r := fast.NewBinaryBufferReader(b)

	if err = bin.Decode(typs, &r, &f2); err != nil {
		panic(err)
	}

	log.Printf("%+v\n", f2)
}
