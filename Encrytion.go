package main

import (
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
}

func main() {
	fmt.Println("Hello")
}

func decrypt() {
	fmt.Print("Decrypt here")
}
