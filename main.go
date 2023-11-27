package main

import (
	"fmt"
	"strings"
)

func main() {
	var key string
	var plaintext string

	fmt.Print("Enter the key (a single letter): ")
	fmt.Scanln(&key)
	key = strings.ToUpper(key)
	if len(key) != 1 || key < "A" || key > "Z" {
		fmt.Println("Invalid key. Please enter a single letter.")
		return
	}

	fmt.Print("Enter plaintext: ")
	fmt.Scanln(&plaintext)
	plaintext = strings.ToUpper(plaintext)

	encryptedText := encrypt(plaintext, key)
	fmt.Println("Encrypted text:", encryptedText)

	decryptedText := decrypt(encryptedText, key)
	fmt.Println("Decrypted text:", decryptedText)
}

func encrypt(plaintext string, key string) string {
	var encryptedText strings.Builder

	for _, char := range plaintext {
		if char >= 'A' && char <= 'Z' {
			shift := int(key[0] - 'A')
			encryptedChar := shiftChar(char, shift)
			encryptedText.WriteRune(encryptedChar)
		} else {
			encryptedText.WriteRune(char)
		}
	}

	return encryptedText.String()
}

func decrypt(ciphertext string, key string) string {
	var decryptedText strings.Builder

	for _, char := range ciphertext {
		if char >= 'A' && char <= 'Z' {
			shift := int(key[0] - 'A')
			decryptedChar := shiftChar(char, -shift)
			decryptedText.WriteRune(decryptedChar)
		} else {
			decryptedText.WriteRune(char)
		}
	}

	return decryptedText.String()
}

func shiftChar(char rune, shift int) rune {
	return 'A' + (char-'A'+rune(shift)+26)%26
}
