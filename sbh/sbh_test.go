package sbh

import "testing"

func getSBH(algorithm, plaintext string, nrots, seed int64) SBH {
	return SBH{
		Plaintext:      plaintext,
		NRots:          nrots,
		Seed:           seed,
		Algorithm:      algorithm,
		Uppercase:      false,
		UppercaseTimes: 0,
		Symbols:        "",
	}
}

func TestDefaultSBH(t *testing.T) {
	result := Generate(getSBH("", "test", 1729, 42))
	expected := "196a7f528702e5ca85cd0ac664843cfb4bdd615ce5bc384d60db65ee20a30fb2"

	if result != expected {
		t.Fatalf("\nExpected\n%s\nGot\n%v", expected, result)
	}
}

func TestSHA256_224SBH(t *testing.T) {
	result := Generate(getSBH("sha256_224", "test", 1729, 42))
	expected := "622e0c3884770a643d71a96c257ccf61934bae32d0fc6203ce4ed9e4"

	if result != expected {
		t.Fatalf("\nExpected\n%s\nGot\n%v", expected, result)
	}
}

func TestSHA256SBH(t *testing.T) {
	result := Generate(getSBH("sha256", "test", 1729, 42))
	expected := "196a7f528702e5ca85cd0ac664843cfb4bdd615ce5bc384d60db65ee20a30fb2"

	if result != expected {
		t.Fatalf("\nExpected\n%s\nGot\n%v", expected, result)
	}
}

func TestSHA512_224SBH(t *testing.T) {
	result := Generate(getSBH("sha512_224", "test", 1729, 42))
	expected := "f88dc0ca340b51b7b08232a9c374e144e6d39ba460340190e42ef46b"

	if result != expected {
		t.Fatalf("\nExpected\n%s\nGot\n%v", expected, result)
	}
}

func TestSHA512_256SBH(t *testing.T) {
	result := Generate(getSBH("sha512_256", "test", 1729, 42))
	expected := "6a19b273eb219d0617b5e81aec263b84186aab22764b0d38890eda3868a4ba1f"

	if result != expected {
		t.Fatalf("\nExpected\n%s\nGot\n%v", expected, result)
	}
}

func TestSHA512_384SBH(t *testing.T) {
	result := Generate(getSBH("sha512_384", "test", 1729, 42))
	expected := "c3ebc2626671811786db37333b2cf27e960d221e3da10c8cd17935dbb2dc466988d8bde1e0f012209fada77af35abb76"

	if result != expected {
		t.Fatalf("\nExpected\n%s\nGot\n%v", expected, result)
	}
}

func TestSHA512SBH(t *testing.T) {
	result := Generate(getSBH("sha512", "test", 1729, 42))
	expected := "d47381726a00d872d9e79859761399d6e601ddd73b5857a571b4c695360f017429648953fc50c58d1cab781a5bfcc5fd8a3958b9d50c241f97195d88a080d6af"

	if result != expected {
		t.Fatalf("\nExpected\n%s\nGot\n%v", expected, result)
	}
}

func benchmarkDefaultSBH(nRots int64, b *testing.B) {
	for i := 0; i < b.N; i++ {
		Generate(getSBH("test", "", nRots, 42))
	}
}

func BenchmarkDefaultSBH100Rots(b *testing.B) {
	benchmarkDefaultSBH(100, b)
}

func BenchmarkDefaultSBH1000Rots(b *testing.B) {
	benchmarkDefaultSBH(1000, b)
}

func BenchmarkDefaultSBH10000Rots(b *testing.B) {
	benchmarkDefaultSBH(10000, b)
}

func BenchmarkDefaultSBH100000Rots(b *testing.B) {
	benchmarkDefaultSBH(100000, b)
}

func BenchmarkDefaultSBH1000000Rots(b *testing.B) {
	benchmarkDefaultSBH(1000000, b)
}

func BenchmarkDefaultSBH10000000Rots(b *testing.B) {
	benchmarkDefaultSBH(10000000, b)
}

func BenchmarkDefaultSBH100000000Rots(b *testing.B) {
	benchmarkDefaultSBH(100000000, b)
}
