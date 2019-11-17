package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/realbucksavage/miniauth/lib"
	"strconv"
)

var bitsizeEnv = "RSA_BIT_SIZE"

func GenerateRSAPrivateKey() (*rsa.PrivateKey, error) {
	bitSize, _ := strconv.Atoi(lib.GetEnv(bitsizeEnv, "4096"))

	privateKey, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		return nil, err
	}

	err = privateKey.Validate()
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func EncodePrivateKeyToPem(privateKey *rsa.PrivateKey) []byte {
	privateDER := x509.MarshalPKCS1PrivateKey(privateKey)

	privateBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   privateDER,
	}

	privatePEM := pem.EncodeToMemory(&privateBlock)

	return privatePEM
}

func GenerateRSAPublicKey(publicKey *rsa.PublicKey) ([]byte, error) {
	pubKeyDer := x509.MarshalPKCS1PublicKey(publicKey)
	pubKeyBlock := pem.Block{
		Type:    "PUBLIC KEY",
		Headers: nil,
		Bytes:   pubKeyDer,
	}

	pubKeyBytes := pem.EncodeToMemory(&pubKeyBlock)
	return pubKeyBytes, nil
}
