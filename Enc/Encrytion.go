package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

//Initialized Vector or Alphabets ==== "Your combination of keys returns error because they don't the same byte encoding with the alphabets"
var IV = []byte{00, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

// Please keep this key secret
var secretKey string = "£*%&><!)\"(@abcdefghijklmnopqrst"

// Encoding here
func encodeBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

//Decoding here
func decodeBase64(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

//Encryption: This takes your words and hides it from external or third party
func Encryption(text, secretKey string) (string, error) {
	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return "", err
	}
	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, IV)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return encodeBase64(cipherText), nil
}

//Decryption: This takes your recieved Secret keys and unhides the texts
func Decryption(text, secretKey string) (string, error) {
	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return "", err
	}
	cipherText := decodeBase64(text)
	cfb := cipher.NewCFBDecrypter(block, IV)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}

func main() {
	fmt.Println("==========================================================")
	fmt.Println("Welcome to the encryption and Decryption Arena")
	fmt.Println("==========================================================")

	var choice int
	fmt.Println("Press the number of your choice, \n" +
		"1. Encrypt a text \n" +
		"2. Decrypt your text")
	fmt.Scanln(&choice)
	fmt.Println("==========================================================")

	switch choice {
	case 1:
		//Encrypting texts
		var phrase string
		fmt.Println("What do you want to encrypt?: ")
		fmt.Scanln(&phrase)

		encText, err := Encryption(phrase, secretKey)
		if err != nil {
			errMessage := "error in encoding your text: "
			fmt.Print(errMessage, err)

		}
		fmt.Println("==========================================================")
		fmt.Println("Your encrypted message key is: ", encText)
		fmt.Println("==========================================================")

	case 2:
		var decryptionKey string
		fmt.Println("What do you want to decrypt?: ")
		fmt.Scanln(&decryptionKey)

		decText, err := Decryption(decryptionKey, secretKey)
		if err != nil {
			errMessage := "error in decoding your text: "
			fmt.Print(errMessage, err)
		}
		fmt.Println("==========================================================")
		fmt.Println("Your encrypted message key is: ", decText)
		fmt.Println("==========================================================")

	default:
		fmt.Println("Wrong choice!!!")
	}
}
