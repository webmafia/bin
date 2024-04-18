package bin

import (
	"encoding/gob"
	"fmt"
	"io"
	"testing"
)

type myStruct struct {
	Slice []mySubStruct
}

type mySubStruct struct {
	Foo string
	Bar float64
	Baz [][]uint16
}

func genBenchStruct(v *myStruct, n int) {
	v.Slice = make([]mySubStruct, n)

	for i := range v.Slice {
		v.Slice[i] = mySubStruct{
			Foo: "foobar",
			Bar: -123.456,
			Baz: make([][]uint16, 16),
		}

		for j := range v.Slice[i].Baz {
			v.Slice[i].Baz[j] = make([]uint16, 16)

			for k := range v.Slice[i].Baz[j] {
				v.Slice[i].Baz[j][k] = uint16(i + j + k)
			}
		}
	}
}

func BenchmarkEncoder_Stream(b *testing.B) {
	sizes := []int{0, 8, 16, 32, 64, 128, 256}

	b.ResetTimer()

	for _, s := range sizes {
		b.Run(fmt.Sprintf("slice-%d", s), func(b *testing.B) {
			var src myStruct
			genBenchStruct(&src, s)

			b.Run("bin", func(b *testing.B) {
				enc, err := NewEncoder(io.Discard, &src)

				if err != nil {
					b.Fatal(err)
				}

				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					err := enc(&src)

					if err != nil {
						b.Fatal(err)
					}
				}
			})

			b.Run("gob", func(b *testing.B) {
				c := gob.NewEncoder(io.Discard)
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					err := c.Encode(&src)

					if err != nil {
						b.Fatal(err)
					}
				}
			})
		})
	}
}

func genBenchSlice(v *[][][]int, n int) {
	*v = make([][][]int, n)

	for i := range *v {
		(*v)[i] = make([][]int, n)

		for j := range (*v)[i] {
			(*v)[i][j] = make([]int, n)

			for k := range (*v)[i][j] {
				(*v)[i][j][k] = i + j + k
			}
		}
	}
}

func BenchmarkEncoder_NestedSlices(b *testing.B) {
	sizes := []int{0, 8, 16, 32, 64, 128, 256}

	b.ResetTimer()

	for _, s := range sizes {
		b.Run(fmt.Sprintf("slice-%d", s), func(b *testing.B) {
			var src [][][]int
			genBenchSlice(&src, s)

			b.Run("bin", func(b *testing.B) {
				enc, err := NewEncoder(io.Discard, &src)

				if err != nil {
					b.Fatal(err)
				}

				b.ResetTimer()
				b.ReportMetric(float64(s*s*s)/float64(b.N), "iter/op")

				for i := 0; i < b.N; i++ {
					err := enc(&src)

					if err != nil {
						b.Fatal(err)
					}
				}
			})

			b.Run("gob", func(b *testing.B) {
				c := gob.NewEncoder(io.Discard)
				b.ResetTimer()
				b.ReportMetric(float64(s*s*s)/float64(b.N), "iter/op")

				for i := 0; i < b.N; i++ {
					err := c.Encode(&src)

					if err != nil {
						b.Fatal(err)
					}
				}
			})
		})
	}
}

func BenchmarkEncoder_Stream_Lite(b *testing.B) {
	type testStruct struct {
		Foo  string
		Bar  float64
		Foo2 string
		Bar2 float64
		Foo3 string
		Bar3 float64
	}

	// src := testStruct{
	// 	Foo:  "hello world",
	// 	Bar:  -123.456,
	// 	Foo2: "hello world",
	// 	Bar2: -123.456,
	// 	Foo3: "hello world",
	// 	Bar3: -123.456,
	// }

	src := []testStruct{
		{
			Foo: "hello world",
			Bar: -123.456,
		},
	}

	b.ResetTimer()

	b.Run("bin", func(b *testing.B) {
		enc, err := NewEncoder(io.Discard, &src)

		if err != nil {
			b.Fatal(err)
		}

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			err := enc(&src)

			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("gob", func(b *testing.B) {
		c := gob.NewEncoder(io.Discard)
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			err := c.Encode(&src)

			if err != nil {
				b.Fatal(err)
			}
		}
	})
}
