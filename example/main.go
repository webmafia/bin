package main

import (
	"log"
	"reflect"

	bin "github.com/webbmaffian/go-binary"
)

type Foo struct {
	Name string
	ID   int
	Bar  Bar
}

type Bar struct {
	Baz []int
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

	hash, err := typs.Register(reflect.TypeOf(Foo{}))

	if err != nil {
		panic(err)
	}

	log.Println("Hash:", hash)
}
