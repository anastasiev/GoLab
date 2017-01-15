package lesson2

import (
	"sync"
	"math/rand"
	"fmt"
	"time"
)

const N  = 10
var(
	P = 10
	C = 10
	buf [N]int
	head, tail, count int
	m sync.Mutex
)

func Lesson2()  {

	notFull := sync.NewCond(&m)
	notEmpty := sync.NewCond(&m)
	for i:=0; i<P; i++{
		go producer(i, notFull, notEmpty)
	}
	for i:=0; i<C; i++{
		go consumer(i, notFull, notEmpty)
	}
	time.Sleep(5 * time.Second)
}

func producer(i int, notFull, notEmpty *sync.Cond){
	for {
		m.Lock()
		for count == N {
			notFull.Wait() // мьютекс відпоскається (реалізація на сигнальних змінних)
		}
		count++
		val := rand.Intn(10)
		fmt.Printf("Producer #%d:%d\n",i,val)
		buf[head] = val

		head = (head + 1) % N
		notEmpty.Signal()
		m.Unlock()
		time.Sleep(500 * time.Millisecond)
	}
}
func consumer(i int, notFull, notEmpty *sync.Cond){
	for {
		m.Lock()
		for count == 0 {
			notEmpty.Wait() // мьютекс відпоскається (реалізація на сигнальних змінних)
		}
		count--
		val := buf[tail]
		fmt.Printf("Consumer #%d:%d\n",i,val)
		buf[head] = val

		tail = (tail + 1) % N
		notFull.Signal()
		m.Unlock()
		time.Sleep(600 * time.Millisecond)
	}
}


