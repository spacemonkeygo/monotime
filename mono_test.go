// Copyright (C) 2014 Space Monkey, Inc.

package time

import (
	"testing"
	"time"
)

func TestMonotonic(t *testing.T) {
	for i := 0; i < 10; i++ {
		start := Monotonic()
		time.Sleep(1 * time.Millisecond)
		if Monotonic() < start+1*time.Millisecond {
			t.Fatal("time didn't advance by 1 millisecond?")
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
