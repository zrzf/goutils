package slices

import (
	"math/rand"
)

func ChunkStr(list []string, chunksize int) [][]string {
	lenlist := len(list)
	var chunks [][]string
	for i := 0; i < lenlist; i += chunksize {
		end := i + chunksize
		if end > lenlist {
			end = lenlist
		}
		chunks = append(chunks, list[i:end])
	}
	return chunks
}

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

func IndexStr(list []string, str string) int {
	for i, v := range list {
		if v == str {
			return i
		}
	}
	return -1
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

func ContainStr(list []string, str string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

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

func SumInt(list []int) int {
	var s int
	for _, v := range list {
		s = s + v
	}
	return s
}

func UniqueStr(list *[]string) {
	var c uint
	check := make(map[string]uint)
	for _, val := range *list {
		check[val] = 1
	}
	res := make([]string,len(check))
	for letter, _ := range check {
		res[c] = letter
		c++
	}
	*list = res
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

func CountStr(list []string, str ...string) int {
	var c int
	for _, val := range list {
		for _, s := range str {
			if val == s {
				c++
				break
			}
		}
	}
	return c
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
	s := make([]int, len(*list) - CountInt(*list, filter...))
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

func FilterStr(list *[]string, filter ...string) {
	var c int
	var add bool
	s := make([]string, len(*list) - CountStr(*list, filter...))
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
	rangeslice := make([]int, max - min + 1)
	for i=0;i<len(rangeslice);i++{
		rangeslice[i] = min + i
	}
	return rangeslice
}

func RangeUint(min uint, max uint) []uint {
	var i uint
	rangeslice := make([]uint, max - min + 1)
	for i=0;i<uint(len(rangeslice));i++{
		rangeslice[i] = min + (i)
	}
	return rangeslice
}

func ShuffleInt(slice []int, seed int64) {
	rand.Seed(seed)
	rand.Shuffle(len(slice), func(i, j int) { (slice)[i], (slice)[j] = (slice)[j], (slice)[i] })
}

func ShuffleUint(slice []uint, seed int64) {
	rand.Seed(seed)
	rand.Shuffle(len(slice), func(i, j int) { (slice)[i], (slice)[j] = (slice)[j], (slice)[i] })
}

func ShuffleStr(slice []string, seed int64) {
	rand.Seed(seed)
	rand.Shuffle(len(slice), func(i, j int) { (slice)[i], (slice)[j] = (slice)[j], (slice)[i] })
}

func ReverseStr(slice []string) []string {
	var c int
	slicelen := len(slice)
	s := make([]string, slicelen)
	for i:=slicelen-1;i>=0;i--{
		s[c] = slice[i]
		c++
	}
	return s
}

func ReverseInt(slice []int) []int {
	var c int
	slicelen := len(slice)
	s := make([]int, slicelen)
	for i:=slicelen-1;i>=0;i--{
		s[c] = slice[i]
		c++
	}
	return s
}
