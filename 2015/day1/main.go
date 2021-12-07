package main

import (
	"fmt"
	"io/ioutil"
)

func fetchInput() string {
    data, err := ioutil.ReadFile("i")
    if err != nil {
	fmt.Printf("Err: %s", err)
    }
    return string(data)
}

func star1() int {
    var flour = 0
    var s = fetchInput()

    for _, c := range s {
	switch string(c) {
	case "(":
	    flour += 1
	case ")":
	    flour -= 1
	}
    }

    return flour
}


func star2() int {
    var flour = 0
    var s = fetchInput()
    //var s = "()())"

    for i, c := range s {
	i++
	switch string(c) {
	case "(":
	    flour++
	case ")":
	    flour--
	}
	//fmt.Printf("%d: %s (%d)\n", i, string(c), flour)
	if flour == -1 {
	    return i
	}
    }

    return flour
}

func main() {
    //fmt.Println(star1())
    fmt.Println(star2())
}
