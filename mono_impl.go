// Copyright (C) 2014 Space Monkey, Inc.
// +build linux,amd64 linux,arm

package monotime

func monotime() (sec int64, nsec int32)
