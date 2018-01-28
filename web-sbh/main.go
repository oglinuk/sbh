package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
	"net/http"
	"encoding/hex"
	"crypto/sha256"
)

func main() {
	http.HandleFunc("/", serveSBH)
	fmt.Println("Starting server ...")
	http.ListenAndServe(":8080", nil)
}

func serveSBH(w http.ResponseWriter, r *http.Request) {
	query := strings.Split(r.FormValue("q"), " ")
	nCiphers, err := strconv.Atoi(query[0])
	if err != nil {
		return
	}
	plainText := query[1]
	for i := 0; i < nCiphers; i++ {
		rot, err := strconv.Atoi(query[i+2])
		if err != nil {
			return
		}
		plainText = getSBH(plainText, rot)
	}
	fmt.Fprintln(w, "SaltBaeHash:", plainText)
}

// Caesar cipher function
// Takes a char(rune) and a rotation -> applies a shift(rotation) -> returns the char(rune)
// NOTE: Need to figure out how to include alphanumerical values and special chars
func CaesarCipher(r rune, rot int) rune {
  c := int(r) + rot
  if c > 'z' {
      return rune(c - 26)
  } else if c < 'a' {
      return rune(c + 26)
  }
  return rune(c)
}

// Get and return the plainText to be hashed
func getPlainText() string {
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Print("Plain text: ")
  scanner.Scan()
  return scanner.Text()
}

// One-off util func for getting user input of int variables -> takes the context of the variable
func getIntValue(context string) int {
  var i int
  fmt.Printf("%s: ", context)
  fmt.Scan(&i)
  return i
}

// SaltBaeHashing function
// Takes plainText and a rotation value -> applies the caesar cipher -> returns the SHA256
func getSBH(plainText string, rot int) string {
  cipherText := strings.Map(func(r rune) rune {
      return CaesarCipher(r, rot)
  }, plainText)
  hasher := sha256.New()
  hasher.Write([]byte(cipherText))
  return hex.EncodeToString(hasher.Sum(nil))
}
