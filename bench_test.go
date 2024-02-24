package bin

import (
	"fmt"
	"reflect"
	"testing"
)

func BenchmarkReflectType_Name(b *testing.B) {
	type Waza struct {
		Foobar string
	}

	t := reflect.TypeOf(Waza{})

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = t.Name()
	}
}

func BenchmarkReflectType_PkgPath(b *testing.B) {
	type Waza struct {
		Foobar string
	}

	t := reflect.TypeOf(Waza{})

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = t.PkgPath()
	}
}

func BenchmarkReflectType_String(b *testing.B) {
	type Waza struct {
		Foobar string
	}

	t := reflect.TypeOf(Waza{})

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = t.String()
	}
}

func BenchmarkUnsafeTab(b *testing.B) {
	type Waza struct {
		Foobar string
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = uintptr(toIface(Waza{}).tab)
	}
}

type itab struct {
	inter uintptr
	_type uintptr
	hash  uint32
	_     uint32
	fun   [1]uintptr
}

func tab(v any) *itab {
	return (*itab)(toIface(v).tab)
}

func ExampleIface() {
	type Foo struct {
		Foobar string
	}

	type Bar struct {
		Foobar string
	}

	type Baz struct {
		Foobar uint64
	}

	type Baz2 struct {
		Foobar  int64
		Foobar2 int64
	}

	fmt.Printf("Foo: %+v\n", tab(Foo{}))
	fmt.Printf("*Foo: %+v\n", tab(&Foo{}))

	// Output: Mjau
}

func BenchmarkReflectElem(b *testing.B) {
	type Foo struct {
		Foobar string
	}

	var f Foo

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = reflect.TypeOf(&f).Elem()
	}
}

func BenchmarkReflectKind(b *testing.B) {
	type Foo struct {
		Foobar string
	}

	var f Foo

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = reflect.TypeOf(&f).Kind()
	}
}
