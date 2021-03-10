package myarray

import (
	"fmt"
	"log"
	"os"
	"runtime/trace"
	"sync"
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

	err := arr.ReplaceSubString("H", "0")
	if err != nil {
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
