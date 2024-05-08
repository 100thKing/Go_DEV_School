//go:build day
// +build day

package main

import (
	"time"
)

func init() {
	CurrentTime := time.Now()
	day := CurrentTime.Weekday()
	day_str := day.String()
	message = (day_str)
}
