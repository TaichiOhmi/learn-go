package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(os.Args)
	args := os.Args[1:]
	var iArgs = []int{}
	for _, i := range args {
		value, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		iArgs = append(iArgs, value)
	}
	max := 0
	for _, value := range iArgs {
		if value > max {
			max = value
		}
	}
	fmt.Println("Max Value :", max)
}
