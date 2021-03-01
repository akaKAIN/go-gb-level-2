// Пакет для работы с паттерном "Пул-воркеров"
package workers

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"
)

type contextKey string

func (c contextKey) String() string {
	return fmt.Sprintf("%v", c)
}

func WorkerHandler(workerQuantity int, handler func()) {
	// Запускаются воркеры в указанном количестве. Каждый выполняет свою работу (handler) в переданный ему канал и
	// заверщает свою работу. Когда количество выполненных воркеров достигнет лимита, канал закроется и произойдет
	// выход из функции
	jobCh := make(chan struct{})
	limit := 0

	for w := 0; w < workerQuantity; w++ {
		go Worker(jobCh)
	}
	for {
		select {
		case <-jobCh:
			handler()
			limit++
			if limit == workerQuantity {
				close(jobCh)
				return
			}
		}
	}
}

func Worker(ch chan<- struct{}) {
	//Простейший воркер
	ch <- struct{}{}
}

func SleepingPool(workerQuantity int, ch <-chan int, sleepTime time.Duration) {
	corpsCh := make(chan struct{})
	ctx := context.Background()
	for w := 1; w <= workerQuantity; w++ {
		val := fmt.Sprintf("%d", w)
		ctx := context.WithValue(ctx, "name", val)
		go SleepingWorker(ctx, ch, sleepTime, corpsCh)
	}

	// Синхронизация с работающими воркерами
	for workerQuantity > 0 {
		select {
		case <-corpsCh:
			workerQuantity--
		}
	}
	fmt.Println("All workers is dead.")
}

func SleepingWorker(ctx context.Context, ch <-chan int, sleepDuration time.Duration, corpsCh chan<- struct{}) {
	workerName := ctx.Value("name")

	t := time.Duration(rand.Intn(10) + 4) * time.Second
	fmt.Println("timer", t)

	ctx, finish := context.WithTimeout(ctx, time.Duration(rand.Intn(6) + 6) * time.Second)

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %s is dead\n", workerName)
			corpsCh <- struct{}{}
			finish()
			return
		case work := <- ch:
			log.Printf("worker %s start job %d\n", workerName, work)
			time.Sleep(sleepDuration)
			log.Printf("worker %s finish job %d\n", workerName, work)
		default:
			finish()
		}
	}
}
