package main

import (
	"math/rand"
	"testing"
	"time"
)

func TestAlwaysPass(t *testing.T) {
	t.Log("TestAlwaysPass: This test always passes")
}

func TestAlwaysSkip(t *testing.T) {
	t.Skip()
	t.Error("TestAlwaysSkip: This test always skips")
}

func TestAlwaysFail(t *testing.T) {
	t.Skip()
	t.Error("TestAlwaysFail: This test always fails")
}

func TestRandomFail(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(100) > 50 {
		t.Log("TestRandomFail: This test passed")
	} else {
		t.Error("TestRandomFail: This test failed")
	}
}
