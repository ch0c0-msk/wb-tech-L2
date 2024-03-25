package dev06

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestOrChannel(t *testing.T) {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	start := time.Now()
	<-OrChannel(
		sig(1*time.Second),
		sig(2*time.Second),
		sig(3*time.Second),
	)
	elapsed := time.Since(start)
	assert.Less(t, elapsed, 2*time.Second)
	assert.Greater(t, elapsed, 999*time.Millisecond)
}
