// Пакет для прикладной обработки сигналов
package mysign

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
)

func SoftShotDown(signList []os.Signal, writer *os.File, msg string) {
	// Функция принимает список сигналов, файл для записи и данные для записи.
	// При срабатывании любого из сигналов из списка, данные записываются в файл спустя определенное время (timeBeforeSD)
	// При истечении определенного времени (timeLimit) будет вызван таймаут контекста
	// При всех сценариях будет записано в файл соответствующее своему сценирию сообщение.
	var (
		timeLimit    = 10 * time.Second
		timeBeforeSD = time.Second
		signCh       = make(chan os.Signal, 1)
		ctx, cancel  = context.WithTimeout(context.Background(), timeLimit)

		print = func(message string) {
			if _, err := fmt.Fprintf(writer, "\n%s\n", message); err != nil {
				log.Println(message)
			}
		}
	)

	signal.Notify(signCh, signList...)

	defer cancel()
	print(fmt.Sprintf("Press 'Ctrl+C' for cancel program of it will stop after %d sec", timeLimit/10e8))
	select {
	case <-time.After(2 * timeLimit):
		print("Over time. Never printing")

	case <-signCh:
		print("program stopping...")
		time.Sleep(timeBeforeSD)
		print(msg)

	case <-ctx.Done():
		print("stop from context timeout")
	}
}
