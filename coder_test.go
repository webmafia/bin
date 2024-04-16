package bin

import (
	"fmt"
	"io"
	"log"
	"testing"

	"github.com/webmafia/fast"
)

func ExampleCoder() {
	type Foo struct {
		Name string
		ID   int
		mjau string
	}

	c := NewCoder()

	b := fast.NewBinaryBuffer(1024)

	err := c.Encode(b, &Foo{
		Name: "mjau",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(b.String())
	fmt.Println(b.Bytes())

	// Output: Todo
}

func ExampleCoder_Slice() {
	type Foo struct {
		Names []string
	}

	c := NewCoder(CoderOptions{
		AllowAllocations: true,
	})

	b := fast.NewBinaryBuffer(1024)

	err := c.Encode(b, &Foo{
		Names: []string{"foo", "bar", "baz"},
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(b.String())
	fmt.Println(b.Bytes())

	r := fast.NewBinaryBufferReader(b)

	var dst Foo

	if err = c.Decode(&r, &dst); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("len %d, cap %d: %#v\n", len(dst.Names), cap(dst.Names), dst)
	r.Reset()

	if err = c.Decode(&r, &dst); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("len %d, cap %d: %#v\n", len(dst.Names), cap(dst.Names), dst)

	// Output: Todo
}

func BenchmarkCoderEncode(b *testing.B) {
	type Foo struct {
		Name string
		ID   int
		mjau string
	}

	c := NewCoder()

	buf := fast.NewBinaryBuffer(1024)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := c.Encode(buf, &Foo{
			Name: "mjau",
		})

		if err != nil {
			b.Fatal(err)
		}

		buf.Reset()
	}
}

func BenchmarkCoderEncode_Stream(b *testing.B) {
	type Foo struct {
		Name string
		ID   int
		mjau string
	}

	c := NewCoder()

	buf := fast.NewBinaryStream(io.Discard)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := c.Encode(buf, &Foo{
			Name: "mjau",
		})

		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCoderDecode(b *testing.B) {
	type Foo struct {
		Name string
		ID   int
		mjau string
	}

	c := NewCoder()

	buf := fast.NewBinaryBuffer(1024)

	err := c.Encode(buf, &Foo{
		Name: "mjau",
	})

	if err != nil {
		b.Fatal(err)
	}

	r := fast.NewBinaryBufferReader(buf)

	var f Foo

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err = c.Decode(&r, &f)

		if err != nil {
			b.Fatal(err)
		}

		r.Reset()
	}
}

func BenchmarkCoderDecodeNocopy(b *testing.B) {
	type Foo struct {
		Name string
		ID   int
		mjau string
	}

	c := NewCoder()

	buf := fast.NewBinaryBuffer(1024)

	err := c.Encode(buf, &Foo{
		Name: "mjau",
	})

	if err != nil {
		b.Fatal(err)
	}

	r := fast.NewBinaryBufferReader(buf)

	var f Foo

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err = c.Decode(&r, &f, true)

		if err != nil {
			b.Fatal(err)
		}

		r.Reset()
	}
}

func BenchmarkCoderParallell(b *testing.B) {
	type Foo struct {
		Name string
		ID   int
		mjau string
	}

	c := NewCoder()
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

func BenchmarkCoderParallellDecode(b *testing.B) {
	type Foo struct {
		Name string
		ID   int
		mjau string
	}

	c := NewCoder()
	buf := fast.NewBinaryBuffer(1024)

	err := c.Encode(buf, &Foo{
		Name: "mjau",
	})

	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	b.RunParallel(func(p *testing.PB) {
		r := fast.NewBinaryBufferReader(buf)
		var f Foo

		for p.Next() {
			err := c.Decode(&r, &f)

			if err != nil {
				b.Fatal(err)
			}

			r.Reset()
		}
	})
}

func BenchmarkCoderParallellDecodeNocopy(b *testing.B) {
	type Foo struct {
		Name string
		ID   int
		mjau string
	}

	c := NewCoder()
	buf := fast.NewBinaryBuffer(1024)

	err := c.Encode(buf, &Foo{
		Name: "mjau",
	})

	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	b.RunParallel(func(p *testing.PB) {
		r := fast.NewBinaryBufferReader(buf)
		var f Foo

		for p.Next() {
			err := c.Decode(&r, &f, true)

			if err != nil {
				b.Fatal(err)
			}

			r.Reset()
		}
	})
}
