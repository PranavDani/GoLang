package main

import (
	"fmt"
	"math/big"
	// "math/big"
	// "math/rand"
	"crypto/rand"
	// "math/rand"
)

func main() {
	fmt.Println("Welcome to Maths command in Golang")

	// var myNumberOne int = 2
	// var myNumberTwo float64 = 4

	// fmt.Println("The sum is ", myNumberOne + int(myNumberTwo))

	// Random number

	// fmt.Println(rand.Intn(5))

	myRandomNumber, _ := rand.Int(rand.Reader, big.NewInt(5))

	fmt.Println(myRandomNumber)

}