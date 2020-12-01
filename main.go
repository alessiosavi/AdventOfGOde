package main

import (
	day1 "github.com/alessiosavi/AdventOfGOde/Day1"
	"log"
	"time"
)

func main() {

	var problems []func()

	problems = append(problems, day1.Problem1)

	for i, fn := range problems {
		start := time.Now()
		fn()
		duration := time.Since(start)
		log.Printf("Function [%d] duration: %v\n", i+1, duration)
	}

}
