package sbh

import (
	"crypto/sha256"
	"encoding/hex"
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

// Generate applies a caesarCipher with a random rot
// based on given seed for nRots and returns a sha256 hash
func Generate(plainText string, nRots, seed int64) string {
	rand.Seed(seed)
	for i := 0; i < int(nRots); i++ {
		rot := rand.Intn(math.MaxInt64)
		plainText = caesarCipher(plainText, rot)
	}
	hasher := sha256.New()
	hasher.Write([]byte(plainText))
	return hex.EncodeToString(hasher.Sum(nil))
}

// caesarCipher applies a (r)otation to each character of given (s)tring
func caesarCipher(s string, r int) string {
	encryptedStr := []string{}
	s = strings.TrimSpace(s)

	for _, runeVal := range s {
		if unicode.IsLetter(runeVal) {
			r = (int(runeVal) + r) % len(LETTERS)
			encryptedStr = append(encryptedStr, string(LETTERS[r]))
		} else if unicode.IsDigit(runeVal) {
			r = (int(runeVal) + r) % len(DIGITS)
			encryptedStr = append(encryptedStr, string(DIGITS[r]))
		} else {
			r = (int(runeVal) + r) % len(SYMBOLS)
			encryptedStr = append(encryptedStr, string(SYMBOLS[r]))
		}
	}

	return strings.Join(encryptedStr, "")
}
