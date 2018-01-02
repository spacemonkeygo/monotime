#include <time.h>
#include <stdint.h>
#include "_cgo_export.h"

#ifdef __MACH__
#include <mach/mach.h>
#include <mach/mach_time.h>
#endif

struct timespec nano_count() {
  struct timespec ts;
#ifdef __MACH__
  static mach_timebase_info_data_t info;
  if (info.denom == 0) {
    mach_timebase_info(&info);
  }
  uint64_t something = mach_absolute_time();
  something /= info.denom;
  something *= info.numer;
  ts.tv_sec = something / 1000000000;
  ts.tv_nsec = something % 1000000000;
  return ts;
#else
#ifdef CLOCK_MONOTONIC_RAW
  if (clock_gettime(CLOCK_MONOTONIC_RAW, &ts) < 0) {
#endif
  if (clock_gettime(CLOCK_MONOTONIC, &ts) < 0) {
    cpanic();
  }
#ifdef CLOCK_MONOTONIC_RAW
  }
#endif
#endif
  return ts;
}
