package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Map struct {
	rows   []MapRow
	height int
	width  int
}

type MapRow struct {
	columns []string
}

func fetchInput() (m Map) {
	data, err := ioutil.ReadFile("i")
	if err != nil {
		fmt.Printf("Err: %s", err)
	}
	input := strings.SplitN(string(data), "\n", 323)
	for _, line := range input {
		line = strings.TrimSuffix(line, "\n")
		row := MapRow{}
		chars := strings.Split(line, "")
		m.width = len(chars)
		for _, char := range chars {
			row.columns = append(row.columns, char)
		}
		m.rows = append(m.rows, row)
	}
	m.height = len(input)
	return m
}

func counter(cordX int, cordY int) int {
	var count = 0
	var stepR = 0
	var stepD = 0
	var data = fetchInput()

	for stepR < data.height {
		char := data.rows[stepR].columns[stepD%data.width]
		if char == "#" {
			count++
		}

		stepR += cordY
		stepD += cordX
	}

	return count
}

func star1() int {
	return counter(3, 1)
}

func star2() int {
	return counter(1, 1) *
		counter(3, 1) *
		counter(5, 1) *
		counter(7, 1) *
		counter(1, 2)
}

func main() {
	fmt.Printf("\nFirst star: %d\n", star1())
	fmt.Printf("Second star: %d\n", star2())
}
