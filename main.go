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
  nCiphers := getNCiphers()
  plainText := getPlainText()
  for i := 0; i < nCiphers; i++ {
    rot := getRot()
    plainText = getSHA(getCipherText(plainText, rot))
  }
  fmt.Printf("%s\n", plainText)
}

// Caesar cipher function
func CaesarCipher(r rune, rot int) rune {
    c := int(r) + rot
    if c > 'z' {
        return rune(c - 26)
    } else if c < 'a' {
        return rune(c + 26)
    }
    return rune(c)
}

// Get and return caesar cipher text
func getCipherText(plainText string, rot int) string {
  cipherText := strings.Map(func(r rune) rune {
      return CaesarCipher(r, rot)
  }, plainText)
  return cipherText
}

// Get and return num of ciphers to be applied
func getNCiphers() int {
  var i int
  fmt.Println("Number of Ciphers: ")
  fmt.Scan(&i)
  return i
}

// Get and return plain text to be hashed
func getPlainText() string {
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Print("Plain text: ")
  scanner.Scan()
  return scanner.Text()
}

// Get and return cipher rotation
func getRot() int {
  var i int
  fmt.Println("Rotation: ")
  fmt.Scan(&i)
  return i
}

// Get and return SHA256
func getSHA(cipherText string) string {
    hasher := sha256.New()
    hasher.Write([]byte(cipherText))
    return hex.EncodeToString(hasher.Sum(nil))
}
