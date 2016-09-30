package lesson


var a int
var C chan int

func F (){
	b := 0
	for i:=0; i<100000; i++{
		b++
	}
	C <- b
}
func Add() int{
	for i:=0; i<10; i++{
		go F()
	}
	for i:=0; i<10;  i++{
		a += <-C
	}

	return a
}
