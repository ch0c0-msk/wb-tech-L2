package dev06

import (
	"dev06/cmdutils"
	"strings"
)

func Cut(lines []string, opt *cmdutils.Options) []string {
	var res []string
	for _, line := range lines {
		fields := strings.Split(line, opt.DFlag)
		if len(fields) == 1 && !opt.SFlag {
			res = append(res, line)
		} else if len(fields) != 1 && opt.FFlag <= len(fields) {
			res = append(res, fields[opt.FFlag-1])
		} else if len(fields) != 1 && opt.FFlag > len(fields) {
			res = append(res, "")
		}
	}
	return res
}
