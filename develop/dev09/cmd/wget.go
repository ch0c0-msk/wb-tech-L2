package main

import (
	"dev09"
	"dev09/fileutils"
	"errors"
	"flag"
)

var url string

func main() {
	data, err := dev09.DownloadSite(url)
	if err != nil {
		panic(err)
	}
	err = fileutils.WriteToFile(data)
	if err != nil {
		panic(err)
	}
}

func init() {
	flag.Parse()
	url = flag.Arg(0)
	if url == "" {
		panic(errors.New("empty url"))
	}
}
