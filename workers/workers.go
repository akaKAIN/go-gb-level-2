package workers

func WorkerHandler(workerQuantity int, handler func()) {
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
	ch <- struct{}{}
}
