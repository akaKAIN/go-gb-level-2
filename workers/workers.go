// Пакет для работы с паттерном "Пул-воркеров"
package workers

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
