package services

import (
	"fmt"

	"github.com/Levap123/go-gobuster/internal/usecase"
)

func (s *Services) Bust(url string, urlChan chan string) {
	for _, route := range s.routes {
		path := fmt.Sprintf("%s/%s", url, route)
		resp, err := usecase.PerfromGet(path)
		if err != nil {
			continue
		}

		if resp.StatusCode == 404 {
			continue
		}

		if resp.StatusCode == 302 {
			continue
		}

		if resp.StatusCode == 429 {
			continue
		}

		urlChan <- path + " " + resp.Status
	}

}
