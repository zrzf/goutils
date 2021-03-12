package slices

import (
	// "fmt"
	"testing"
)

func TestIndexStr(t *testing.T) {
	a := []string{"a", "b", "c"}
	if IndexStr(a, "b") != 1 {
		t.Fatal()
	}
}

func TestIndexInt(t *testing.T) {
	a := []int{2, 3, 1}
	if IndexInt(a, 1) != 2 {
		t.Fatal()
	}
}

func TestContainInt(t *testing.T) {
	a := []int{2, 3, 1}
	if ContainInt(a, 2) != true {
		t.Fatal()
	}
}

func TestContainStr(t *testing.T) {
	a := []string{"a", "b", "c"}
	if ContainStr(a, "b") != true {
		t.Fatal()
	}
}

func TestUniqueStr(t *testing.T) {
	a := []string{"a", "b", "c", "a", "a"}
	UniqueStr(&a)
	if len(a) != 3 {
		t.Fatal()
	}
}

func TestFilterStr(t *testing.T) {
	a := []string{"a", "b", "c", "a", "a"}
	FilterStr(&a, "a")
	if len(a) != 2 {
		t.Fatal()
	}

}

func TestFilterInt(t *testing.T) {
	a := []int{1, 2, 3, 1, 1}
	FilterInt(&a, 1)
	if len(a) != 2 {
		t.Fatal()
	}
}

func TestRangeInt(t *testing.T) {
	if len(RangeInt(0, 5)) != 6 {
		t.Fatal()
	}
}

func TestRangeUint(t *testing.T) {
	if len(RangeUint(0, 5)) != 6 {
		t.Fatal()
	}
}

func TestShuffleInt(t *testing.T) {
	a := RangeInt(0, 30)
	b := make([]int, len(a))
	copy(b, a)
	ShuffleInt(a, 12345)
	if a[0] == b[0] && a[1] == b[1] {
		t.Fatal()
	}
}
