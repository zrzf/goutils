package file

import (
	// "os"
	"strings"
	"io/ioutil"
)

func Read(filename string) string {
	data, _ := ioutil.ReadFile(filename)
	return string(data)
}

func ReadLines(filename string) []string {
	return strings.Split(Read(filename), "\n")
}
