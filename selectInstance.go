package main

import (
	"fmt"
	"strconv"
)

func selectInstanceIndex(instance_list []instance) int {
	var input string
	fmt.Println("\nWhich one do you want to shh in?")
	fmt.Scanln(&input)
	index, err := strconv.Atoi(input)
	if err != nil || index > len(instance_list)-1 || index < 0 {
		fmt.Println("Please enter a valid number.")
		index = selectInstanceIndex(instance_list)
	}
	return index
}