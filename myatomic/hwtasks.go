package myatomic

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

func SimpleWorker(ch chan<- struct{}) {
	// Простая функция с ожиданием. По окончании работы, сообщает в канал.
	time.Sleep(time.Second)
	ch <- struct{}{}
}

func StartGo(quantity int, worker func(ch chan<- struct{})) {
	// Функция запускает заданное количество горутин и ожидает их завершения
	var (
		wg sync.WaitGroup
		ch = make(chan struct{})
	)

	wg.Add(quantity)
	for i := 0; i < quantity; i++ {
		go worker(ch)
	}

	go func() {
		for {
			_, ok := <-ch
			if !ok {
				break
			}
			wg.Done()
		}
	}()

	wg.Wait()
	close(ch)
	fmt.Printf("%d goroutins was done\n", quantity)
}

func StartReadAndWrite(writeQuantity, readQuantity int, store Buffer) {
	ctx, finish := context.WithCancel(context.Background())
	go func() {
		defer finish()
		ReadAndWrite(writeQuantity, readQuantity, store)
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Second):
			fmt.Println("Timeout")
			return
		}
	}
}

func ReadAndWrite(writeQuantity, readQuantity int, store Buffer) {
	var (
		wg sync.WaitGroup
	)
	wg.Add(writeQuantity)
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

	wg.Add(readQuantity)
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
}
