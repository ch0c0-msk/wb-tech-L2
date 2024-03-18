package main

import (
	"dev01"
	"fmt"
	"os"
)

func main() {
	currentTime, correctTime, err := dev01.GetCurrentAndCorrectTime()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	fmt.Printf("Current time: %s\nCorrect NTP time: %s\n", currentTime.String(), correctTime.String())
}
