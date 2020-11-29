package util

import (
	"log"
	"strconv"
)

func Str2Int(str string) int {
	i, err := strconv.Atoi(str)
	if err!=nil {
		log.Fatal("str can not cast to int")
	}

	return i
}
