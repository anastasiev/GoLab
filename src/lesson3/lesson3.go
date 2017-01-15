package lesson3


//Обідаючі філософи
// є стіл навколо якого сидять 5 філософів, їм треба їсти
// між кожною парою філософів по 1 виделці, щоб поїсти йому треба взяти 2 виделки



import(
	"fmt"
	"sync"
	"math/rand"
	"time"
	"sync/atomic"
)
const N = 5 // кількість філософів

type philState int

const (
	Thinking philState = iota
	Hungry
	Eating
)

type philosopher struct {
	canEat *sync.Cond
	state philState
	num int
}

var(
	fork[N] sync.Mutex
	a int64

	philosophers [N] *philosopher
	m sync.Mutex
)

func New(i int) *philosopher  {
	p := new(philosopher)
	p.num = i
	p.state = Thinking
	p.canEat = sync.NewCond(&m)
	return p
}

func (p *philosopher) Think(){
	fmt.Printf("Philosopher #%d thinking\n", p.num);
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
}

func (p *philosopher) Eat(){
	fmt.Printf("Philosopher #%d eating\n", p.num);
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
}

func (p *philosopher) TakeForks(){
	m.Lock()
	p.state = Hungry
	for p.Left().state == Eating || p.Right().state == Eating{
		p.canEat.Wait()
	}
	p.state = Eating
	m.Unlock()
}

func (p *philosopher) PutForks(){
	m.Lock()
	p.state = Thinking
	if p.Left().state == Hungry && p.Left().Left().state != Eating{
		p.Left().canEat.Signal()
	}
	if p.Right().state == Hungry && p.Right().Right().state != Eating{
		p.Right().canEat.Signal()
	}
	m.Unlock()
}

func (p *philosopher) Left() *philosopher{
	return philosophers[(N + p.num - 1) % N]
}

func (p *philosopher) Right() *philosopher{
	return philosophers[(p.num + 1) % N]
}



// i номер філософа
// повертає номер виделки
// номер виделки зліва == номеру філософа
func left(i int) int{
	return i;
}
func right(i int) int{
	return (i+1)%N;
}

//так буде дедлок
//func philosopher(i int){
//	for {
//		fmt.Printf("Philosopher #%d thinking\n", i);
//		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
//		fork[left(i)].Lock()
//		fork[right(i)].Lock()
//		fmt.Printf("Philosopher #%d eating\n", i);
//		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
//		fork[right(i)].Unlock()
//		fork[left(i)].Unlock()
//	}
//}

//без дедлоку, але повільно
func philosopher2(i int, c chan bool){

	for {
		atomic.AddInt64(&a, 1)
		c <- true

		fmt.Printf("Philosopher #%d thinking\n", i);
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		fork[left(i)].Lock()
		fork[right(i)].Lock()
		fmt.Printf("Philosopher #%d eating\n", i);
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		fork[right(i)].Unlock()
		fork[left(i)].Unlock()

		<- c
	}
}

func philosopher3(i int){
	for {
		atomic.AddInt64(&a, 1)
		fmt.Printf("Philosopher #%d thinking\n", i);
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		fork[min(left(i), right(i))].Lock()
		fork[max(left(i), right(i))].Lock()
		fmt.Printf("Philosopher #%d eating\n", i);
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		fork[max(left(i), right(i))].Unlock()
		fork[min(left(i), right(i))].Unlock()
	}
}

func philosopher4(i int){
	for {
		atomic.AddInt64(&a, 1)

		fmt.Printf("Philosopher #%d thinking\n", i);
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		fork[min(left(i), right(i))].Lock()
		fork[max(left(i), right(i))].Lock()
		fmt.Printf("Philosopher #%d eating\n", i);
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		fork[max(left(i), right(i))].Unlock()
		fork[min(left(i), right(i))].Unlock()
	}
}

func (p *philosopher) Dine()  {
	for{
		atomic.AddInt64(&a, 1)
		p.Think()
		p.TakeForks()
		p.Eat()
		p.PutForks()
	}
}

func Lesson3(){
	for i:= 0; i<N; i++{
		philosophers[i] = New(i)
	}
	for _, p := range philosophers{
		go p.Dine()
	}

}

func min(x, y int)  int{
	if x < y{
		return x
	}
	return y
}

func max(x, y int)  int{
	if x > y{
		return x
	}
	return y
}

func PrintA()  {
	fmt.Println(a)
}



