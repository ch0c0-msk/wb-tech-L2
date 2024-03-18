package dev02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnpackString(t *testing.T) {
	cases := []struct {
		name     string
		common   string
		expected string
	}{
		{
			"without changes",
			"abcd",
			"abcd",
		},
		{
			"default",
			"a4bc2d5e",
			"aaaabccddddde",
		},
		{
			"empty string",
			"",
			"",
		},
		{
			"escape digits",
			"qwe\\4\\5",
			"qwe45",
		},
		{
			"escape escape`s",
			"qwe\\\\5",
			"qwe\\\\\\\\\\",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := UnpackString(tc.common)
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, tc.expected, actual)
		})
	}
}
