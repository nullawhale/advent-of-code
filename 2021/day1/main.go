package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func fetchInput() (res []int) {
	data, err := ioutil.ReadFile("i")
	if err != nil {
		fmt.Printf("Err: %s", err)
	}
	str := strings.SplitN(string(data), "\n", 2000)
	for _, s := range str {
		num, _ := strconv.Atoi(strings.TrimSuffix(s, "\n"))
		res = append(res, num)
	}

	return res
}

func star1() int {
	var sum = 0
	var data = fetchInput()

	for i := 1; i < len(data); i++ {
		if data[i] > data[i-1] {
			sum++
		}
	}

	return sum
}

func star2() int {
	var a2 []int
	var sum = 0
	var data = fetchInput()

	for i := 1; i < len(data)-1; i++ {
		a2 = append(a2, data[i-1]+data[i]+data[i+1])
	}
	for i := 1; i < len(a2); i++ {
		if a2[i] > a2[i-1] {
			sum++
		}
	}

	return sum
}

func main() {
	fmt.Printf("First star: %d\n", star1())
	fmt.Printf("Second star: %d\n", star2())
}
