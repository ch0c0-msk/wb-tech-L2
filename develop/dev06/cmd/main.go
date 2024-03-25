package main

import (
	"dev06"
	"dev06/cmdutils"
	"dev06/fileutils"
	"log"
	"os"
)

func main() {
	opt, err := cmdutils.GetOptions()
	if err != nil {
		log.Fatal(err)
	}
	inLines, err := fileutils.ReadLinesFromFile(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	outLines := dev06.Cut(inLines, opt)
	err = fileutils.WriteLinesToFile(os.Stdout, outLines)
	if err != nil {
		log.Fatal(err)
	}
}
