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
	return *(*iface)(noescape(unsafe.Pointer(&v)))
}

//go:inline
func fromIface(v iface) any {
	return *(*any)(unsafe.Pointer(&v))
}

// noescape hides a pointer from escape analysis. It is the identity function
// but escape analysis doesn't think the output depends on the input.
// noescape is inlined and currently compiles down to zero instructions.
// USE CAREFULLY!
// This was copied from the runtime; see issues 23382 and 7921.
//
//go:nosplit
//go:nocheckptr
func noescape(p unsafe.Pointer) unsafe.Pointer {
	x := uintptr(p)
	return unsafe.Pointer(x ^ 0)
}
