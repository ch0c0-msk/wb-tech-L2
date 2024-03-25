package dev06

import (
	"dev06/cmdutils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCut(t *testing.T) {
	cases := []struct {
		name     string
		input    []string
		opt      *cmdutils.Options
		expected []string
	}{
		{
			"default delimeter",
			[]string{
				"first	second	third",
				"aaa;bbb;ccc",
			},
			&cmdutils.Options{
				FFlag: 2,
				DFlag: "\t",
				SFlag: false,
			},
			[]string{
				"second",
				"aaa;bbb;ccc",
			},
		},
		{
			"custom delimeter",
			[]string{
				"first	second	third",
				"aaa;bbb;ccc",
			},
			&cmdutils.Options{
				FFlag: 2,
				DFlag: ";",
				SFlag: false,
			},
			[]string{
				"first	second	third",
				"bbb",
			},
		},
		{
			"f flag is more than the number of fields",
			[]string{
				"first	second	third",
				"aaa;bbb;ccc",
			},
			&cmdutils.Options{
				FFlag: 10,
				DFlag: ";",
				SFlag: false,
			},
			[]string{
				"first	second	third",
				"",
			},
		},
		{
			"separated only",
			[]string{
				"first	second	third",
				"aaa;bbb;ccc",
			},
			&cmdutils.Options{
				FFlag: 2,
				DFlag: ";",
				SFlag: true,
			},
			[]string{
				"bbb",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := Cut(tc.input, tc.opt)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
