package main

import (
	"fmt"
	"github.com/akaKAIN/go-gb-level-2/workers"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// 1. С помощью пула воркеров написать программу, которая запускает 1000 горутин,
// каждая из которых увеличивает число на 1. Дождаться завершения всех горутин
// и убедиться, что при каждом запуске программы итоговое число равно 1000.

// 2. Написать программу, которая при получении в канал сигнала SIGTERM
// останавливается не позднее, чем за одну секунду (установить таймаут).
func main() {
	var counter int
	jobHandler := func() {counter++}
	workers.WorkerHandler(1000, jobHandler)
	fmt.Printf("Counter: %d\n", counter)

	//workers.Start()
	//mysign.SoftShotDown([]os.Signal{syscall.SIGINT}, os.Stdout, "well done")

	channelBuffer := 100
	ch := make(chan int, channelBuffer)
	for i:=1; i<=channelBuffer; i++ {
		ch <- i
	}
	workers.SleepingPool(3, ch, time.Second)
}
