package Day6

import (
	"io/ioutil"
	"log"
	"strings"
)

func Problem() {
	log.Println("========== Problem 6 ==========")

	file, err := ioutil.ReadFile("data/day6.txt")
	if err != nil {
		panic(err)
	}

	data := strings.Split(string(file), "\n\n")
	//for i := range data {
	//	data[i] = strings.ReplaceAll(data[i], "\n", "")
	//}

	log.Println("First solution:", core1(data))
	log.Println("Second solution:", core2(data))

}

func core1(data []string) int {
	var uniq int
	for _, line := range data {
		line = strings.ReplaceAll(line, "\n", "")
		var m = make(map[rune]struct{})
		for _, c := range line {
			m[c] = struct{}{}
		}
		uniq += len(m)
	}
	return uniq
}
func core2(data []string) int {
	var uniq int
	for _, lines := range data {
		line := strings.Split(lines, "\n")
		var m = make(map[rune]int)
		for _, row := range line {
			for _, c := range row {
				if _, ok := m[c]; !ok {
					m[c] = 1
				} else {
					m[c]++
				}
			}
		}
		if len(line) > 1 {
			for k := range m {
				if m[k] == len(line) {
					uniq++
				}
			}
		} else {
			uniq += len(m)
		}
	}
	return uniq

}
