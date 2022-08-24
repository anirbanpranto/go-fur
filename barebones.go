//main package is the default package of this repo for practicing GoLang
package main

import (
	"fmt"
	"math"
)

//basic functions
func addNums(a int, b int) int {
	return a + b
}

//multiple returns
func minMax(a, b int) (int, int) {
	mn := 0
	mx := 0
	if a > b {
		mx = a
	} else {
		mx = b
	}
	if a < b {
		mn = a
	} else {
		mn = b
	}
	return mn, mx
}

//named return, bad practice generally
func namedRet() (x, y float64) {
	x = math.Pi
	y = x * 3 * 3
	return
}

//variables
func basicVariableTest() {
	var a = 40
	var b = 50
	var c = float64(a*b) + 2.80
	fmt.Println(c)
}

//constants
func constantTest() {
	const someConstVar = 45.0
	fmt.Printf("Type of someConstVar %T", someConstVar)
	//unAssignable
}

func isPrime(x int) bool {
	for i := 2; i <= int(math.Sqrt(float64(x))); i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func isPrimeWeird(x int) bool {
	for i := 2; i <= int(math.Sqrt(float64(x))); i++ {
		//short declaration of y and evaluation of y
		if y := x%i == 0; y {
			return false
		}
	}
	return true
}

func tester(cond bool) {
	if cond {
		fmt.Println("Passed.")
	} else {
		fmt.Println("Failed.")
	}
}

//testing those functions
func testPrime() {
	//this is defer. A defer statement is something that runs after this function has returned itself
	isOk := true
	defer fmt.Println("Finished.")
	for i := 0; i < 1000; i++ {
		a := isPrime(i)
		b := isPrimeWeird(i)
		if a != b {
			isOk = false
		}
	}
	tester(isOk)
	//output :
	//Passed.
	//Finished.
}

/*
Notes : Go does not support runtime constants, means you cannot assign a returned variable from some functions or something in
Go like we do in Javascript. Consts in Go can only store numbers or strings that are already known at compile time
allowed -
const foo = "bar"
const blazeIt = 69420

not allowed -
const ans = isPrime(57) [Just to traumatize you - 57 is not a prime, The factors of 57 are 1, 3, 19, and 57]
*/
