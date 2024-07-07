package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <int> <int>")
		return
	}

	num1 := os.Args[1]
	num2 := os.Args[2]

	isNum := checkNum(num1, num2)
	var a float64
	var b float64

	if isNum {
		a, _ = strconv.ParseFloat(num1, 32)
		b, _ = strconv.ParseFloat(num2, 32)
	} else {
		fmt.Println("Usage: go run . <int> <int>")
		return
	}

	sum := Add(a, b)
	mul := Mult(a, b)
	div := Div(a, b)

	fmt.Printf("The sum is: %.3f \n", sum)
	fmt.Printf("The product is: %.3f \n", mul)
	fmt.Printf("The quotient is: %.3f \n", div)

}

// Function Add() takes two digits and returns their sum
func Add(a, b float64) float64 {
	return a + b
}

// Function Mult() take two integers and returns their product
func Mult(a, b float64) float64 {
	return a * b
}

// Function Div() takes two integers and returns the quotient
func Div(a, b float64) float64 {
	return float64(a / b)
}

// Func isNum() check if the values given are nums
func checkNum(a, b string) bool {
	num1 := true
	num2 := true
	decimal1 := false
	decimal2 := false

	for _, n := range a {
		if (n >= '0' && n <= '9') || n == '.' {
			if n == '.' && decimal1 == false {
				decimal1 = true
				continue
			}
			if n == '.' && decimal1 == true {
				return false
			}
			continue
		} else {
			num1 = false
			break
		}
	}

	for _, n := range b {
		if (n >= '0' && n <= '9') || n == '.' {
			if n == '.' && decimal2 == false {
				decimal2 = true
				continue
			}
			if n == '.' && decimal2 == true {
				return false
			}
			continue
		} else {
			num1 = false
			break
		}
	}

	if !num1 || !num2 {
		return false
	}
	return true
}
