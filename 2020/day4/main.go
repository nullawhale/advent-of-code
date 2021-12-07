package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func fetchInput() (passports []Passport) {
	data, err := ioutil.ReadFile("i")
	if err != nil {
		fmt.Printf("Err: %s", err)
	}
	str := strings.Split(string(data), "\n\n")
	for _, s := range str {
		s = strings.TrimSuffix(s, "\n")
		temp := strings.Split(s, "\n")
		var passport Passport
		var byr, iyr, eyr, hgt, hcl, ecl, pid, cid string
		for _, t := range temp {
			items := strings.Split(t, " ")
			for _, i := range items {
				item := strings.Split(i, ":")
				switch item[0] {
				case "byr":
					byr = item[1]
				case "iyr":
					iyr = item[1]
				case "eyr":
					eyr = item[1]
				case "hgt":
					hgt = item[1]
				case "hcl":
					hcl = item[1]
				case "ecl":
					ecl = item[1]
				case "pid":
					pid = item[1]
				case "cid":
					cid = item[1]
				}
			}
		}
		passport = Passport{byr: byr, iyr: iyr, eyr: eyr, hgt: hgt, hcl: hcl, ecl: ecl, pid: pid, cid: cid}
		passports = append(passports, passport)
	}

	return passports
}

func star1() int {
	var count = 0
	var data = fetchInput()

	for _, i := range data {
		if i.byr != "" && i.iyr != "" && i.eyr != "" && i.hgt != "" && i.hcl != "" && i.ecl != "" && i.pid != "" {
			count++
		}
	}

	return count
}

func star2() int {
	var count = 0
	var data = fetchInput()

	hclR, _ := regexp.Compile("^#([a-fA-F0-9]{6}|[a-fA-F0-9]{3})$")

	for _, i := range data {
		if i.byr != "" && i.iyr != "" && i.eyr != "" && i.hgt != "" && i.hcl != "" && i.ecl != "" && i.pid != "" {
			byrI, _ := strconv.Atoi(i.byr)
			iyrI, _ := strconv.Atoi(i.iyr)
			eyrI, _ := strconv.Atoi(i.eyr)
			if (len(i.byr) == 4 && byrI >= 1920 && byrI <= 2002) &&
				(len(i.iyr) == 4 && iyrI >= 2010 && iyrI <= 2020) &&
				(len(i.eyr) == 4 && eyrI >= 2020 && eyrI <= 2030) &&
				(hclR.MatchString(i.hcl)) &&
				(i.ecl == "amb" || i.ecl == "blu" || i.ecl == "brn" || i.ecl == "gry" ||
					i.ecl == "grn" || i.ecl == "hzl" || i.ecl == "oth") &&
				len(i.pid) == 9 {
				var hCm int
				var hIn int
				if strings.HasSuffix(i.hgt, "cm") {
					hCm, _ = strconv.Atoi(strings.TrimSuffix(i.hgt, "cm"))
					if hCm >= 150 && hCm <= 193 {
						count++
					}
				} else if strings.HasSuffix(i.hgt, "in") {
					hIn, _ = strconv.Atoi(strings.TrimSuffix(i.hgt, "in"))
					if hIn >= 59 && hIn <= 76 {
						count++
					}
				}
			}
		}
	}

	return count
}

func main() {
	fmt.Printf("First star: %d\n", star1())
	fmt.Printf("Second star: %d\n", star2())
}
