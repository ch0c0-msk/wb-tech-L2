package cmdutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRegExp(t *testing.T) {
	cases := []struct {
		name           string
		opt            *Options
		args           []string
		expectedRegExp string
	}{
		{
			"NoneFlag",
			&Options{
				AFlag:     0,
				BFlag:     0,
				CFlag:     0,
				CountFlag: false,
				IFlag:     false,
				VFlag:     false,
				FFlag:     false,
				NFlag:     false,
			},
			[]string{"somepattern"},
			"somepattern",
		},
		{
			"RegisterFlag",
			&Options{
				AFlag:     0,
				BFlag:     0,
				CFlag:     0,
				CountFlag: false,
				IFlag:     true,
				VFlag:     false,
				FFlag:     false,
				NFlag:     false,
			},
			[]string{"somepattern"},
			"(?i)somepattern",
		},
		{
			"FixedFlag",
			&Options{
				AFlag:     0,
				BFlag:     0,
				CFlag:     0,
				CountFlag: false,
				IFlag:     false,
				VFlag:     false,
				FFlag:     true,
				NFlag:     false,
			},
			[]string{"^somepattern"},
			"\\Q^somepattern\\E",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual, _, err := parseArgs(tc.args, tc.opt)
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, tc.expectedRegExp, actual)
		})
	}
}
