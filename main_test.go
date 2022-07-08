package main

import "testing"

func TestGreeting(t *testing.T) {
	msg := getMessage()
	expected := "Welcome to deploybot with GitHub Actions!"
	if msg != expected {
		t.Errorf("expected %v, got %v", expected, msg)
	}
}
