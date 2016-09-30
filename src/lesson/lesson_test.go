package lesson

import "testing"

func TestAdd(t *testing.T) {
	C = make(chan int)
	if Add() != 1000000{
		t.Error("Failed")
	}
}
func BenchmarkAdd(b *testing.B) {
	C = make(chan int)
	for i:= 0; i < b.N; i++{
		Add()
	}
}