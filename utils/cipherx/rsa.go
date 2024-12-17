package cipherx

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

func GenerateRSAPemKeyPair(bitSize uint32) ([]string, error) {
	prv, err := rsa.GenerateKey(rand.Reader, int(bitSize))
	if err != nil {
		return nil, err
	}

	prvPem := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(prv),
	})

	pubPem := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(&prv.PublicKey),
	})

	return []string{string(prvPem), string(pubPem)}, nil
}
