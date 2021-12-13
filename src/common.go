package main

import (
	"errors"
	"strconv"
	"strings"
)

func extractId(url string) (int, error) {
	list := strings.Split(url, "/")
	if len(list) > 3 {
		id, err := strconv.Atoi(list[3])
		if err!=nil {
			return 0,nil
		}
		return id, nil
	}
	return 0, errors.New("Incorrect url")
}
