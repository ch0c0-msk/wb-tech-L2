package dev01

import (
	"testing"
	"time"
)

func TestGetCurrentAndCorrectTime(t *testing.T) {
	currentTime, correctTime, err := GetCurrentAndCorrectTime()
	if err != nil {
		t.Error(err)
	}
	diff := currentTime.Sub(correctTime)
	if diff > time.Second {
		t.Log("difference between local and correct times is too large")
		t.Fail()
	}
}
