package core

import "testing"

func runUint2Bytes(ui uint, e []byte, t *testing.T) {
	for i, v := range uint2Bytes(ui) {
		if v != e[i] {
			t.Error("AAAAh")
		}
	}
}

func runBytes2Uint(bs []byte, e uint, t *testing.T) {
	if bytes2Uint(bs) != e {
		t.Error("AAAAAAAAAHHHHHHHH")
	}
}

func TestIo(t *testing.T) {
	runUint2Bytes(1023, []byte{1, 0, 2, 3}, t)
	runUint2Bytes(99, []byte{0, 0, 9, 9}, t)
	runBytes2Uint([]byte{1, 0, 2, 3}, 1023, t)
	runBytes2Uint([]byte{9, 9, 9, 9, 9}, 99999, t)
	runBytes2Uint([]byte{1, 0, 2, 0}, 1020, t)
	runBytes2Uint([]byte{0, 0, 0, 0, 0, 2, 0}, 20, t)
}
