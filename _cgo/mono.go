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

struct timespec nano_count();
*/
import "C"

//export cpanic
func cpanic() {
	panic("failed getting monotonic clock")
}

func Monotime() (sec int64, nsec int32) {
	rv := C.nano_count()
	return int64(rv.tv_sec), int32(rv.tv_nsec)
}
