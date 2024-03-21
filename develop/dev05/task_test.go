package dev05

import (
	"dev05/cmdutils"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrep(t *testing.T) {
	cases := []struct {
		name     string
		lines    []string
		regExp   *regexp.Regexp
		opt      *cmdutils.Options
		expected []string
	}{
		{
			"NoneFlags",
			[]string{"golang is effective language", "java is ...", "c++ is ..."},
			regexp.MustCompile("golang"),
			&cmdutils.Options{
				AFlag:     0,
				BFlag:     0,
				CFlag:     0,
				CountFlag: false,
				IFlag:     false,
				VFlag:     false,
				FFlag:     false,
				NFlag:     false,
			},
			[]string{"golang is effective language"},
		},
		{
			"AfterFlag",
			[]string{"golang is effective language", "java is ...", "c++ is ..."},
			regexp.MustCompile("golang"),
			&cmdutils.Options{
				AFlag:     1,
				BFlag:     0,
				CFlag:     0,
				CountFlag: false,
				IFlag:     false,
				VFlag:     false,
				FFlag:     false,
				NFlag:     false,
			},
			[]string{"golang is effective language", "java is ..."},
		},
		{
			"BeforeFlag",
			[]string{"golang is effective language", "java is ...", "c++ is ..."},
			regexp.MustCompile("java"),
			&cmdutils.Options{
				AFlag:     0,
				BFlag:     1,
				CFlag:     0,
				CountFlag: false,
				IFlag:     false,
				VFlag:     false,
				FFlag:     false,
				NFlag:     false,
			},
			[]string{"golang is effective language", "java is ..."},
		},
		{
			"ContextFlag",
			[]string{"golang is effective language", "java is ...", "c++ is ..."},
			regexp.MustCompile("java"),
			&cmdutils.Options{
				AFlag:     0,
				BFlag:     0,
				CFlag:     1,
				CountFlag: false,
				IFlag:     false,
				VFlag:     false,
				FFlag:     false,
				NFlag:     false,
			},
			[]string{"golang is effective language", "java is ...", "c++ is ..."},
		},
		{
			"ContextFlagWithOutOfRange",
			[]string{"golang is effective language", "java is ...", "cpp is ..."},
			regexp.MustCompile("cpp"),
			&cmdutils.Options{
				AFlag:     0,
				BFlag:     0,
				CFlag:     100,
				CountFlag: false,
				IFlag:     false,
				VFlag:     false,
				FFlag:     false,
				NFlag:     false,
			},
			[]string{"golang is effective language", "java is ...", "cpp is ..."},
		},
		{
			"CountFlag",
			[]string{"golang is effective language", "java is ...", "cpp is ..."},
			regexp.MustCompile("java"),
			&cmdutils.Options{
				AFlag:     0,
				BFlag:     0,
				CFlag:     0,
				CountFlag: true,
				IFlag:     false,
				VFlag:     false,
				FFlag:     false,
				NFlag:     false,
			},
			[]string{"1"},
		},
		{
			"InvertFlag",
			[]string{"golang is effective language", "java is ...", "cpp is ..."},
			regexp.MustCompile("java"),
			&cmdutils.Options{
				AFlag:     0,
				BFlag:     0,
				CFlag:     0,
				CountFlag: false,
				IFlag:     false,
				VFlag:     true,
				FFlag:     false,
				NFlag:     false,
			},
			[]string{"golang is effective language", "cpp is ..."},
		},
		{
			"NumberFlag",
			[]string{"golang is effective language", "java is ...", "cpp is ..."},
			regexp.MustCompile("java"),
			&cmdutils.Options{
				AFlag:     0,
				BFlag:     0,
				CFlag:     0,
				CountFlag: false,
				IFlag:     false,
				VFlag:     false,
				FFlag:     false,
				NFlag:     true,
			},
			[]string{"2 java is ..."},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := Grep(tc.lines, tc.regExp, tc.opt)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
