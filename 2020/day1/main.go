package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func fetchInput() (res []int) {
    data, err := ioutil.ReadFile("i")
    if err != nil {
	fmt.Printf("Err: %s", err)
    }
    str := strings.SplitN(string(data), "\n", 200)
    for _, s := range str {
	num, _ := strconv.Atoi(strings.TrimSuffix(s, "\n"))
	res = append(res, num)
    }

    return res
}

func star1() int {
    var sum = 0
    var data = fetchInput()

    for i := 0; i < len(data); i++ {
	for j := 0; j < len(data); j++ {
	    if i != j {
		if data[i] + data[j] == 2020 {
		    return data[i] * data[j]
		}
	    }
	}
    }

    return sum
}

func star2() int {
    var sum = 0
    var data = fetchInput()

    for i := 0; i < len(data); i++ {
	for j := 0; j < len(data); j++ {
	    for k := 0; k < len(data); k++ {
		if i != j && i != k && j != k {
		    if data[i] + data[j] + data[k] == 2020 {
			return data[i] * data[j] * data[k]
		    }
		}
	    }
	}
    }

    return sum
}

func main() {
    fmt.Printf("First star: %d\n", star1())
    fmt.Printf("Second star: %d\n", star2())
}
