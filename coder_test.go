package bin

import (
	"fmt"
	"testing"

	"github.com/webbmaffian/go-fast"
)

func ExampleCoder() {
	type Foo struct {
		Name string
		ID   int
		mjau string
	}

	var key [32]byte

	c, err := NewCoder(key[:])

	if err != nil {
		panic(err)
	}

	b := fast.NewBinaryBuffer(1024)

	err = c.Encode(b, &Foo{
		Name: "mjau",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(b.String())
	fmt.Println(b.Bytes())

	// Output: Unknown
}

func BenchmarkCoder(b *testing.B) {
	type Foo struct {
		Name string
		ID   int
		mjau string
	}

	var key [32]byte

	c, err := NewCoder(key[:])

	if err != nil {
		b.Fatal(err)
	}

	buf := fast.NewBinaryBuffer(1024)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err = c.Encode(buf, &Foo{
			Name: "mjau",
		})

		if err != nil {
			b.Fatal(err)
		}

		buf.Reset()
	}
}

func BenchmarkCoderParallell(b *testing.B) {
	type Foo struct {
		Name string
		ID   int
		mjau string
	}

	var key [32]byte

	c, err := NewCoder(key[:])

	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	b.RunParallel(func(p *testing.PB) {
		buf := fast.NewBinaryBuffer(1024)

		for p.Next() {
			err := c.Encode(buf, &Foo{
				Name: "mjau",
			})

			if err != nil {
				b.Fatal(err)
			}

			buf.Reset()
		}
	})
}
