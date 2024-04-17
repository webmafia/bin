package bin

import (
	"errors"
	"hash"
	"reflect"
	"sync/atomic"
	"unsafe"

	"github.com/webmafia/fast/binary"
)

type Coder struct {
	items [64]unsafe.Pointer
	opt   CoderOptions
}

type CoderOptions struct {
	AllowAllocations     bool
	KeepUnexportedFields bool
}

func NewCoder(opt ...CoderOptions) *Coder {
	c := &Coder{}

	if len(opt) > 0 {
		c.opt = opt[0]
	}

	return c
}

func (c *Coder) Encode(b binary.Writer, v any) (err error) {
	typ := reflect.TypeOf(v)

	if typ.Kind() != reflect.Pointer {
		return errors.New("expected pointer")
	}

	ifs := toIface(v)
	t, err := c.getType(uintptr(ifs.tab), typ.Elem())

	if err != nil {
		return
	}

	t.encode(ifs.data, b)

	return
}

func (c *Coder) Decode(b binary.Reader, v any, nocopy ...bool) (err error) {
	typ := reflect.TypeOf(v)

	if typ.Kind() != reflect.Pointer {
		return errors.New("expected pointer")
	}

	ifs := toIface(v)
	t, err := c.getType(uintptr(ifs.tab), typ.Elem())

	if err != nil {
		return
	}

	return t.decode(ifs.data, b, len(nocopy) > 0 && nocopy[0])
}

func (c *Coder) TypeHash(h hash.Hash, v any) (err error) {
	typ := reflect.TypeOf(v)

	if typ.Kind() != reflect.Pointer {
		return errors.New("expected pointer")
	}

	ifs := toIface(v)
	t, err := c.getType(uintptr(ifs.tab), typ.Elem())

	if err != nil {
		return
	}

	t.typeHash(h)

	return
}

func (c *Coder) getType(tab uintptr, typ reflect.Type) (t Type, err error) {
	idx := tab % 64

	pp := &c.items[idx]
	p := atomic.LoadPointer(pp)

	for p != nil {
		it := (*item)(p)

		if it.tab == tab {
			return it.typ, nil
		}

		pp = &it.next
		p = atomic.LoadPointer(pp)
	}

	// If we came here, no coder exists - create one
	t, err = getType(typ, 0, &c.opt)

	if err != nil {
		return
	}

	it := &item{
		tab: tab,
		typ: t,
	}

	p = atomic.SwapPointer(pp, unsafe.Pointer(it))

	// The pointer might have been swapped by another thread
	for p != nil {
		it = (*item)(p)
		pp = &it.next
		p = atomic.SwapPointer(pp, p)
	}

	return
}

type item struct {
	tab  uintptr
	next unsafe.Pointer
	typ  Type
}
