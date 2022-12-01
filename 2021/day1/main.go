package main

import (
	"fmt"
	"log"
	advent_of_code "nullawhale.com/aoc"
)

func fetchInput(filename string) ([]int, error) {
	data, err := advent_of_code.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	res, err := advent_of_code.StringSliceToIntSlice(data)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func star1(data []int) (int, error) {
	var sum = 0

	for i := 1; i < len(data); i++ {
		if data[i] > data[i-1] {
			sum++
		}
	}

	return sum, nil
}

func star2(data []int) (int, error) {
	var a2 []int
	var sum = 0

	for i := 1; i < len(data)-1; i++ {
		a2 = append(a2, data[i-1]+data[i]+data[i+1])
	}
	for i := 1; i < len(a2); i++ {
		if a2[i] > a2[i-1] {
			sum++
		}
	}

	return sum, nil
}

func main() {
	data, err := fetchInput("i")
	if err != nil {
		log.Fatal(err)
	}

	star, err := star1(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("First star: %d\n", star)

	star, err = star2(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Second star: %d\n", star)
}
