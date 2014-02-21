// Copyright (C) 2014 Space Monkey, Inc.

package time

import (
    "time"
)

func Monotonic() time.Duration {
    sec, nsec := monotime()
    return time.Duration(sec * 1000000000 + int64(nsec))
}
