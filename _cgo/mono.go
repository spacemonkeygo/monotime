// Copyright (C) 2014 Space Monkey, Inc.

package _cgo

/*
#cgo LDFLAGS: -lrt
#include <time.h>
#include <stdint.h>

int nano_count(struct timespec *ts) {
#ifdef CLOCK_MONOTONIC_RAW
  if (clock_gettime(CLOCK_MONOTONIC_RAW, ts) < 0) {
#endif
  if (clock_gettime(CLOCK_MONOTONIC, ts) < 0) {
    return -1;
  }
#ifdef CLOCK_MONOTONIC_RAW
  }
#endif
  return 0;
}
*/
import "C"

func Monotime() (sec int64, nsec int32) {
    var rv C.struct_timespec
    if C.nano_count(&rv) != 0 {
        panic("failed getting monotonic clock")
    }
    return int64(rv.tv_sec), int32(rv.tv_nsec)
}
