package Day5

import (
	fileutils "github.com/alessiosavi/GoGPUtils/files"
	"log"
	"math"
	"sort"
)

func Problem() {
	log.Println("========== Problem 5 ==========")
	seats := fileutils.ReadFileInArray("data/day5.txt")
	maxRow := 127
	maxCol := 7
	var val []int
	for _, seat := range seats {
		seatRow := seat[0 : len(seat)-3]
		seatCol := seat[len(seat)-3:]
		row := core(seatRow, maxRow)
		col := core(seatCol, maxCol)
		val = append(val, uniqSeatID(row, col))
	}
	sort.Ints(val)
	log.Println("First solution:", val[len(val)-1])

	for i, v := range val {
		if i+val[0] != v {
			log.Println("Second solution:", v-1)
			break
		}
	}
}

func uniqSeatID(row, col int) int {
	return row*8 + col
}
func core(seatRow string, maxRow int) int {
	m, M := 0, maxRow
	for _, c := range seatRow {
		if c == 'F' || c == 'L' {
			M -= int(math.Ceil(float64(M-m) / 2.0))
		} else if c == 'B' || c == 'R' {
			m += int(math.Round(float64(M-m) / 2.0))
		} else {
			panic("Row not managed! " + seatRow)
		}
	}
	return m
}
