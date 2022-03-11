package sbh

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
	"math"
	"math/rand"
	"strings"
	"unicode"
)

// LETTERS, DIGITS, and SYMBOLS are the string representations of
// all possible letters, digits, and symbols (minus some symbols).
const (
	LETTERS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	DIGITS  = "0123456789"
	SYMBOLS = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"

	LLEN = int64(len(LETTERS))
	DLEN = int64(len(DIGITS))
	SLEN = int64(len(SYMBOLS))
)

// SBH is a struct containing all the required parts to create a SBH
type SBH struct {
	Plaintext      string
	NRots          int64
	Seed           int64
	Algorithm      string
	UppercaseTimes int
	Symbols        string
	Length         int
}

// getHasher returns the corresponding hash function for the given
// algorithm. If a non-existing algorithm is given, default to sha256.
func getHasher(algorithm string) hash.Hash {
	hashers := make(map[string]hash.Hash)
	hashers["sha256_224"] = sha256.New224()
	hashers["sha256"] = sha256.New()
	hashers["sha512_224"] = sha512.New512_224()
	hashers["sha512_256"] = sha512.New512_256()
	hashers["sha512_384"] = sha512.New384()
	hashers["sha512"] = sha512.New()

	var h hash.Hash

	if _, exists := hashers[algorithm]; !exists {
		h = hashers["sha256"]
	} else {
		h = hashers[algorithm]
	}

	return h
}

// caesarCipher applies a (r)otation to each rune of given (s)tring
func caesarCipher(r int64, s string) string {
	str := []string{}
	s = strings.TrimSpace(s)

	for _, runeVal := range s {
		iRune := int64(runeVal)
		if unicode.IsLetter(runeVal) {
			r = (iRune + r) % LLEN
			str = append(str, string(LETTERS[r]))
		} else if unicode.IsDigit(runeVal) {
			r = (iRune + r) % DLEN
			str = append(str, string(DIGITS[r]))
		} else {
			r = (iRune + r) % SLEN
			str = append(str, string(SYMBOLS[r]))
		}
	}

	return strings.Join(str, "")
}

// Generate sets the rng seed with SBH.Seed, then applies a caesarCipher
// for the number of times specified by SBH.NRots. The resulting string
// is then trimmed to the specified SBH.Length. Replace each letter with
// an uppercase value for SBH.UppercaseTimes. Finally append any
// SBH.Symbols and return the result.
func Generate(secbaehash SBH) string {
	rand.Seed(secbaehash.Seed)
	for i := 0; i < int(secbaehash.NRots); i++ {
		rot := rand.Int63n(math.MaxInt64)
		secbaehash.Plaintext = caesarCipher(rot, secbaehash.Plaintext)
	}
	hasher := getHasher(secbaehash.Algorithm)
	hasher.Write([]byte(secbaehash.Plaintext))
	hash := hex.EncodeToString(hasher.Sum(nil))

	if secbaehash.Length == 0 {
		secbaehash.Length = len(hash)
	}

	secbaehash.Length -= len(secbaehash.Symbols)

	hash = hash[:secbaehash.Length]

	c := 0
	for _, r := range hash {
		if c >= secbaehash.UppercaseTimes {
			break
		}

		if unicode.IsLetter(r) {
			hash = strings.Replace(hash, string(r), strings.ToUpper(string(r)), 1)
			c++
		}
	}

	if secbaehash.Symbols != "" {
		hash += secbaehash.Symbols
	}

	return hash
}
