package sbh

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getSBH(algorithm, plaintext string, nrots, seed int64) SBH {
	return SBH{
		Plaintext: plaintext,
		NRots: nrots,
		Seed: seed,
		Algorithm: algorithm,
		Uppercase: false,
		UppercaseTimes: 0,
		Symbols: "",
	}
}

func TestDefaultSBH(t *testing.T) {
	result := Generate(getSBH("", "test", 1729, 42))
	expected := "505c3c3fab111175642a3073756472621d0312388dec72ee9268c05f96463ecb"

	assert.Equal(t, expected, result)
}

func TestMD5SBH(t *testing.T) {
	result := Generate(getSBH("md5", "test", 1729, 42))
	expected := "57b830ea86331cc8f669202fad3d1850"

	assert.Equal(t, expected, result)
}

func TestSHA1SBH(t *testing.T) {
	result := Generate(getSBH("sha1", "test", 1729, 42))
	expected := "b1add2443ddd007f04484c3736b340c9db743a55"

	assert.Equal(t, expected, result)
}

func TestSHA256_224SBH(t *testing.T) {
	result := Generate(getSBH("sha256_224", "test", 1729, 42))
	expected := "8abc1cdbbc47d3e3876f86fccd0715a956177ea181249621f90a6222"

	assert.Equal(t, expected, result)
}

func TestSHA256SBH(t *testing.T) {
	result := Generate(getSBH("sha256", "test", 1729, 42))
	expected := "505c3c3fab111175642a3073756472621d0312388dec72ee9268c05f96463ecb"

	assert.Equal(t, expected, result)
}

func TestSHA512_224SBH(t *testing.T) {
	result := Generate(getSBH("sha512_224", "test", 1729, 42))
	expected := "62ff96f2cab0dc8ce549b0ba9cefe079532bb5143897f4ec7b4b66c9"

	assert.Equal(t, expected, result)
}

func TestSHA512_256SBH(t *testing.T) {
	result := Generate(getSBH("sha512_256", "test", 1729, 42))
	expected := "814e45a6755086fe09ac2f52c1412154fa6d5048ff26b27b9393b20149243a9a"

	assert.Equal(t, expected, result)
}

func TestSHA512_384SBH(t *testing.T) {
	result := Generate(getSBH("sha512_384", "test", 1729, 42))
	expected := "58014bc206308721b4575ca55800be7a4c54ca7f8c65a35d5b1271dee7e9ce89eaf750eca6af4d4b95aa32448a412490"

	assert.Equal(t, expected, result)
}

func TestSHA512SBH(t *testing.T) {
	result := Generate(getSBH("sha512", "test", 1729, 42))
	expected := "8be80abeb9fa101f7e67b4d7da5ce00b1914e32d361eb67627d6fd4fd85cadf13a0bb079332508f443445a8ce410c1664e9bfcfcd86e6c7ca47cf47d234d7d49"

	assert.Equal(t, expected, result)
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
