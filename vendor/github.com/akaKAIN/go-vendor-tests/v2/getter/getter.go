package getter

import (
	_ "github.com/valyala/fasthttp@v1.21.0"
	"net/http"
)

func GetStatusCode(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	return resp.StatusCode, nil
}

