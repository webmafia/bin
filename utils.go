package bin

func sizeUvarint(x uint64) (l int) {
	for x >= 0x80 {
		l++
		x >>= 7
	}

	l++

	return
}

func sizeVarint(x int64) int {
	ux := uint64(x) << 1
	if x < 0 {
		ux = ^ux
	}

	return sizeUvarint(ux)
}
