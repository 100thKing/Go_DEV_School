//go:build time
// +build time

package main

import (
	"time"
)

func init() {
	CurrentTime := time.Now()
	TimeFormat := CurrentTime.Format("15:04:05")
	message = (TimeFormat)
}
