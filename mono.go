// Copyright (C) 2014 Space Monkey, Inc.

// Package monotime provides a monotonic timer for Go 1.2
// Go 1.3 will support monotonic time on its own.
package monotime

import (
	"time"
)

// Monotonic returns a time duration from some fixed point in the past.
func Monotonic() time.Duration {
	sec, nsec := monotime()
	return time.Duration(sec*1000000000 + int64(nsec))
}
