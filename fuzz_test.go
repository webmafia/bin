package bin

import (
	"bytes"
	"fmt"
	"math/rand"
	"testing"

	"github.com/webmafia/fast"
	"github.com/webmafia/fast/binary"
)

func TestEncoder(t *testing.T) {
	var src bigStruct
	rand := rand.New(rand.NewSource(1))

	for i := 0; i < 128; i++ {
		if err := testEncoder(&src, rand, i); err != nil {
			t.Fatal(err)
		}
	}
}

func testEncoder(src *bigStruct, rand *rand.Rand, maxSlice int) (err error) {
	generateStruct(src, rand, maxSlice)

	c := NewCoder(Options{
		AllowAllocations:     true,
		KeepUnexportedFields: true,
	})

	var buf bytes.Buffer
	w := binary.NewStreamWriter(&buf)

	if err = c.Encode(w, src); err != nil {
		return
	}

	var dst bigStruct
	r := binary.NewBufferReader(buf.Bytes())

	if err = c.Decode(r, &dst); err != nil {
		return
	}

	return src.compare(&dst)
}

type bigStruct struct {
	items []bigStructItem
}

type bigStructItem struct {
	bigSlice    []string
	integer     uint16
	nestedSlice [][][]int
}

func (a *bigStruct) compare(b *bigStruct) (err error) {
	if len(a.items) != len(b.items) {
		return fmt.Errorf("items length mismatch; %d vs %d", len(a.items), len(b.items))
	}

	for i := range a.items {
		if len(a.items[i].bigSlice) != len(b.items[i].bigSlice) {
			return fmt.Errorf("items.%d.bigSlice length mismatch; %d vs %d", i, len(a.items[i].bigSlice), len(b.items[i].bigSlice))
		}

		for j := range a.items[i].bigSlice {
			if a.items[i].bigSlice[j] != b.items[i].bigSlice[j] {
				return fmt.Errorf("items.%d.bigSlice.%d string mismatch; '%s' vs '%s'", i, j, a.items[i].bigSlice[j], b.items[i].bigSlice[j])
			}
		}

		if a.items[i].integer != b.items[i].integer {
			return fmt.Errorf("items.%d.integer integer mismatch; %d vs %d", i, a.items[i].integer, b.items[i].integer)
		}

		if len(a.items[i].nestedSlice) != len(b.items[i].nestedSlice) {
			return fmt.Errorf("items.%d.nestedSlice length mismatch; %d vs %d", i, len(a.items[i].nestedSlice), len(b.items[i].nestedSlice))
		}

		for j := range a.items[i].nestedSlice {
			if len(a.items[i].nestedSlice[j]) != len(b.items[i].nestedSlice[j]) {
				return fmt.Errorf("items.%d.nestedSlice.%d length mismatch; %d vs %d", i, j, len(a.items[i].nestedSlice[j]), len(b.items[i].nestedSlice[j]))
			}

			for k := range a.items[i].nestedSlice[j] {
				if len(a.items[i].nestedSlice[j][k]) != len(b.items[i].nestedSlice[j][k]) {
					return fmt.Errorf("items.%d.nestedSlice.%d.%d length mismatch; %d vs %d", i, j, k, len(a.items[i].nestedSlice[j][k]), len(b.items[i].nestedSlice[j][k]))
				}

				for l := range a.items[i].nestedSlice[j][k] {
					if a.items[i].nestedSlice[j][k][l] != b.items[i].nestedSlice[j][k][l] {
						return fmt.Errorf("items.%d.nestedSlice.%d.%d.%d integer mismatch; %d vs %d", i, j, k, l, a.items[i].nestedSlice[j][k][l], b.items[i].nestedSlice[j][k][l])
					}
				}
			}
		}
	}

	return
}

func generateStruct(v *bigStruct, rand *rand.Rand, maxSlice int) {
	randomResize(&v.items, rand, maxSlice)

	for i := range v.items {
		randomResize(&v.items[i].bigSlice, rand, maxSlice)

		for j := range v.items[i].bigSlice {
			str := make([]byte, rand.Intn(128))
			rand.Read(str)
			v.items[i].bigSlice[j] = fast.BytesToString(str)
		}

		v.items[i].integer = uint16(i)

		generateNestedSlice(&v.items[i].nestedSlice, rand, maxSlice)
	}

}

func generateNestedSlice(v *[][][]int, rand *rand.Rand, maxSlice int) {
	randomResize(v, rand, maxSlice)

	for j := range *v {
		randomResize(&(*v)[j], rand, maxSlice)

		for k := range (*v)[j] {
			randomResize(&(*v)[j][k], rand, maxSlice)

			for l := range (*v)[j][k] {
				(*v)[j][k][l] = rand.Int()
			}
		}
	}
}

func randomResize[T any](v *[]T, rand *rand.Rand, maxRand int) {
	if maxRand > 0 {
		resize(v, rand.Intn(maxRand))
	} else {
		resize(v, 0)
	}
}

func resize[T any](v *[]T, size int) {
	if cap(*v) >= size {
		*v = (*v)[:size]
	} else {
		*v = make([]T, size)
	}
}
