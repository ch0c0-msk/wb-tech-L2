package cmdutils

import (
	"errors"
	"flag"
)

type Options struct {
	FFlag int
	DFlag string
	SFlag bool
}

func GetOptions() (*Options, error) {
	opt := new(Options)
	flag.IntVar(&opt.FFlag, "f", 0, "fields")
	flag.StringVar(&opt.DFlag, "d", "\t", "delimeter")
	flag.BoolVar(&opt.SFlag, "s", false, "separeted only")
	flag.Parse()
	if !isOptionsCorrect(opt) {
		return nil, errors.New("wrong options")
	}
	return opt, nil
}

func isOptionsCorrect(opt *Options) bool {
	return opt.FFlag > 0
}
