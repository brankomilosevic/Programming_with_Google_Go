package main

import (
	"fmt"
	"sync"
	"time"
)

type ChopStick struct {
	sync.Mutex
}

type Philosopher struct {
	num, count        int
	left_cs, right_cs *ChopStick
}

func (p Philosopher) eat(c chan *Philosopher, wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		c <- &p

		if p.count < 3 {
			p.left_cs.Lock()
			p.right_cs.Lock()

			fmt.Println("starting to eat : ", p.num)
			p.count = p.count + 1
			fmt.Println("finishing eating: ", p.num)
			p.right_cs.Unlock()
			p.left_cs.Unlock()

			wg.Done()
		}

	}
}

func host(c chan *Philosopher, wg *sync.WaitGroup) {
	for {
		if len(c) == 2 {
			<-c
			<-c
			//time delay
			time.Sleep(20 * time.Millisecond)
		}
	}
}

func main() {
	var i int
	var wg sync.WaitGroup

	c := make(chan *Philosopher, 2)

	wg.Add(15)

	ChopSticks := make([]*ChopStick, 5)
	for i = 0; i < 5; i++ {
		ChopSticks[i] = new(ChopStick)
	}

	Philosophers := make([]*Philosopher, 5)
	for i = 0; i < 5; i++ {
		Philosophers[i] = &Philosopher{i + 1, 0, ChopSticks[i], ChopSticks[(i+1)%5]}
	}

	go host(c, &wg)

	for i = 0; i < 5; i++ {
		go Philosophers[i].eat(c, &wg)
	}

	wg.Wait()
}
