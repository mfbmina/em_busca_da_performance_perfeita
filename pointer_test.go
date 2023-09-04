package main

import (
	"testing"
)

func BenchmarkReferencePassMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReferencePassMain()
	}
}

func BenchmarkValuePassMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ValuePassMain()
	}
}
