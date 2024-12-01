package parse

import (
	"log"
	"strconv"
)

func Int[T []byte | string | rune](s T) int {
	i, err := strconv.Atoi(string(s))
	if err != nil {
		log.Fatalf("failed to convert \"%s\" to int\n", s)
	}
	return i
}
