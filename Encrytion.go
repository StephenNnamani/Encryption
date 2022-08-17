package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

//Initialized Vector or Alphabets
var alphabets = []byte("a b c d e f g h i j k l m n o p q r s t u v w x y z")

// Please keep this key secret
var secretKey string = "£ * % & > < ! ) \" ( @ a b c d e f g h i j k l m n o"

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
	cfb := cipher.NewCFBEncrypter(block, alphabets)
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
	cfb := cipher.NewCFBDecrypter(block, alphabets)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}

func main() {
	fmt.Println("Hello")
}
