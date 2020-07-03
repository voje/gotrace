package main

func someBytes(b []byte) {
	return
}

func someUInts(a, b, c []uint8) {
	return
}

func main() {
	someBytes([]byte{0x61, 0x62, 0x63})
	someUInts(
		[]uint8{1, 2, 3},
		[]uint8{88, 89, 99},
		[]uint8{0xff, 0x32, 0x1f},
	)
}
