package main

import (
	"euler/common"
	euler001 "euler001/solution"
	"fmt"
)

func main() {
	if result, err := euler001.Solution(); err == nil {
		fmt.Println("Problem 001:", result)
	} else {
		fmt.Println("Error:", err)
	}
}
