package lesson3

import (
	"sync"
)


// A -> B in the same time B -> A

// перша заблокує A, друга B

type Account struct{
	mutex sync.Mutex
	sum int
	characteristic int
}


//func transferMoney(a, b Account, sum int){
//	if a.characteristic < b.characteristic{
//		a.Lock()
//		b.Lock()
//	}else{
//		b.Lock()
//		a.Lock()
//	}
//
//	a-=sum
//	b-=sum
//
//	if a.characteristic < b.characteristic{
//		b.Unlock()
//		a.Unlock()
//	}else{
//		a.Unlock()
//		b.Unlock()
//	}
//
//}
