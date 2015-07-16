package lifegame

import (
	"testing"
)

func TestReadRLE(t *testing.T) {
	size := 1000
	bitmap := NewBitMap(size, size)
	err := ReadRLE("../broken-lines.rle", bitmap)
	if err != nil {
		t.Error(err)
	}
}

func TestBitMap(t *testing.T) {
	size := 1000
	bitmap := NewBitMap(size, size)
	var i, j uint64
	for i = 0; i < uint64(size); i++ {
		for j = 0; j < uint64(size); j++ {
			bitmap.Set(i, j, true)
			if !bitmap.Get(i, j) {
				t.Error("should be true %d %d", i, j)
				t.Fail()
			}
			bitmap.Set(i, j, false)
			if bitmap.Get(i, j) {
				t.Error("should be false %d %d", i, j)
				t.Fail()
			}
		}
	}
}
