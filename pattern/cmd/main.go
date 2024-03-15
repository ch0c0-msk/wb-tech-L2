package main

import (
	"fmt"
	"pattern"
)

func main() {
	fmt.Println("INFO: Facade example")
	computer := pattern.NewComputer(2400, 100, 500)
	computer.Start()
	fmt.Println()
}
