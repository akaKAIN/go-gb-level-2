package getter

import (
	log "github.com/sirupsen/logrus"
	_ "github.com/valyala/fasthttp@v1.21.0"
	"net/http"
)

func GetStatusCode(url string) int {
	resp, err := http.Get(url)

	if err != nil {
		log.Warn(err)
		return -1
	}
	return resp.StatusCode
}

