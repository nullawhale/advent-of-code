package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"sort"
)

type Combin struct {
    C map[int]int
}

func combinatorika(c Combin, data []int) int {
	if co, exists := c.C[data[0]]; exists {
	    return co
	}

	co := 0
	for i := 0; i < len(data)-1; i++ {
	    for n := 1; i+n+1 < len(data); n++ {
		var j = data[i+n+1] - data[i]
		if j <= 3 {
			co = co + 1 + combinatorika(c, data[i+n+1:])
		}
	    }
	}

	c.C[data[0]] = co
	return co
}

func fetchInput() (res []int) {
    data, err := ioutil.ReadFile("i")
    if err != nil {
	fmt.Printf("Err: %s", err)
    }
    str := strings.SplitN(string(data), "\n", 104)
    for _, s := range str {
	num, _ := strconv.Atoi(strings.TrimSuffix(s, "\n"))
	res = append(res, num)
    }

    return res
}

func star1() int {
    var data = fetchInput()

    sort.Ints(data)
    data = append(make([]int, 1), data...)
    data = append(data, data[len(data)-1] + 3)

    var j1 = 0
    var j3 = 0
    rI := len(data) - 1
    for i, _ := range data[:rI] {
	    var jolt = data[i+1] - data[i]
	    switch jolt {
		case 3:
		    j3++
		case 0, 2:
		    break
		case 1:
		    j1++
	    }
    }

    return j1 * j3
}

func star2() int {
    var data = fetchInput()

    sort.Ints(data)
    data = append(make([]int, 1), data...)
    data = append(data, data[len(data)-1] + 3)

    var comb = &Combin{
	C: make(map[int]int),
    }

    return combinatorika(*comb, data) + 1
}

func main() {
    fmt.Printf("First star: %d\n", star1())
    fmt.Printf("Second star: %d\n", star2())
}
