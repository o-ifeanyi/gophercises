package quiz

import (
	"testing"
)

func TestParseContent(t *testing.T) {
	sample := [][]string{
		{"5 + 5", "10"},
		{"5 + 2", "7"},
	}
	expected := []Quiz{
		{Question: "5 + 5", Answer: "10"},
		{Question: "5 + 2", Answer: "7"},
	}

	res := ParseContent(sample, true)

	for i, v := range res {
		if !contains(expected, v) {
			t.Error("Got", v, "Expected", expected[i])
		}
	}
}

func TestContains(t *testing.T) {
	sample := []Quiz{
		{Question: "5 + 5", Answer: "10"},
		{Question: "5 + 2", Answer: "7"},
	}
	exp1 := sample[1]
	exp2 := Quiz{Question: "20 + 43", Answer: "63"}

	if !contains(sample, exp1) {
		t.Error("sample should contain", exp1)
	}

	if contains(sample, exp2) {
		t.Error("sample should NOT contain", exp2)
	}
}

func BenchmarkParseContent(b *testing.B) {
	sample := [][]string{
		{"5 + 5", "10"},
		{"5 + 2", "7"},
	}
	for i := 0; i < b.N; i++ {
		ParseContent(sample, true)
	}
}

func BenchmarkContains(b *testing.B) {
	sample := []Quiz{
		{Question: "5 + 5", Answer: "10"},
		{Question: "5 + 2", Answer: "7"},
	}
	for i := 0; i < b.N; i++ {
		contains(sample, Quiz{Question: "5 + 5", Answer: "10"})
	}
}
