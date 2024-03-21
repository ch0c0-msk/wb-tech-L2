package cmdutils

import (
	"errors"
	"flag"
)

type Options struct {
	AFlag     int
	BFlag     int
	CFlag     int
	CountFlag bool
	IFlag     bool
	VFlag     bool
	FFlag     bool
	NFlag     bool
}

func GetArgsAndOptions() (string, string, *Options, error) {
	opt := &Options{}
	flag.IntVar(&opt.AFlag, "A", 0, "After")
	flag.IntVar(&opt.BFlag, "B", 0, "Before")
	flag.IntVar(&opt.CFlag, "C", 0, "Context")
	flag.BoolVar(&opt.CountFlag, "c", false, "Count")
	flag.BoolVar(&opt.IFlag, "i", false, "Register")
	flag.BoolVar(&opt.VFlag, "v", false, "Invert")
	flag.BoolVar(&opt.FFlag, "F", false, "Fixed pattern")
	flag.BoolVar(&opt.NFlag, "n", false, "Number of string")
	flag.Parse()
	if !isOptionsCorrect(opt) {
		return "", "", nil, errors.New("invalid flag values")
	}
	args := flag.Args()
	regExpStr, fileName, err := parseArgs(args, opt)
	if err != nil {
		return "", "", nil, errors.New("invalid args")
	}
	return regExpStr, fileName, opt, nil
}

func isOptionsCorrect(opt *Options) bool {
	return (opt.AFlag >= 0) && (opt.BFlag >= 0) && (opt.CFlag >= 0)
}

func parseArgs(args []string, opt *Options) (string, string, error) {
	var regExp, fileName string
	switch len(args) {
	case 0:
		return regExp, fileName, errors.New("missing pattern")
	case 1:
		regExp = args[0]
		if opt.IFlag {
			regExp = "(?i)" + regExp
		}
		if opt.FFlag {
			regExp = `\Q` + regExp + `\E`
		}
		fileName = ""
	default:
		regExp = args[0]
		if opt.IFlag {
			regExp = "(?i)"
		}
		if opt.FFlag {
			regExp = `\Q` + regExp + `\E`
		}
		for i := 1; i < len(args)-1; i++ {
			if opt.FFlag {
				regExp += "|" + `\Q` + args[i] + `\E`
			} else {
				regExp += "|" + args[i]
			}
		}
		fileName = args[1]
	}
	return regExp, fileName, nil
}
