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
	var (
		timeLimit   = 10 * time.Second
		signCh      = make(chan os.Signal, 1)
		ctx, cancel = context.WithTimeout(context.Background(), timeLimit)

		print = func(message string) {
			if _, err := fmt.Fprintf(writer, "\n%s\n", message); err != nil {
				log.Println(message)
			}
		}
	)

	signal.Notify(signCh, signList...)

	defer cancel()
	print(fmt.Sprintf("Press 'Ctrl+C' for cancel program of it will stop after %d sec", time.Second / 10e8))
	select {
	case <-time.After(2 * timeLimit):
		print("Over time. Never printing")

	case <-signCh:
		time.Sleep(time.Second)
		print(msg)

	case <-ctx.Done():
		print("stop from context timeout")
	}
}
