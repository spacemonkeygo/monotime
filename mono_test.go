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

package monotime

import (
	"runtime"
	"testing"
	"time"
)

func TestMonotonic(t *testing.T) {
	upper := 10
	if runtime.GOOS == "darwin" {
		// a bug only showed up when we did this test for seconds on darwin
		upper = 10000
	}
	for i := 0; i < upper; i++ {
		start := Monotonic()
		time.Sleep(1 * time.Millisecond)
		end := Monotonic()
		if end < start+1*time.Millisecond {
			t.Fatalf("time didn't advance by 1 millisecond? %s", end-start)
		}
	}
}

func BenchmarkMonotonic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Monotonic()
	}
}

func BenchmarkTimeNow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		time.Now()
	}
}
