package slices

import (
	"math/rand"
	"sort"
)

func ChunkInt(list []int, chunksize int) [][]int {
	lenlist := len(list)
	var chunks [][]int
	for i := 0; i < lenlist; i += chunksize {
		end := i + chunksize
		if end > lenlist {
			end = lenlist
		}
		chunks = append(chunks, list[i:end])
	}
	return chunks
}

func IndexInt(list []int, num int) int {
	for i, v := range list {
		if v == num {
			return i
		}
	}
	return -1
}

func ContainInt(list []int, num int) bool {
	for _, v := range list {
		if v == num {
			return true
		}
	}
	return false
}

func SumInt(list []int) int {
	var s int
	for _, v := range list {
		s = s + v
	}
	return s
}

func UniqueInt(list *[]int) {
	var c uint
	check := make(map[int]uint)
	for _, val := range *list {
		check[val] = 1
	}
	res := make([]int, len(check))
	for num, _ := range check {
		res[c] = num
		c++
	}
	*list = res
}

func CountInt(list []int, num ...int) int {
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

func FilterInt(list *[]int, filter ...int) {
	var c int
	var add bool
	s := make([]int, len(*list)-CountInt(*list, filter...))
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

func RangeInt(min int, max int) []int {
	var i int
	rangeslice := make([]int, max-min+1)
	for i = 0; i < len(rangeslice); i++ {
		rangeslice[i] = min + i
	}
	return rangeslice
}

func ShuffleInt(slice []int, seed int64) {
	rand.Seed(seed)
	rand.Shuffle(len(slice), func(i, j int) { (slice)[i], (slice)[j] = (slice)[j], (slice)[i] })
}

func ReverseInt(slice []int) []int {
	var c int
	slicelen := len(slice)
	s := make([]int, slicelen)
	for i := slicelen - 1; i >= 0; i-- {
		s[c] = slice[i]
		c++
	}
	return s
}

func PopInt(slice *[]int, indexes ...int) {
	var c, ci int
	var compare bool = true
	sort.Ints(indexes)
	s := make([]int, len(*slice)-len(indexes))
	leni := len(indexes)
	for i, v := range *slice {
		if compare && i == indexes[ci] {
			ci++
			if ci == leni {
				compare = false
			}
		} else {
			s[c] = v
			c++
		}
	}
	*slice = s
}
