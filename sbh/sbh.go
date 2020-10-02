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

const (
	LETTERS = "abcdefghijklmnopqrstuvwxyz"
	DIGITS  = "0123456789"
	SYMBOLS = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
)

type SBH struct {
	Plaintext      string
	NRots          int64
	Seed           int64
	Algorithm      string
	Uppercase      bool
	UppercaseTimes int
	Symbols        string
}

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
		h = hashers["sha256"] // Default to using sha256 if non-existing algorithm given
	} else {
		h = hashers[algorithm]
	}

	return h
}

// Generate a hash using a caesarCipher with the specified
// hashing algorithm. A pseudo-random rot is generated based
// on given seed for number of rotations (nRots) specified.
func Generate(secbaehash SBH) string {
	rand.Seed(secbaehash.Seed)
	for i := 0; i < int(secbaehash.NRots); i++ {
		rot := rand.Intn(math.MaxInt64)
		secbaehash.Plaintext = caesarCipher(rot, secbaehash.Plaintext)
	}
	hasher := getHasher(secbaehash.Algorithm)
	hasher.Write([]byte(secbaehash.Plaintext))
	hash := hex.EncodeToString(hasher.Sum(nil))

	// TODO: Change a rune to uppercase based on seed or rotation given rather than the first rune that IsLetter
	if secbaehash.Uppercase {
		for _, r := range hash {
			if unicode.IsLetter(r) {
				hash = strings.Replace(hash, string(r), strings.ToUpper(string(r)), secbaehash.UppercaseTimes)
				break
			}
		}
	}

	// TODO: Either figure out why certain combinations cause errors or change how to get symbols
	// Also need to figure out how to append the symbols in pseudo-random (same based on seed) places
	if secbaehash.Symbols != "" {
		hash += secbaehash.Symbols
	}

	return hash
}

// caesarCipher applies a (r)otation to each character of given (s)tring
func caesarCipher(r int, s string) string {
	str := []string{}
	s = strings.TrimSpace(s)

	for _, runeVal := range s {
		if unicode.IsLetter(runeVal) {
			r = (int(runeVal) + r) % len(LETTERS)
			str = append(str, string(LETTERS[r]))
		} else if unicode.IsDigit(runeVal) {
			r = (int(runeVal) + r) % len(DIGITS)
			str = append(str, string(DIGITS[r]))
		} else {
			r = (int(runeVal) + r) % len(SYMBOLS)
			str = append(str, string(SYMBOLS[r]))
		}
	}

	return strings.Join(str, "")
}
