package main

import (
	"fmt"
	"study-golang/12-Test-testfy/sub"
	"study-golang/12-Test-testfy/sum"
)

func main() {
	sumResult := sum.Sum(1, 2)
	subResult := sub.Sub(3, 2)
	fmt.Println(sumResult)
	fmt.Println(subResult)
	fmt.Printf("Hello!")
}
