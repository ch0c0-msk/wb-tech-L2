package dev03

import (
	"dev03/cmdutils"
	"errors"
	"sort"
	"strconv"
	"strings"
)

// SortContext is common Context for sorting
type SortContext struct {
	opt     *cmdutils.Options
	strSort *StringSort
	numSort *NumSort
}

// NewSortContext creates SortContext with given options
func NewSortContext(opt *cmdutils.Options, lines []string) (*SortContext, error) {
	ctx := &SortContext{opt: opt}
	if opt.UFlag {
		lines = deleteNoneUniqueLines(lines)
	}
	strSort, err := NewStringSort(opt.KFlag-1, lines)
	ctx.strSort = strSort
	if err != nil {
		return nil, err
	}
	numSort, err := NewNumSort(opt.KFlag-1, lines)
	ctx.numSort = numSort
	if err != nil {
		return nil, err
	}
	if opt.NFlag && !isKeyColumnsNumbers(numSort.lines, numSort.key) {
		return nil, errors.New("for n flag key columns should be numbers")
	}
	return ctx, nil
}

func isKeyColumnsNumbers(lines []string, key int) (res bool) {
	res = true
	for _, line := range lines {
		words := strings.Fields(line)
		_, err := strconv.ParseFloat(words[key], 64)
		if err != nil {
			res = false
		}
	}
	return
}

func deleteNoneUniqueLines(lines []string) []string {
	var resLines []string
	uniqueLines := make(map[string]bool)
	for _, line := range lines {
		if !uniqueLines[line] {
			uniqueLines[line] = true
			resLines = append(resLines, line)
		}
	}
	return resLines
}

// Sort returns ordered lines
func (ctx *SortContext) Sort() []string {
	var sortAlg sort.Interface
	sortAlg = ctx.strSort
	if ctx.opt.NFlag {
		sortAlg = ctx.numSort
	}
	sort.Sort(sortAlg)
	if ctx.opt.RFlag {
		sort.Sort(sort.Reverse(sortAlg))
	}
	if ctx.opt.NFlag {
		return ctx.numSort.lines
	}
	return ctx.strSort.lines
}

// Checker for key and columns count
type Checker interface {
	isKeyCorrect(key int, lines []string) bool
}

// Check is Checker implementation
type Check struct{}

func (c *Check) isKeyCorrect(key int, lines []string) bool {
	isCorrect := true
	for _, line := range lines {
		wordsCount := len(strings.Fields(line))
		if key > wordsCount {
			isCorrect = false
			break
		}
	}
	return isCorrect
}

// StringSort is Sort interface implementations
type StringSort struct {
	lines []string
	key   int
	Checker
}

// NewStringSort creates StringSort with given strings and key column
func NewStringSort(key int, lines []string) (*StringSort, error) {
	strSort := &StringSort{key: key, lines: lines, Checker: &Check{}}
	if isCorrect := strSort.isKeyCorrect(key, lines); !isCorrect {
		return strSort, errors.New("invalid key, there are not so many words in one of the lines")
	}
	return strSort, nil
}

func (s *StringSort) Len() int {
	return len(s.lines)
}

func (s *StringSort) Less(i, j int) bool {
	iKeyColumn := strings.Fields(s.lines[i])[s.key]
	jKeyColumn := strings.Fields(s.lines[j])[s.key]
	return iKeyColumn < jKeyColumn
}

func (s *StringSort) Swap(i, j int) {
	s.lines[i], s.lines[j] = s.lines[j], s.lines[i]
}

// NumSort is Sort interface implementations
type NumSort struct {
	lines []string
	key   int
	Checker
}

// NewNumSort creates NumSort with given strings and key column
func NewNumSort(key int, lines []string) (*NumSort, error) {
	numSort := &NumSort{key: key, lines: lines, Checker: &Check{}}
	if isCorrect := numSort.isKeyCorrect(key, lines); !isCorrect {
		return numSort, errors.New("invalid key, there are not so many words in one of the lines")
	}
	return numSort, nil
}

func (n *NumSort) Len() int {
	return len(n.lines)
}

func (n *NumSort) Less(i, j int) bool {
	iKeyColumn := strings.Fields(n.lines[i])[n.key]
	iKeyNumColumn, _ := strconv.ParseFloat(iKeyColumn, 64)
	jKeyColumn := strings.Fields(n.lines[j])[n.key]
	jKeyNumColumn, _ := strconv.ParseFloat(jKeyColumn, 64)
	return iKeyNumColumn < jKeyNumColumn
}

func (n *NumSort) Swap(i, j int) {
	n.lines[i], n.lines[j] = n.lines[j], n.lines[i]
}
