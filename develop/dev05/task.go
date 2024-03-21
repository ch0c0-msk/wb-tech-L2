package dev05

import (
	"dev05/cmdutils"
	"fmt"
	"regexp"
	"strconv"
)

// Grep filters text by pattern
func Grep(lines []string, regExp *regexp.Regexp, opt *cmdutils.Options) []string {
	matchLineNums := match(lines, regExp, opt)

	if opt.CountFlag {
		return []string{strconv.Itoa(len(matchLineNums))}
	}

	var after, before int
	if opt.AFlag > opt.CFlag {
		after = opt.AFlag
	} else {
		after = opt.CFlag
	}
	if opt.BFlag > opt.CFlag {
		before = opt.BFlag
	} else {
		before = opt.CFlag
	}

	var res []string
	lastAdded := -1
	for _, num := range matchLineNums {
		low := 0
		if (num - before) > low {
			low = num - before
		}
		high := len(lines) - 1
		if (num + after) < high {
			high = num + after
		}
		for i := low; i <= high; i++ {
			if i > lastAdded {
				if opt.NFlag {
					res = append(res, fmt.Sprintf("%d %s", i+1, lines[i]))
				} else {
					res = append(res, lines[i])
				}
				lastAdded = i
			}
		}
	}
	return res
}

func match(lines []string, regExp *regexp.Regexp, opt *cmdutils.Options) []int {
	var res []int
	for i, line := range lines {
		if regExp.MatchString(line) != opt.VFlag {
			res = append(res, i)
		}
	}
	return res
}
