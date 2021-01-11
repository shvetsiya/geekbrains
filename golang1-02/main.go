package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main1() {
	epsilon := math.Nextafter(1, 2) - 1 // this is machine epsilon to work with floating points
	var snum1, snum2, op string
	var res float64

	fmt.Println("Enter first number: ")
	fmt.Scanln(&snum1)

	fmt.Println("Enter second number: ")
	fmt.Scanln(&snum2)

	fmt.Println("Enter first operator: ")
	fmt.Scanln(&op)

	num1, err := strconv.ParseFloat(snum1, 64)
	if err != nil {
		fmt.Printf("Your first number: %s is incorrect!", snum1)
		os.Exit(1)
	}

	num2, err := strconv.ParseFloat(snum2, 64)
	if err != nil {
		fmt.Printf("Your second number: %s is incorrect!", snum2)
		os.Exit(1)
	}

	switch op {
	case "+":
		res = num1 + num2
	case "-":
		res = num1 - num2
	case "*":
		res = num1 * num2
	case "/":
		res = num1 / (num2 + epsilon)
	case "^":
		res = math.Pow(num1, num2)
	case "Min":
		res = math.Min(num1, num2)
	case "Max":
		res = math.Max(num1, num2)
	default:
		fmt.Print("Not supported operation!")
		os.Exit(1)
	}

	fmt.Printf("The result: %f\n", res)
}
