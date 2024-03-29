// In this list, the two entries that sum to 2020 are 1721 and 299.
// Multiplying them together produces 1721 * 299 = 514579, so the correct answer is 514579.
// Of course, your expense report is much larger.
// Find the two entries that sum to 2020; what do you get if you multiply them together?
package Day1

import (
	fileutils "github.com/alessiosavi/GoGPUtils/files"
	stringutils "github.com/alessiosavi/GoGPUtils/string"
	"log"
	"strconv"
)

func Problem() {
	log.Println("========== Problem 1 ==========")
	rawData := fileutils.ReadFileInArray("data/day1.txt")
	data := make([]int, len(rawData))
	for i := range rawData {
		atoi, err := strconv.Atoi(stringutils.Trim(rawData[i]))
		if err != nil {
			panic(err)
		}
		data[i] = atoi
	}

	log.Println("First solution:", core1(data))
	log.Println("Second solution:", core2(data))
}

func core2(data []int) int {
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			for k := j + 1; k < len(data); k++ {
				if data[i]+data[j]+data[k] == 2020 {
					return data[i] * data[j] * data[k]
				}
			}
		}
	}
	return 0
}

func core1(data []int) int {
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i]+data[j] == 2020 {
				return data[i] * data[j]
			}
		}
	}
	return 0
}
