package Day3

import (
	"log"
	"strings"
	"testing"
)

var raw = `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`

func TestProblem3(t *testing.T) {
	data := strings.Split(raw, "\n")
	var tree int = 0
	var x = 0
	for y := 1; y < len(data); y++ {
		x += 3
		if x >= 11 {
			x = x % 11
		}
		if data[y][x] == '#' {
			tree++
		}
	}
	log.Println(tree)
}
