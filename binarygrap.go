package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := 0
	fmt.Print("Enter integer to find binary grap = ")
	fmt.Scanf("%d", &input)
	fmt.Println("Hello ", input)
	num := int64(input)
	binary := strconv.FormatInt(num, 2)
	s := strings.Split(binary, "")
	fmt.Println(s)
	countZero(s)
}

func countZero(num []string) {
	countZero := 0
	tempZero := 0
	lenLoop := len(num)
	for i := 0; i < lenLoop; i++ {
		if num[i] == "0" {
			countZero++
		} else {
			countZero = 0
		}
		if tempZero < countZero {
			tempZero = countZero

		}
	}
	fmt.Println("have 0 = ", tempZero)
}
