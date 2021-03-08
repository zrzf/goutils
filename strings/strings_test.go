package strings

import (
	"testing"
	"github.com/zrzf/goutils/semaphore"
)

func BenchmarkRemoveBetweenRune(b *testing.B){
	for i:=0;i<b.N;i++{
		RemoveBetweenRune("<abc>abc<abc>",'<','>', "")
	}
}

func BenchmarkSemaphore(b *testing.B){
	s := semaphore.New(100)
	for i:=0;i<b.N;i++{
		s.Add()
		s.Done()
	}
	s.Wait()
}
