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

func IndexStr(list []string, str string) int {
	for i, v := range list {
		if v == str {
			return i
		}
	}
	return -1
}

func ContainStr(list []string, str string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

func UniqueStr(list *[]string) {
	var c uint
	check := make(map[string]uint)
	for _, val := range *list {
		check[val] = 1
	}
	res := make([]string, len(check))
	for letter, _ := range check {
		res[c] = letter
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

func FilterStr(list *[]string, filter ...string) {
	var c int
	var add bool
	s := make([]string, len(*list)-CountStr(*list, filter...))
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

func ShuffleStr(slice []string, seed int64) {
	rand.Seed(seed)
	rand.Shuffle(len(slice), func(i, j int) { (slice)[i], (slice)[j] = (slice)[j], (slice)[i] })
}

func ReverseStr(slice []string) []string {
	var c int
	slicelen := len(slice)
	s := make([]string, slicelen)
	for i := slicelen - 1; i >= 0; i-- {
		s[c] = slice[i]
		c++
	}
	return s
}

func PopStr(slice *[]string, indexes ...int) {
	var add bool
	var c int
	s := make([]string, len(*slice)-len(indexes))
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
