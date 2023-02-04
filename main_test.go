package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestAlwaysPass(t *testing.T) {
	t.Log("TestAlwaysPass: This test always passes")
}

func TestAlwaysFail(t *testing.T) {
	t.Skip()
	t.Error("TestAlwaysFail: This test always fails")
}

func TestRandomFail(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(100)
	fmt.Println(i)
	if i > 50 {
		t.Log("TestRandomFail: This test passed")
	} else {
		t.Error("TestRandomFail: This test failed")
	}
}
