package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
)

func main() {

	testRsa()

}

func testRsa() {

	bitSize := 2048

	privateKey, err := rsa.GenerateKey(rand.Reader, bitSize)

	if err != nil {
		panic(err)
	}

	err = privateKey.Validate()
	if err != nil {
		panic(err)
	}

	publicKey := &privateKey.PublicKey

	bytesPrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
	stringPrivateKey := base64.StdEncoding.EncodeToString(bytesPrivateKey)

	bytesPublicKey := x509.MarshalPKCS1PublicKey(publicKey)
	stringPublicKey := base64.StdEncoding.EncodeToString(bytesPublicKey)

	log.Println("PRIVATE KEY")
	log.Println(stringPrivateKey)

	decodedBytesPrivateKey, err := base64.StdEncoding.DecodeString(stringPrivateKey)
	if err != nil {
		panic(err)
	}

	log.Println(bytes.Equal(bytesPrivateKey, decodedBytesPrivateKey))

	log.Println("PUBLIC KEY")
	log.Println(stringPublicKey)

	decodedBytesPublicKey, err := base64.StdEncoding.DecodeString(stringPublicKey)
	if err != nil {
		panic(err)
	}

	decodedPublicKey, err := x509.ParsePKCS1PublicKey(decodedBytesPublicKey)
	if err != nil {
		panic(err)
	}

	log.Println(bytes.Equal(bytesPublicKey, decodedBytesPublicKey))

	t := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iss": "auth.ivixlabs.com",
		"sub": "userId",
		"foo": 1,
	})

	stringToken, err := t.SignedString(privateKey)
	if err != nil {
		panic(err)
	}

	log.Println("TOKEN")
	log.Println(stringToken)

	token, err := jwt.Parse(stringToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return decodedPublicKey, nil
	})
	if err != nil {
		panic(err)
	}

	log.Println("DECODED TOKEN")
	log.Println(token)

}

func testHmac() {
	key := []byte("It is signing key")
	//wrongKey := []byte("It is signing key wrong")

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "auth.ivixlabs.com",
		"sub": "userId",
		"foo": 1,
	})

	s, err := t.SignedString(key)

	if err != nil {
		panic(err)
	}

	log.Println(s)

	token, err := jwt.Parse(s, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return key, nil
	})

	if err != nil {
		panic(err)
	}

	log.Println(token)
}
