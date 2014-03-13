// Copyright (C) 2014 Space Monkey, Inc.

package _cgo

/*
#include <time.h>
#include <stdint.h>

#ifdef __MACH__
#include <mach/clock.h>
#include <mach/mach.h>
#endif

int nano_count(struct timespec *ts) {
#ifdef __MACH__
  uint64_t nanoseconds = mach_absolute_time();
  ts->tv_sec = nanoseconds / 1000000000;
  ts->tv_nsec = nanoseconds % 1000000000;
  return 0;
#else
#ifdef CLOCK_MONOTONIC_RAW
  if (clock_gettime(CLOCK_MONOTONIC_RAW, ts) < 0) {
#endif
  if (clock_gettime(CLOCK_MONOTONIC, ts) < 0) {
    return -1;
  }
#ifdef CLOCK_MONOTONIC_RAW
  }
#endif
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
