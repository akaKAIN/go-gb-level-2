package myatomic

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

func SimpleHandler(wg *sync.WaitGroup) {
	// Простая функция с ожиданием. По окончанию в wg идет оповещение и завершении работы функции
	time.Sleep(time.Second)
	defer wg.Done()
}

func StartGo(quantity int, handler func(wg *sync.WaitGroup)) {
	// Функция зупускает заданное количество горутин и ожидает их завершения
	var wg sync.WaitGroup
	wg.Add(quantity)
	for i := 0; i < quantity; i++ {
		go handler(&wg)
	}
	wg.Wait()
	fmt.Println("All goroutine was done")
}

func StartReadAndWrite(writeQuantity, readQuantity int) {

	var (
		store = IntMap{
			lock: sync.RWMutex{},
			Map:  make(map[int]int),
		}
		ctx, finish = context.WithCancel(context.Background())
	)

	for i := 1; i <= writeQuantity+readQuantity; i++ {
		_ = store.Add(i, i)
	}

	go ReadAndWrite(writeQuantity, readQuantity, &store, finish)

	for {
		select {
		case <-ctx.Done():
			return
		}
	}
}

func ReadAndWrite(writeQuantity, readQuantity int, store Buffer, finish context.CancelFunc) {
	var (
		wg sync.WaitGroup
	)
	wg.Add(writeQuantity + readQuantity)
	go func() {
		for w := 0; w < writeQuantity; w++ {
			go func(i int) {
				defer wg.Done()
				if err := store.Add(i, i); err != nil {
					log.Printf("%s for %d\n", err, i)
				}
			}(w)
		}
	}()

	go func() {
		for r := 0; r < readQuantity; r++ {
			go func(i int) {
				defer wg.Done()
				if _, err := store.Get(r); err != nil {
					log.Printf("%s for %d\n", err, i)
				}
			}(r)
		}
	}()
	wg.Wait()
	finish()
}
