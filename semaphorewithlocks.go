package semaphore

import (
	"fmt"
	"sync"
	"time"
)

type Sema struct {
	cond  *sync.Cond
	count int
}

func (s *Sema) acquire(th int) {
	s.cond.L.Lock()
	if s.count == 0 {
		fmt.Println("lease not available, waiting", th)
		s.cond.Wait()
		fmt.Println("woke up from waiting", th)
	}
	fmt.Println("acquired lease", th)
	s.count--
	s.cond.L.Unlock()
}

func (s *Sema) release(th int) {
	s.cond.L.Lock()
	fmt.Println("release lease", th)
	s.count++
	s.cond.Signal()
	s.cond.L.Unlock()
}

func printHello(s *Sema, th int) {
	s.acquire(th)
	time.Sleep(10 * time.Millisecond)
	fmt.Println("I am in critical section", th, time.Now().Unix())
	s.release(th)
}
