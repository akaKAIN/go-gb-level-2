package myatomic

import (
	"fmt"
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
	for i:=0; i<quantity; i++ {
		go handler(&wg)
	}
	wg.Wait()
	fmt.Println("All goroutine was done")
}
