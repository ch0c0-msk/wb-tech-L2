package dev03

import (
	"dev03/cmdutils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	cases := []struct {
		name            string
		opt             *cmdutils.Options
		input           []string
		expected        []string
		isErrorExpected bool
	}{
		{
			"NoneFlag",
			&cmdutils.Options{KFlag: 1, NFlag: false, RFlag: false, UFlag: false},
			[]string{"B string", "C string", "A string"},
			[]string{"A string", "B string", "C string"},
			false,
		},
		{
			"KeyFlag",
			&cmdutils.Options{KFlag: 2, NFlag: false, RFlag: false, UFlag: false},
			[]string{"It`s B string", "It`s C string", "It`s A string"},
			[]string{"It`s A string", "It`s B string", "It`s C string"},
			false,
		},
		{
			"ReverseFlag",
			&cmdutils.Options{KFlag: 1, NFlag: false, RFlag: true, UFlag: false},
			[]string{"B string", "C string", "A string"},
			[]string{"C string", "B string", "A string"},
			false,
		},
		{
			"UniqueFlag",
			&cmdutils.Options{KFlag: 1, NFlag: false, RFlag: false, UFlag: true},
			[]string{"B string", "C string", "A string", "B string"},
			[]string{"A string", "B string", "C string"},
			false,
		},
		{
			"NumFlagWithNumValues",
			&cmdutils.Options{KFlag: 1, NFlag: true, RFlag: false, UFlag: false},
			[]string{"2.75 string", "25.1256 string", "1.25 string"},
			[]string{"1.25 string", "2.75 string", "25.1256 string"},
			false,
		},
		{
			"NumFlagWithNotNumValues",
			&cmdutils.Options{KFlag: 1, NFlag: true, RFlag: false, UFlag: false},
			[]string{"B string", "C string", "A string"},
			[]string{},
			true,
		},
		{
			"KeyGreaterThanColumns",
			&cmdutils.Options{KFlag: 4, NFlag: false, RFlag: false, UFlag: false},
			[]string{"B string", "C string", "A string"},
			[]string{"A string", "B string", "C string"},
			true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ctx, err := NewSortContext(tc.opt, tc.input)
			if err != nil && !tc.isErrorExpected {
				t.Error(err)
			} else if err != nil && tc.isErrorExpected {
				return
			} else if err == nil && tc.isErrorExpected {
				t.Fail()
			} else {
				actual := ctx.Sort()
				assert.Equal(t, tc.expected, actual)
			}
		})
	}
}
