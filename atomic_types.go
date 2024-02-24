package bin

import (
	"hash"
	"reflect"
	"sync/atomic"
	"unsafe"

	"github.com/minio/highwayhash"
)

type AtomicTypes struct {
	items [64]unsafe.Pointer
	hash  hash.Hash64
}

func NewAtomicTypes(key []byte) (t *AtomicTypes, err error) {
	hash, err := highwayhash.New64(key)

	if err != nil {
		return
	}

	t = &AtomicTypes{
		hash: hash,
	}

	return
}

type item struct {
	tab  uintptr
	next unsafe.Pointer
	typ  Type
}

func (t *AtomicTypes) GetCoder(typ any) (c Type, err error) {
	ifs := toIface(typ)
	tab := uintptr(ifs.tab)
	idx := tab % 64

	pp := &t.items[idx]
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
	c, err = getType(reflect.TypeOf(typ).Elem(), 0, func(k reflect.Kind) {
		t.hash.Write([]byte{byte(k)})
	})

	if err != nil {
		return
	}

	it := &item{
		tab: tab,
		typ: c,
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
