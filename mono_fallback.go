// Copyright (C) 2014 Space Monkey, Inc.
// +build !linux !amd64,!arm

package time

import (
    "code.spacemonkey.com/go/space/time/_cgo"
)

func monotime() (sec int64, nsec int32) {
    return _cgo.Monotime()
}
