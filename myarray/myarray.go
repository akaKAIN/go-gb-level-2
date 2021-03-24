package myarray

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
	"time"
)

func InitByteArray(word string) *ByteArray {
	return &ByteArray{
		lock:  sync.Mutex{},
		Array: []byte(word),
	}
}

func TaskOne(quantity int) {
	if err := trace.Start(os.Stderr); err != nil {
		log.Println(err)
		return
	}
	defer trace.Stop()

	var (
		wg  sync.WaitGroup
		arr = InitByteArray("Hello World")
	)

	if err := arr.ReplaceSubString("H", "0"); err != nil {
		log.Println(err)
		return
	}

	wg.Add(quantity)
	for w := 0; w < quantity; w += 1 {
		go func(ind int) {
			defer wg.Done()
			if _, err := arr.GetByteByInd(ind); err != nil {
				log.Println(err)
			}
		}(w)
	}
	wg.Wait()
	fmt.Println("end")
}

func TaskTwo(limit int) {
	if err := trace.Start(os.Stderr); err != nil {
		log.Println(err)
		return
	}
	defer trace.Stop()

	counter := 0
	for {
		if counter%limit == 0 {
			break
		}
		go func(i int) {
			time.Sleep(20 * time.Millisecond)
			if i%10e3 == 0 {
				runtime.Gosched()
			}

		}(counter)
		counter += 1
	}
}

func TaskThree() {
	// Run 10e3 goroutines with data race condition.
	var (
		store    int = 0
		quantity int = 10e3
		wg       sync.WaitGroup
	)

	wg.Add(quantity)
	for i := 0; i < quantity; i += 1 {
		func() {
			defer wg.Done()
			go RaceIncrement(&store)
		}()
	}

	wg.Wait()
	fmt.Printf("Done: %d/%d\n", store, quantity)
}

func RaceIncrement(targetNum *int) {
	// func with data race
	*targetNum += 1
}
