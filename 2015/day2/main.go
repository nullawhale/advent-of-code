package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func fetchInput() []string {
    data, err := ioutil.ReadFile("i")
    if err != nil {
	fmt.Printf("Err: %s", err)
    }
    return strings.SplitN(string(data), "\n", 1000)
}


func min(x int, y int, z int) (min int) {
    var nums = []int{x, y, z}
    min = nums[0]
    for _, a := range nums {
	if a < min {
	    min = a
	}
    }

    return min
}

func star1() int {
    var flour = 0
    var data = fetchInput()
    //var data = []string{"2x3x4", "1x1x10"}

    for _, s := range data {
	s = strings.TrimSuffix(s, "\n")
	var tmp = strings.Split(s, "x")
	var l, _ = strconv.Atoi(tmp[0])
	var w, _ = strconv.Atoi(tmp[1])
	var h, _ = strconv.Atoi(tmp[2])

	var surfaceArea = 2*l*w + 2*w*h + 2*h*l

	flour += surfaceArea + min(l*w, w*h, h*l)
    }

    return flour
}


func star2() int {
    var flour = 0
    var data = fetchInput()
    //var data = []string{"2x3x4", "1x1x10"}

    for _, s := range data {
	s = strings.TrimSuffix(s, "\n")
	var tmp = strings.Split(s, "x")
	var l, _ = strconv.Atoi(tmp[0])
	var w, _ = strconv.Atoi(tmp[1])
	var h, _ = strconv.Atoi(tmp[2])

	var ribbon = min(l+l+w+w, l+l+h+h, w+w+h+h)

	flour += ribbon + l*w*h
    }

    return flour
}

func main() {
    fmt.Println(star1())
    fmt.Println(star2())
}
