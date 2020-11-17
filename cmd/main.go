package main

import (
	"Noah/internal/library"
	"time"
)

func main() {
	targetTime := time.Now()
	library.Scanner()

	println(time.Since(targetTime).String())
}
