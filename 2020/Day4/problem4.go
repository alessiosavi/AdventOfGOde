package Day4

import (
	"encoding/json"
	"github.com/alessiosavi/GoGPUtils/helper"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type PassportStruct struct {
	Byr string `json:"byr"`
	Ecl string `json:"ecl"`
	Eyr string `json:"eyr"`
	Hcl string `json:"hcl"`
	Hgt string `json:"hgt"`
	Iyr string `json:"iyr"`
	Pid string `json:"pid"`
}

var reg *regexp.Regexp

func init() {
	var err error
	reg, err = regexp.Compile(`#((\d+)*([a-f]+)*)+`)
	if err != nil {
		panic(err)
	}
}
func (p *PassportStruct) Validate() bool {
	atoi, err := strconv.Atoi(p.Byr)
	if err != nil {
		return false
	}
	if !(atoi >= 1920 && atoi <= 2002) {
		return false
	}

	atoi, err = strconv.Atoi(p.Iyr)
	if err != nil {
		return false
	}
	if !(atoi >= 2010 && atoi <= 2020) {
		return false
	}
	atoi, err = strconv.Atoi(p.Eyr)
	if err != nil {
		return false
	}
	if !(atoi >= 2020 && atoi <= 2030) {
		return false
	}

	switch p.Hgt[len(p.Hgt)-2:] {
	case "cm":
		atoi, err = strconv.Atoi(p.Hgt[0:3])
		if err != nil {
			return false
		}
		if !(atoi >= 150 && atoi <= 193) {
			return false
		}
	case "in":
		atoi, err = strconv.Atoi(p.Hgt[0:2])
		if err != nil {
			return false
		}
		if !(atoi >= 59 && atoi <= 76) {
			return false
		}
	default:
		return false
	}

	if len(p.Hcl) != 7 {
		return false
	}
	if !reg.MatchString(p.Hcl) {
		return false
	}

	switch p.Ecl {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
	default:
		return false
	}
	if len(p.Pid) != 9 {
		return false
	}
	all := strings.ReplaceAll(p.Pid, "0", "")
	_, err = strconv.Atoi(all)
	if err != nil {
		return false
	}
	return true
}

func Problem() {
	log.Println("========== Problem 4 ==========")

	file, err := ioutil.ReadFile("data/day4.txt")
	if err != nil {
		panic(err)
	}
	data := strings.Split(string(file), "\n\n")

	replacer := strings.NewReplacer("\n", " ", "\t", " ")
	for i := range data {
		data[i] = replacer.Replace(data[i])
	}

	var passports []map[string]string

	for _, passport := range data {
		var p map[string]string = make(map[string]string)
		split := strings.Split(passport, " ")
		for _, field := range split {
			f := strings.Split(field, ":")
			p[f[0]] = f[1]
		}
		passports = append(passports, p)
	}

	validPassport := core1(passports)

	log.Println("First solution:", len(validPassport))

	valid := 0
	for _, d := range validPassport {
		if d.Validate() {
			valid++
		}
	}
	log.Println("Second solution:", valid)
}
func core1(passports []map[string]string) []PassportStruct {
	keys := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	valid := 0
	var datastructure []PassportStruct

	for _, passport := range passports {
		isValid := true
		for _, key := range keys {
			if _, ok := passport[key]; !ok {
				isValid = false
				break
			}
		}
		if isValid {
			var d PassportStruct
			raw := []byte(helper.MarshalIndent(passport))
			if err := json.Unmarshal(raw, &d); err != nil {
				panic(err)
			}
			datastructure = append(datastructure, d)
			valid++
		}
	}
	return datastructure
}
