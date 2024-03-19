package main

import (
	"dev03"
	"dev03/cmdutils"
	"dev03/fileutils"
	"log"
)

func main() {
	fileName, opt, err := cmdutils.GetArgAndOptions()
	if err != nil {
		log.Fatal(err)
	}
	inFile, err := fileutils.GetInputFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	inLines, err := fileutils.ReadLinesFromFile(inFile, opt)
	if err != nil {
		log.Fatal(err)
	}
	ctx, err := dev03.NewSortContext(opt, inLines)
	if err != nil {
		log.Fatal(err)
	}
	outLines := ctx.Sort()
	outFile, err := fileutils.CreateOutputFile()
	if err != nil {
		log.Fatal(err)
	}
	if err := fileutils.WriteLinesToFile(outFile, outLines); err != nil {
		log.Fatal(err)
	}
}
