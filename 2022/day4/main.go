package main

import (
	"errors"
	"fmt"
	"log"
	advent_of_code "nullawhale.com/aoc"
	"strconv"
	"strings"
)

type Pair struct {
	left  string
	right string
}

func fetchInput(filename string) ([]string, error) {
	data, err := advent_of_code.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func star1(data []string) int {
	pairs, err := parsePairs(data)
	if err != nil {
		log.Fatal(err)
	}

	res := 0
	for _, pair := range pairs {
		leftRange := strings.Split(pair.left, "-")
		rightRange := strings.Split(pair.right, "-")
		if len(leftRange) != 2 || len(rightRange) != 2 {
			log.Fatal("mismatch number of elements in range")
		}

		if checkFullyOverlap(leftRange, rightRange) {
			res++
		}
	}

	return res
}

func star2(data []string) int {
	pairs, err := parsePairs(data)
	if err != nil {
		log.Fatal(err)
	}

	res := 0
	for _, pair := range pairs {
		leftRange := strings.Split(pair.left, "-")
		rightRange := strings.Split(pair.right, "-")
		if len(leftRange) != 2 || len(rightRange) != 2 {
			log.Fatal("mismatch number of elements in range")
		}

		if checkOverlap(leftRange, rightRange) {
			res++
		}
	}

	return res
}

func checkFullyOverlap(left []string, right []string) bool {
	ll, _ := strconv.Atoi(left[0])
	lr, _ := strconv.Atoi(left[1])
	rl, _ := strconv.Atoi(right[0])
	rr, _ := strconv.Atoi(right[1])

	if ll <= rl && lr >= rr || rl <= ll && rr >= lr {
		return true
	}
	return false
}

func checkOverlap(left []string, right []string) bool {
	ll, _ := strconv.Atoi(left[0])
	lr, _ := strconv.Atoi(left[1])
	rl, _ := strconv.Atoi(right[0])
	rr, _ := strconv.Atoi(right[1])

	if ll <= rr && rl <= lr {
		return true
	}
	return false
}

func parsePairs(sections []string) ([]Pair, error) {
	var res []Pair
	for _, section := range sections {
		dividedSection := strings.Split(section, ",")
		if len(dividedSection) != 2 {
			return nil, errors.New("mismatch number of elements in pair")
		}
		res = append(res, Pair{
			left:  dividedSection[0],
			right: dividedSection[1],
		})
	}
	return res, nil
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	data, err := fetchInput("i")
	if err != nil {
		log.Fatal(err)
	}

	star := star1(data)
	fmt.Printf("First star: %d\n", star)

	star = star2(data)
	fmt.Printf("Second star: %d\n", star)
}
