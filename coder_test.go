package bin

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"runtime"
	"testing"

	"github.com/webmafia/fast/binary"
)

func ExampleCoder() {
	type Foo struct {
		Name string
		ID   int
		mjau string
	}

	c := NewCoder()

	b := binary.NewBufferWriter(1024)

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
	type Baz struct {
		Name string
	}

	type Bar struct {
		A Baz
		B Baz
	}

	type Foo struct {
		Items []Bar
	}

	c := NewCoder(CoderOptions{
		AllowAllocations: true,
	})

	var buf bytes.Buffer
	b := binary.NewStreamWriter(&buf)

	err := c.Encode(b, &Foo{
		Items: []Bar{
			{
				A: Baz{Name: "a1"},
				B: Baz{Name: "b1"},
			},
			{
				A: Baz{Name: "a2"},
				B: Baz{Name: "b2"},
			},
			{
				A: Baz{Name: "a3"},
				B: Baz{Name: "b3"},
			},
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	b.Flush()

	fmt.Println(buf.Len(), buf.String())
	fmt.Println(buf.Len(), buf.Bytes())

	r := binary.NewBufferReader(buf.Bytes())

	var dst Foo

	fmt.Printf("len %d, cap %d: %#v\n", len(dst.Items), cap(dst.Items), dst)

	if err = c.Decode(&r, &dst); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("len %d, cap %d: %#v\n", len(dst.Items), cap(dst.Items), dst)
	r.Reset()

	if err = c.Decode(&r, &dst); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("len %d, cap %d: %#v\n", len(dst.Items), cap(dst.Items), dst)

	runtime.GC()

	fmt.Printf("len %d, cap %d: %#v\n", len(dst.Items), cap(dst.Items), dst)

	// Output: Todo
}

func BenchmarkCoderEncode(b *testing.B) {
	type Foo struct {
		Name string
		ID   int
		mjau string
	}

	c := NewCoder()

	buf := binary.NewBufferWriter(1024)

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

	buf := binary.NewStreamWriter(io.Discard)

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

	buf := binary.NewBufferWriter(1024)

	err := c.Encode(buf, &Foo{
		Name: "mjau",
	})

	if err != nil {
		b.Fatal(err)
	}

	r := binary.NewBufferReader(buf.Bytes())

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

	buf := binary.NewBufferWriter(1024)

	err := c.Encode(buf, &Foo{
		Name: "mjau",
	})

	if err != nil {
		b.Fatal(err)
	}

	r := binary.NewBufferReader(buf.Bytes())

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
		buf := binary.NewBufferWriter(1024)

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
	buf := binary.NewBufferWriter(1024)

	err := c.Encode(buf, &Foo{
		Name: "mjau",
	})

	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	b.RunParallel(func(p *testing.PB) {
		r := binary.NewBufferReader(buf.Bytes())
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
	buf := binary.NewBufferWriter(1024)

	err := c.Encode(buf, &Foo{
		Name: "mjau",
	})

	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	b.RunParallel(func(p *testing.PB) {
		r := binary.NewBufferReader(buf.Bytes())
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
