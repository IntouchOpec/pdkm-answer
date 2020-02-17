package main

import (
	"testing"
)

func TestAnswerSix(t *testing.T) {
	given := []string{"ABCDEZXY", "abcdezxy", "AbCdEzXy", "A1b2C3d4/+*"}
	wants := []string{CapitalResult, SmallResult, MixResult, InResult}
	var result string
	for index, text := range given {
		result = AnswerSix(text)
		if result != wants[index] {
			t.Errorf("given %s wants %s but get %s\n", given[index], wants[index], result)
		}
	}
}

func benchmarkAnswerSix(text string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		AnswerSix(text)
	}
}

func benchmarkAnswerSeven(text string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		AnswerSeven(text)
	}
}
func BenchmarkAnswerSeven1(b *testing.B) { benchmarkAnswerSeven("dddd", b) }
func BenchmarkAnswerSeven2(b *testing.B) { benchmarkAnswerSeven("world", b) }
func BenchmarkAnswerSeven3(b *testing.B) { benchmarkAnswerSeven("cycle", b) }
func BenchmarkAnswerSeven4(b *testing.B) { benchmarkAnswerSeven("abba", b) }
func BenchmarkAnswerSix1(b *testing.B)   { benchmarkAnswerSix("ABCDEZXY", b) }
func BenchmarkAnswerSix2(b *testing.B)   { benchmarkAnswerSix("abcdezxy", b) }
func BenchmarkAnswerSix3(b *testing.B)   { benchmarkAnswerSix("AbCdEzXy", b) }
func BenchmarkAnswerSix4(b *testing.B)   { benchmarkAnswerSix("A1b2C3d4/+*", b) }
