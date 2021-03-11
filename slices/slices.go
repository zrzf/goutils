package slices

import (
	"math/rand"
)

func SumFloat64(list []float64) float64 {
	var s float64
	for _, v := range list {
		s = s + v
	}
	return s
}

func SumFloat32(list []float32) float32 {
	var s float32
	for _, v := range list {
		s = s + v
	}
	return s
}

func SumUint(list []uint) uint {
	var s uint
	for _, v := range list {
		s = s + v
	}
	return s
}

func UniqueUint(list *[]uint) {
	var c uint
	check := make(map[uint]uint)
	for _, val := range *list {
		check[val] = 1
	}
	res := make([]uint, len(check))
	for num, _ := range check {
		res[c] = num
		c++
	}
	*list = res
}

func RangeUint(min uint, max uint) []uint {
	var i uint
	rangeslice := make([]uint, max-min+1)
	for i = 0; i < uint(len(rangeslice)); i++ {
		rangeslice[i] = min + (i)
	}
	return rangeslice
}

func ShuffleUint(slice []uint, seed int64) {
	rand.Seed(seed)
	rand.Shuffle(len(slice), func(i, j int) { (slice)[i], (slice)[j] = (slice)[j], (slice)[i] })
}
