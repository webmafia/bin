package bin

import (
	"bytes"
	"math/rand"
	"reflect"
	"testing"

	"github.com/webmafia/bin"
	"github.com/webmafia/fast"
	"github.com/webmafia/fast/binary"
)

type bigStruct struct {
	items []item
}

type item struct {
	bigSlice    []string
	nestedSlice [][][]int
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

		randomResize(&v.items[i].nestedSlice, rand, maxSlice)

		for j := range v.items[i].nestedSlice {
			randomResize(&v.items[i].nestedSlice[j], rand, maxSlice)

			for k := range v.items[i].nestedSlice[j] {
				randomResize(&v.items[i].nestedSlice[j][k], rand, maxSlice)

				for l := range v.items[i].nestedSlice[j][k] {
					v.items[i].nestedSlice[j][k][l] = rand.Int()
				}
			}
		}
	}

}

func randomResize[T any](v *[]T, rand *rand.Rand, maxRand int) {
	resize(v, rand.Intn(maxRand))
}

func resize[T any](v *[]T, size int) {
	if cap(*v) >= size {
		*v = (*v)[:size]
	} else {
		*v = make([]T, size)
	}
}

func TestStream(t *testing.T) {
	var src bigStruct
	rand := rand.New(rand.NewSource(1))

	generateStruct(&src, rand, 2)

	c := bin.NewCoder(bin.CoderOptions{
		AllowAllocations:     true,
		KeepUnexportedFields: true,
	})

	var buf bytes.Buffer
	w := binary.NewStreamWriter(&buf)

	if err := c.Encode(w, &src); err != nil {
		t.Fatal(err)
	}

	var dst bigStruct
	r := binary.NewBufferReader(buf.Bytes())

	if err := c.Decode(&r, &dst); err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(src, dst) {
		t.Fatal("src and dst are not equal")
	}
}
