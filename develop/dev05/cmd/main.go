package main

import (
	"dev05"
	"dev05/cmdutils"
	"dev05/fileutils"
	"log"
	"os"
	"regexp"
)

func main() {
	regExpStr, fileName, opt, err := cmdutils.GetArgsAndOptions()
	if err != nil {
		log.Fatal(err)
	}
	inputFile, err := fileutils.GetInputFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	inputLines, err := fileutils.ReadLinesFromFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	regExp, err := regexp.Compile(regExpStr)
	if err != nil {
		log.Fatal(err)
	}
	outputLines := dev05.Grep(inputLines, regExp, opt)
	fileutils.WriteLinesToFile(os.Stdout, outputLines)
}
