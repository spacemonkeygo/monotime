// Copyright (C) 2014 Space Monkey, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package _cgo

/*
#include <time.h>
#include <stdint.h>

#ifdef __MACH__
#include <mach/mach.h>
#include <mach/mach_time.h>
#endif

int nano_count(struct timespec *ts) {
#ifdef __MACH__
  static mach_timebase_info_data_t info;
  if (info.denom == 0) {
    mach_timebase_info(&info);
  }
  uint64_t something = mach_absolute_time();
  something /= info.denom;
  something *= info.numer;
  ts->tv_sec = something / 1000000000;
  ts->tv_nsec = something % 1000000000;
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
