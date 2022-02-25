package main

import (
	"strings"
	"testing"
)

func TestGenerateEmail(t *testing.T) {
	email := GenerateEmail("jonathan@example.com")
	if !strings.Contains(email, "jonathan+") {
		t.Fatalf("email contains right base email")
	}
}

func BenchmarkGenerateEmail(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GenerateEmail("jonathan@example.com")
	}
}
