package slices

import (
	"math/rand"
)

func SumUint(list []uint) uint {
	var s uint
	for _, v := range list {
		s = s + v
	}
	return s
}

func CountUint(list []uint, num ...uint) int {
	var c int
	for _, val := range list {
		for _, n := range num {
			if val == n {
				c++
				break
			}
		}
	}
	return c
}

func FilterUint(list *[]uint, filter ...uint) {
	var c int
	var add bool
	s := make([]uint, len(*list)-CountUint(*list, filter...))
	for _, val := range *list {
		add = true
		for _, f := range filter {
			if val == f {
				add = false
				break
			}
		}
		if add {
			s[c] = val
			c++
		}
	}
	*list = s
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

func PopUint(slice *[]uint, indexes ...int) {
	var add bool
	var c int
	s := make([]uint, len(*slice)-len(indexes))
	for i, v := range *slice {
		add = true
		for _, index := range indexes {
			if i == index {
				add = false
				break
			}
		}
		if add {
			s[c] = v
			c++
		}
	}
	*slice = s
}
