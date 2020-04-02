package binaryconversion

import "testing"

func TestCheckNumber(t *testing.T) {
	got1, got2 := checkNumber("10000")
	want1, want2 := 16, true
	if got1 != want1 || got2 != want2 {
		t.Errorf("got1 %d want1 %d got1 %v want2 %v", got1, want1, got2, want2)
	}
}

//基准测试
func BenchmarkCheckNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checkNumber("10000")

	}
}
