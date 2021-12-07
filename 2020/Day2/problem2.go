package Day2

import (
	fileutils "github.com/alessiosavi/GoGPUtils/files"
	"log"
	"strconv"
	"strings"
)

type Datastructure struct {
	Min      int
	Max      int
	Char     string
	Password string
}

func Problem() {

	var (
		correct1, correct2, min, max = 0, 0, 0, 0
		err                          error
	)
	log.Println("========== Problem 2 ==========")
	rawData := fileutils.ReadFileInArray("data/day2.txt")

	for _, data := range rawData {
		if strings.Count(data, " ") != 2 {
			panic("Inconsistent input")
		}

		split := strings.Split(data, " ")
		n := strings.Split(split[0], "-")
		if min, err = strconv.Atoi(n[0]); err != nil {
			panic(err)
		}
		if max, err = strconv.Atoi(n[1]); err != nil {
			panic(err)
		}
		letter := split[1]
		password := split[2]
		letter = string(letter[0])
		nCount := strings.Count(password, letter)
		if nCount >= min && nCount <= max {
			correct1++
		}
		if (string(password[min-1]) == letter || string(password[max-1]) == letter) &&
			!(string(password[min-1]) == letter && string(password[max-1]) == letter) {
			correct2++
		}
	}

	log.Println("First solution:", correct1)
	log.Println("Second solution:", correct2)
	return
}
