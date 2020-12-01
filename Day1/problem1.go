// In this list, the two entries that sum to 2020 are 1721 and 299.
// Multiplying them together produces 1721 * 299 = 514579, so the correct answer is 514579.
// Of course, your expense report is much larger.
// Find the two entries that sum to 2020; what do you get if you multiply them together?
package day1

import (
	"fmt"
	fileutils "github.com/alessiosavi/GoGPUtils/files"
	stringutils "github.com/alessiosavi/GoGPUtils/string"
	"strconv"
)

func Problem1() {
	rawData := fileutils.ReadFileInArray("Day1/day1.txt")
	data := make([]int, len(rawData))
	for i := range rawData {
		atoi, err := strconv.Atoi(stringutils.Trim(rawData[i]))
		if err != nil {
			panic(err)
		}
		data[i] = atoi
	}

	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i]+data[j] == 2020 {
				fmt.Println(data[i] * data[j])
				return
			}
		}
	}
}
