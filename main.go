package main

import (
	"fmt"
	"sync"
	"time"
)

type Philosopher struct {
	Name string
	Left int
	Right int
}

func (p Philosopher) Eat(wg *sync.WaitGroup, table *Table) {
	defer wg.Done()
	table.Forks[p.Left].Mux.Lock()
	defer table.Forks[p.Left].Mux.Unlock()

	time.Sleep(150 * time.Millisecond)

	table.Forks[p.Right].Mux.Lock()
	defer table.Forks[p.Right].Mux.Unlock()

	fmt.Println(fmt.Sprintf("%v is eating", p.Name))

	time.Sleep(1000 * time.Millisecond)

	fmt.Println(fmt.Sprintf("%v is done eating", p.Name))
}

type Table struct {
	Forks []Fork
}

type Fork struct {
	Mux sync.Mutex
}

func main() {
	var wg sync.WaitGroup
	var forks = []Fork{
		Fork{},
		Fork{},
		Fork{},
		Fork{},
		Fork{},
	}
	var table = Table{
		forks,
	}
	var philosophers = []Philosopher{
		Philosopher{
			"Judith Butler", 0, 1,
		},
		Philosopher{
			"Gilles Deleuze", 1, 2,
		},
		Philosopher{
			"Karl Marx", 2, 3,
		},
		Philosopher{
			"Emma Goldman", 3, 4,
		},
		Philosopher{
			"Michel Foucault", 0, 4,
		},
	}

	for _, philosopher := range philosophers {
		wg.Add(1)
		go philosopher.Eat(&wg, &table)
	}
	wg.Wait()
}
