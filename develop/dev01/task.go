package dev01

import (
	"time"

	"github.com/beevik/ntp"
)

// GetCurrentAndCorrectTime take current time from localhost and correct time from NTP
func GetCurrentAndCorrectTime() (time.Time, time.Time, error) {
	resp, err := ntp.Query("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		return time.Now(), time.Now(), err
	}
	currentTime := time.Now()
	correctTime := currentTime.Add(resp.ClockOffset)
	return currentTime, correctTime, nil
}
