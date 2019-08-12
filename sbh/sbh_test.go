package sbh

import "testing"

func TestSBH(t *testing.T) {
	result := Generate("test", 1729, 42)
	expected := "196a7f528702e5ca85cd0ac664843cfb4bdd615ce5bc384d60db65ee20a30fb2"

	if result != expected {
		t.Fatalf("Expected %s\nGot %v", expected, result)
	}
}

func benchmarkSBH(nRots int64, b *testing.B) {
	for i := 0; i < b.N; i++ {
		Generate("test", nRots, 42)
	}
}

func BenchmarkSBH100(b *testing.B) {
	benchmarkSBH(100, b)
}

func BenchmarkSBH1000(b *testing.B) {
	benchmarkSBH(1000, b)
}

func BenchmarkSBH10000(b *testing.B) {
	benchmarkSBH(10000, b)
}

func BenchmarkSBH100000(b *testing.B) {
	benchmarkSBH(100000, b)
}

func BenchmarkSBH1000000(b *testing.B) {
	benchmarkSBH(1000000, b)
}

func BenchmarkSBH10000000(b *testing.B) {
	benchmarkSBH(10000000, b)
}

func BenchmarkSBH100000000(b *testing.B) {
	benchmarkSBH(100000000, b)
}
