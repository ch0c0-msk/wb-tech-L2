package cmdutils

import (
	"errors"
	"flag"
)

type Options struct {
	KFlag int
	NFlag bool
	RFlag bool
	UFlag bool
}

func GetArgAndOptions() (string, *Options, error) {
	opt := &Options{}
	flag.IntVar(&opt.KFlag, "k", 1, "key column number")
	flag.BoolVar(&opt.NFlag, "n", false, "sorting by numeric values")
	flag.BoolVar(&opt.RFlag, "r", false, "output the result in reverse order")
	flag.BoolVar(&opt.UFlag, "u", false, "do not output duplicate lines")
	flag.Parse()
	if !isOptionsCorrect(opt) {
		return "", nil, errors.New("invalid flag values")
	}
	fileName := flag.Arg(0)
	if len(fileName) == 0 {
		return "", nil, errors.New("input file name is missing")
	}
	return fileName, opt, nil
}

func isOptionsCorrect(opt *Options) bool {
	return opt.KFlag > 0
}
