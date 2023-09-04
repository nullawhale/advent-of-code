package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Command struct {
	move string
	step int
}

func fetchInput() (res []Command) {
	data, err := os.Open("i")
	if err != nil {
		fmt.Printf("Err: %s", err)
	}
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		str := strings.Split(scanner.Text(), " ")
		s, _ := strconv.Atoi(str[1])
		res = append(res, Command{
			move: str[0],
			step: s,
		})
	}

	return res
}

func star1() int {
	var horizontal = 0
	var depth = 0
	var data = fetchInput()

	for _, c := range data {
		switch c.move {
		case "forward":
			horizontal += c.step
			break
		case "down":
			depth += c.step
			break
		case "up":
			depth -= c.step
			break
		}
	}

	return horizontal * depth
}

func star2() int {
	var horizontal = 0
	var aim = 0
	var depth = 0
	var data = fetchInput()

	for _, c := range data {
		switch c.move {
		case "forward":
			horizontal += c.step
			depth += c.step * aim
			break
		case "down":
			aim += c.step
			break
		case "up":
			aim -= c.step
			break
		}
	}

	return horizontal * depth
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Printf("First star: %d\n", star1())
	fmt.Printf("Second star: %d\n", star2())
}
