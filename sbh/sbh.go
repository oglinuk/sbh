package sbh

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
	"hash/adler32"
	"hash/crc32"
	"hash/fnv"
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

func getHasher(algorithm string) hash.Hash {
	hashers := make(map[string]hash.Hash)
	hashers["sha1"] = sha1.New()
	hashers["md5"] = md5.New()
	hashers["adler32"] = adler32.New()
	hashers["crc32"] = crc32.NewIEEE()
	hashers["fnv32"] = fnv.New32()
	hashers["fnv32a"] = fnv.New32a()
	hashers["fnv64"] = fnv.New64()
	hashers["fnv64a"] = fnv.New64a()
	hashers["fnv128"] = fnv.New128()
	hashers["fnv128a"] = fnv.New128a()
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
func Generate(algorithm, plainText string, nRots, seed int64) string {
	rand.Seed(seed)
	for i := 0; i < int(nRots); i++ {
		rot := rand.Intn(math.MaxInt64)
		plainText = caesarCipher(rot, plainText)
	}
	hasher := getHasher(algorithm)
	hasher.Write([]byte(plainText))
	return hex.EncodeToString(hasher.Sum(nil))
}

// caesarCipher applies a (r)otation to each character of given (s)tring
func caesarCipher(r int, s string) string {
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
