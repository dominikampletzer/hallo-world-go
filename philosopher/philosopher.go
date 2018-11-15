package philosopher

import (
	"fmt"
	"time"
)

type Table struct {
	reservedChannels  []chan bool
	putForkChannel    chan int
	takeChannel       chan int
	forks             []bool
	philosophersCount int
}

func NewTable(i int) *Table {
	table := new(Table)

	table.putForkChannel = make(chan int)
	table.takeChannel = make(chan int)
	table.reservedChannels = make([]chan bool, i)
	table.philosophersCount = i
	for j := 0; j < i; j++ {
		table.reservedChannels[j] = make(chan bool)
	}

	table.forks = make([]bool, i)

	return table
}

func (t Table) getReservedChannel(id int) chan bool {
	return t.reservedChannels[id]
}
func (t Table) GetPutForkChannel() chan int {
	return t.putForkChannel
}

func (t Table) GetTakeChannel() chan int {
	return t.takeChannel
}

func (t Table) Run() {
	for {
		select {
		case p := <-t.putForkChannel:
			{
				t.forks[p], t.forks[(p+1)%t.philosophersCount] = false, false
			}
		case forkNo := <-t.takeChannel:
			{
				if !t.forks[forkNo] && !t.forks[(forkNo+1)%t.philosophersCount] {
					t.forks[forkNo], t.forks[(forkNo+1)&t.philosophersCount] = true, true
					t.reservedChannels[forkNo] <- true
				} else {
					t.reservedChannels[forkNo] <- false
				}
			}
		}
	}
}

type Philosopher struct {
	id    int
	table *Table
}

func (p *Philosopher) think() {
	fmt.Printf("[->]: Philosopher #%d thinks ...\n", p.id)
	time.Sleep(2 * time.Millisecond)
	fmt.Printf("[<-]: Philosopher #%d thinking ends\n", p.id)
}

func (p *Philosopher) eat() {
	fmt.Printf("[->]: Philosopher #%d thinks ...\n", p.id)
	time.Sleep(3 * time.Millisecond)
	fmt.Printf("[<-]: Philosopher #%d thinking ends\n", p.id)
}

func (p *Philosopher) takeForks() {

	gotForks := false
	for !gotForks {
		c := p.table.GetTakeChannel()
		c <- p.id
		gotForks = <-p.table.getReservedChannel(p.id)
	}
}

func (p *Philosopher) putForks() {
	p.table.GetPutForkChannel() <- p.id
}

func (p *Philosopher) run() {
	for {
		p.takeForks()
		p.eat()
		p.putForks()
		p.think()
	}
}
func NewPhilosopher(name int, table *Table) *Philosopher {
	p := new(Philosopher)
	p.id = name
	p.table = table
	return p
}
