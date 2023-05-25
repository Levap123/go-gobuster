package usecase

import (
	"os"
	"strings"
)

func ReadTxt(path string) ([]string, error) {
	body, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	routes := strings.Split(string(body), "\n")

	return routes, nil
}
