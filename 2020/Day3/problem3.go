package Day3

import (
	"io/ioutil"
	"log"
	"strings"
)

func Problem() {
	log.Println("========== Problem 3 ==========")
	rawData, err := ioutil.ReadFile("data/day3.txt")
	if err != nil {
		panic(err)
	}
	data := strings.Split(string(rawData), "\n")
	tree := traverse(data, 3, 1)

	log.Println("First solution:", tree)
	// Second solution
	var res []int
	for _, v := range [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}} {
		res = append(res, traverse(data, v[0], v[1]))
	}
	var total int = res[0]
	for i := 1; i < len(res); i++ {
		total *= res[i]
	}
	log.Println("Second solution:", total)
	return
}

func traverse(data []string, x_step, y_step int) int {
	var tree = 0
	var x, y = 0, 0
	var step = len(data[0])
	for {
		y += y_step
		if y >= len(data) {
			break
		}
		x += x_step
		if x >= step {
			x = x % step
		}
		if data[y][x] == '#' {
			tree++
		}
	}
	return tree
}
