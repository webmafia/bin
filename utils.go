package bin

import "unsafe"

func sizeUvarint(x uint64) (l int) {
	for x >= 0x80 {
		l++
		x >>= 7
	}

	l++

	return
}

type iface struct {
	tab  unsafe.Pointer
	data unsafe.Pointer
}

//go:inline
func toIface(v any) iface {
	return *(*iface)(unsafe.Pointer(&v))
}

//go:inline
func fromIface(v iface) any {
	return *(*any)(unsafe.Pointer(&v))
}
