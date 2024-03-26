package main

import (
	"bufio"
	"dev08"
	"os"
)

// TODO make fork/exec
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		err := dev08.PrintPrefix()
		if err != nil {
			panic(err)
		}
		scanner.Scan()
		tasks := scanner.Text()
		if tasks == "\\quit" {
			return
		}
		commands, err := dev08.ParsePipes(tasks)
		if err != nil {
			panic(err)
		}
		for _, cmd := range commands {
			err = cmd.Exec()
			if err != nil {
				panic(err)
			}
		}
	}
}
