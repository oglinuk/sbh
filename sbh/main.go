package main

import (
    "os"
    "fmt"
    "bufio"
    "strings"
    "crypto/sha256"
    "encoding/hex"
)

func main() {

}

// Initialize -> execute on start
func init() {
  print("\033[H\033[2J") // clear terminal
  nCiphers := getIntValue("Number of Ciphers")
  plainText := getPlainText()
  for i := 0; i < nCiphers; i++ {
    rot := getIntValue("Rotation: ")
    plainText = getSBH(plainText, rot)
  }
  print("\033[H\033[2J")
  fmt.Printf("SBH: %s\n", plainText)
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
