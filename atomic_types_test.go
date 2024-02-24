package bin

import (
	"fmt"
	"testing"

	"github.com/webbmaffian/go-fast"
)

func ExampleAtomicTypes() {
	type Foo struct {
		Name string
		ID   int
		mjau string
	}

	var key [32]byte
	typs, err := NewAtomicTypes(key[:])

	if err != nil {
		panic(err)
	}

	// c, err := typs.GetCoder(Foo{})

	// if err != nil {
	// 	panic(err)
	// }

	b := fast.NewBinaryBuffer(1024)

	if err = Encode2(typs, b, &Foo{
		Name: "mjau",
	}); err != nil {
		panic(err)
	}

	fmt.Println(b.String())
	fmt.Println(b.Bytes())
	b.Reset()

	if err = Encode2(typs, b, &Foo{
		Name: "mjau2",
	}); err != nil {
		panic(err)
	}

	fmt.Println(b.String())
	fmt.Println(b.Bytes())

	// Output: Unknown
}

func BenchmarkAtomicTypes(b *testing.B) {
	type Foo struct {
		Name string
		ID   int
		mjau string
	}

	var key [32]byte
	typs, err := NewAtomicTypes(key[:])

	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		c, err := typs.GetCoder(&Foo{})
		_, _ = c, err
	}
}

func BenchmarkAtomicTypesParallell(b *testing.B) {
	type Foo struct {
		Name string
		ID   int
		mjau string
	}

	var key [32]byte
	typs, err := NewAtomicTypes(key[:])

	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			c, err := typs.GetCoder(&Foo{})
			if err != nil {
				b.Fatal(err)
			}
			_, _ = c, err
		}
	})
}
