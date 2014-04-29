// Copyright (C) 2014 Space Monkey, Inc.
// +build !linux !amd64,!arm

package monotime

import (
	"github.com/SpaceMonkeyGo/monotime/_cgo"
)

func monotime() (sec int64, nsec int32) {
	return _cgo.Monotime()
}
