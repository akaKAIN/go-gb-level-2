package workers

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func Start() {
	rand.Seed(time.Now().UnixNano())
	var (
		ctx, cancel = context.WithCancel(context.Background())
		hr          = func(cancel context.CancelFunc) {
			for t := 1; t <= 4; t++ {
				time.Sleep(time.Second)
			}
			cancel()
			fmt.Println("hr goes home")
		}
		jobsChan    = make(chan int)
		managerFunc = func(ctx context.Context) {
			for job := 1; ; job++ {
				select {
				case <-ctx.Done():
					close(jobsChan)
					fmt.Println("manager goes home")
					return
				default:
					fmt.Printf("manager create job %d\n", job)
					jobsChan <- job
				}
			}
		}
		resource = make(chan struct{}, 3)
		worker   = func(id int) {
			defer func() { <-resource }()
			for job := range jobsChan {
				fmt.Printf("worker %d starts processing of %d\n", id, job)
				<-time.NewTicker(6 * time.Second).C
				fmt.Printf("worker %d completes processing of %d\n", id, job)
			}
			fmt.Printf("worker %d goes home\n", id)
		}
	)
	go managerFunc(ctx)
	go hr(cancel)

	for i := 1; i <= cap(resource); i++ {
		resource <- struct{}{}
		go worker(i)
	}

	select {
	case <-ctx.Done():
		for i := 1; i <= cap(resource); i++ {
			resource <- struct{}{}
		}
		close(resource)
		return
	}
}
