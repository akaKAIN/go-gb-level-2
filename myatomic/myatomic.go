// Тут приведены практические примеры из методичек и издевательства над ними
package myatomic

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

const count = 1000

func Increment() {
	var (
		counter int
		mutex   sync.Mutex
		wg      sync.WaitGroup

		// Вспомогательная часть нашего кода
		ch = make(chan struct{}, count)
	)
	wg.Add(count)
	for i := 0; i < count; i += 1 {
		go func() {
			defer wg.Done()
			// Захват мьютекса
			mutex.Lock()
			defer mutex.Unlock()

			counter += 1

			// Фиксация факта запуска горутины в канале
			ch <- struct{}{}

		}()
	}
	wg.Wait()
	close(ch)

	i := 0
	for range ch {
		i += 1
	}
	// Выводим показание счетчика
	fmt.Println(counter)
	// Выводим показания канала
	fmt.Println(i)
}

func SimpleCounter() {
	var (
		counter int64 = 0
		limit         = 1000
		m       sync.Mutex
	)
	ch := make(chan struct{}, limit)

	for i := 0; i < limit; i++ {
		ch <- struct{}{}
		go func() {
			m.Lock()
			defer m.Unlock()
			counter++
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(counter, "/", len(ch))
}

func CompareAndSwap(num *int) {
	var a = int64(2)
	atomic.CompareAndSwapInt64(&a, a, 1)
	fmt.Println(a)
}

func Swap() {
	var a = int64(2)
	r := atomic.SwapInt64(&a, 1)
	fmt.Println(r, a)
}

type Set struct {
	sync.RWMutex
	mm map[int]struct{}
}

func NewSet() *Set {
	return &Set{
		mm: make(map[int]struct{}),
	}
}

func (s *Set) Add(i int) {
	s.Lock()
	s.mm[i] = struct{}{}
	s.Unlock()
}

func (s *Set) Has(i int) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.mm[i]
	return ok
}
