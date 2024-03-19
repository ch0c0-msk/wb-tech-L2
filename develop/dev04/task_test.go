package dev04

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAnagrams(t *testing.T) {
	cases := []struct {
		name     string
		input    []string
		expected map[string][]string
	}{
		{
			"default",
			[]string{"пятак", "тяпка", "пятка", "листок", "слиток", "столик"},
			map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
			},
		},
		{
			"setWithOneElement",
			[]string{"пятак", "тяпка", "пятка", "листок", "машина", "слиток", "столик"},
			map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
			},
		},
		{
			"ordered",
			[]string{"cbaa", "caab", "abac", "aabc"},
			map[string][]string{
				"cbaa": {"aabc", "abac", "caab", "cbaa"},
			},
		},
		{
			"differentRegister",
			[]string{"cbaa", "caab", "abac", "aabc", "AaBc"},
			map[string][]string{
				"cbaa": {"aabc", "abac", "caab", "cbaa"},
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := FindAnagrams(tc.input)
			t.Logf("\nExpected: %v\nActual: %v\n", tc.expected, actual)
			assert.True(t, reflect.DeepEqual(actual, tc.expected))
		})
	}
}
