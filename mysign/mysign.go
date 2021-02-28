package mysign

import (
	"fmt"
	"log"
	"os"
	"os/signal"
)

func SoftShotDown(signList []os.Signal, writer *os.File, msg string) {
	signCh := make(chan os.Signal, 1)
	signal.Notify(signCh, signList...)
	for {
		select {
		case <-signCh:
			if _, err := fmt.Fprintf(writer, "%s", msg); err != nil {
				log.Println(err)
			}
			os.Exit(0)
		}
	}
}
