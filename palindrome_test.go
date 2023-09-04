package main

import (
	"fmt"
	"testing"
)

type testEntry struct {
	value string
}

var tableTest = []testEntry{
	{value: "aba"},
	{value: "1000"},
	{value: "a long text that is not a palindrome"},
	{value: "soccorammesubinoonibuosemmarrocos"},
}

func BenchmarkPalindromeA(b *testing.B) {
	for _, v := range tableTest {
		b.Run(fmt.Sprintf("Test %s", v.value), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				PalindromoA(v.value)
			}
		})
	}
}

func BenchmarkPalindromeB(b *testing.B) {
	for _, v := range tableTest {
		b.Run(fmt.Sprintf("Test %s", v.value), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				PalindromoB(v.value)
			}
		})
	}
}
