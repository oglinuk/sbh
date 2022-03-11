package sbh

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getSBH(algo, pt, symbols string, nr, seed int64, ut, l int) SBH {
	return SBH{
		Plaintext:      pt,
		NRots:          nr,
		Seed:           seed,
		Algorithm:      algo,
		UppercaseTimes: ut,
		Symbols:        symbols,
		Length:         l,
	}
}

// TODO: Should Generate return a valid string or should it return ""?
func TestBadInputSBH(t *testing.T) {
	expected := "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
	actual := Generate(SBH{})
	assert.Equal(t, expected, actual)
}

func TestDefaultGenerate(t *testing.T) {
	expected := "505c3c3fab111175642a3073756472621d0312388dec72ee9268c05f96463ecb"
	actual := Generate(getSBH("", "test", "", 1729, 42, 0, 0))
	assert.Equal(t, expected, actual)
}

func TestSymbolsGenerate(t *testing.T) {
	expected := "505c3!@#$%^&" //c3fab111175642a3073756472621d0312388dec72ee9268c05f96463ecb!@#$%^&"
	actual := Generate(getSBH("", "test", "!@#$%^&", 1729, 42, 0, 12))
	assert.Equal(t, expected, actual)
}

func TestUppercaseTimesGenerate(t *testing.T) {
	expected := "505C3C3Fab111" //175642a3073756472621d0312388dec72ee9268c05f96463ecb"
	actual := Generate(getSBH("", "test", "", 1729, 42, 3, 13))
	assert.Equal(t, expected, actual)
}

func TestSHA256_224Generate(t *testing.T) {
	expected := "8abc1cdbbc47d3" //e3876f86fccd0715a956177ea181249621f90a6222"
	actual := Generate(getSBH("sha256_224", "test", "", 1729, 42, 0, 14))
	assert.Equal(t, expected, actual)
}

func TestSHA256Generate(t *testing.T) {
	expected := "505c3c3fab11117" //5642a3073756472621d0312388dec72ee9268c05f96463ecb"
	actual := Generate(getSBH("sha256", "test", "", 1729, 42, 0, 15))
	assert.Equal(t, expected, actual)
}

func TestSHA512_224Generate(t *testing.T) {
	expected := "62ff96f2cab0dc8c" //e549b0ba9cefe079532bb5143897f4ec7b4b66c9"
	actual := Generate(getSBH("sha512_224", "test", "", 1729, 42, 0, 16))
	assert.Equal(t, expected, actual)
}

func TestSHA512_256Generate(t *testing.T) {
	expected := "814e45a6755086fe0" //9ac2f52c1412154fa6d5048ff26b27b9393b20149243a9a"
	actual := Generate(getSBH("sha512_256", "test", "", 1729, 42, 0, 17))
	assert.Equal(t, expected, actual)
}

func TestSHA512_384Generate(t *testing.T) {
	expected := "58014bc206308721b4" //575ca55800be7a4c54ca7f8c65a35d5b1271dee7e9ce89eaf750eca6af4d4b95aa32448a412490"
	actual := Generate(getSBH("sha512_384", "test", "", 1729, 42, 0, 18))
	assert.Equal(t, expected, actual)
}

func TestSHA512Generate(t *testing.T) {
	expected := "8be80abeb9fa101f7e6" //7b4d7da5ce00b1914e32d361eb67627d6fd4fd85cadf13a0bb079332508f443445a8ce410c1664e9bfcfcd86e6c7ca47cf47d234d7d49"
	actual := Generate(getSBH("sha512", "test", "", 1729, 42, 0, 19))
	assert.Equal(t, expected, actual)
}

func benchmarkDefaultGenerate(nRots int64, b *testing.B) {
	for i := 0; i < b.N; i++ {
		Generate(getSBH("", "test", "", nRots, 42, 0, 0))
	}
}

func BenchmarkDefaultGenerate100Rots(b *testing.B) {
	benchmarkDefaultGenerate(100, b)
}

func BenchmarkDefaultSBH1000Rots(b *testing.B) {
	benchmarkDefaultGenerate(1000, b)
}

func BenchmarkDefaultSBH10000Rots(b *testing.B) {
	benchmarkDefaultGenerate(10000, b)
}

func BenchmarkDefaultSBH100000Rots(b *testing.B) {
	benchmarkDefaultGenerate(100000, b)
}

func BenchmarkDefaultSBH1000000Rots(b *testing.B) {
	benchmarkDefaultGenerate(1000000, b)
}

func BenchmarkDefaultSBH10000000Rots(b *testing.B) {
	benchmarkDefaultGenerate(10000000, b)
}

func BenchmarkDefaultSBH100000000Rots(b *testing.B) {
	benchmarkDefaultGenerate(100000000, b)
}
