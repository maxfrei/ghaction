package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})

	result := sum(5, 7)
	log.Info(fmt.Sprintf("Result: %d", result))

}

func sum(arg1, arg2 int) int {
	return arg1 + arg2
}
