// Copyright (C) 2014 Space Monkey, Inc.

package time

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
