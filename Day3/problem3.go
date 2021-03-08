package Day3

import (
	fileutils "github.com/alessiosavi/GoGPUtils/files"
	"log"
)

func Problem3() {
	rawData := fileutils.ReadFileInArray("Day3/day3.txt")
	//for _, data := range rawData {
	//	log.Println(data)
	//}
	log.Println(rawData[0])
	log.Println(len(rawData[0]))
}
