package dev08

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsePipes(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected []Command
	}{
		{
			"one command",
			"cd ./somedir",
			[]Command{
				{
					Name: "cd",
					Args: []string{"./somedir"},
				},
			},
		},
		{
			"multiple commands",
			"pwd | cd ./somedir | echo aaa bbb ccc",
			[]Command{
				{
					Name: "pwd",
				},
				{
					Name: "cd",
					Args: []string{"./somedir"},
				},
				{
					Name: "echo",
					Args: []string{"aaa", "bbb", "ccc"},
				},
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := ParsePipes(tc.input)
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, len(tc.expected), len(actual))
			for i := range tc.expected {
				assert.True(t, reflect.DeepEqual(tc.expected[i], actual[i]))
			}
		})
	}
}
