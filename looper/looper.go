package looper

import (
	"reflect"
)

type Looper struct {
	List []interface{}
	Index int
	Len int
}

type LooperStr struct {
	List []string
	Index int
	Len int
}

type LooperInt struct {
	List []int
	Index int
	Len int
}

func NewInt(list []int) LooperInt {
	return LooperInt{List: list, Len:len(list)}
}

func NewStr(list []string) LooperStr {
	return LooperStr{List: list, Len:len(list)}
}

func New(list interface{}) Looper {
	s := reflect.ValueOf(list)
	r := make([]interface{}, s.Len())
	for i:=0; i<s.Len(); i++ {
        r[i] = s.Index(i).Interface()
    }
	return Looper{List: r, Len: s.Len()}
}

func (l *Looper) Next() (j interface{}) {
	j = l.List[l.Index]
	l.Index++
	if l.Index + 1 > l.Len {
		l.Index = 0
	}
	return
}

func (l *LooperStr) Next() (j string) {
	j = l.List[l.Index]
	l.Index++
	if l.Index + 1 > l.Len {
		l.Index = 0
	}
	return
}

func (l *LooperInt) Next() (j int) {
	j = l.List[l.Index]
	l.Index++
	if l.Index + 1 > l.Len {
		l.Index = 0
	}
	return
}

