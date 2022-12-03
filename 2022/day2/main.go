package main

import (
	"fmt"
	"log"
	advent_of_code "nullawhale.com/aoc"
	"strings"
)

type Pair struct {
	elf int
	me  int
}

func fetchInput(filename string) ([]string, error) {
	data, err := advent_of_code.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func star1(data []string) int {
	pairs := parsePairs(data)

	res := 0
	for _, pair := range pairs {
		if (pair.elf+1)%3 == pair.me {
			res += pair.me + 1 + 6
		} else if pair.elf == pair.me {
			res += pair.me + 1 + 3
		} else {
			res += pair.me + 1
		}
	}

	return res
}

func star2(data []string) int {
	pairs := parsePairs(data)

	res := 0
	for _, pair := range pairs {
		switch pair.me {
		case 0:
			res += (pair.elf+2)%3 + 1
		case 1:
			res += 3 + pair.elf + 1
		case 2:
			res += 6 + (pair.elf+1)%3 + 1
		}
	}

	return res
}

func parsePairs(pairs []string) []Pair {
	var res []Pair
	for _, pair := range pairs {
		pairSlice := strings.Split(pair, " ")
		var tempPair Pair
		switch pairSlice[0] {
		case "A":
			tempPair.elf = 0
		case "B":
			tempPair.elf = 1
		case "C":
			tempPair.elf = 2
		}
		switch pairSlice[1] {
		case "X":
			tempPair.me = 0
		case "Y":
			tempPair.me = 1
		case "Z":
			tempPair.me = 2
		}
		res = append(res, tempPair)
	}
	return res
}

func main() {
	data, err := fetchInput("i")
	if err != nil {
		log.Fatal(err)
	}

	star := star1(data)
	fmt.Printf("First star: %d\n", star)

	star = star2(data)
	fmt.Printf("Second star: %d\n", star)
}
