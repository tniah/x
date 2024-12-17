package cipherx

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

type AEAD interface {
	Encrypt(plaintext []byte, opts ...[]byte) ([]byte, error)
	Decrypt(ciphertext []byte, opts ...[]byte) ([]byte, error)
}

type aeadImpl struct {
	aesgcm cipher.AEAD
}

func New(key []byte) (AEAD, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	return &aeadImpl{aesgcm: aesgcm}, nil
}

func (im *aeadImpl) Encrypt(plaintext []byte, opts ...[]byte) ([]byte, error) {
	var additionalData []byte
	if len(opts) > 0 {
		additionalData = opts[0]
	}

	nonce := make([]byte, im.aesgcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ct := im.aesgcm.Seal(nil, nonce, plaintext, additionalData)
	return append(nonce, ct...), nil
}

func (im *aeadImpl) Decrypt(ciphertext []byte, opts ...[]byte) ([]byte, error) {
	var additionalData []byte
	if len(opts) > 0 {
		additionalData = opts[0]
	}

	nonce := ciphertext[:im.aesgcm.NonceSize()]
	ct := ciphertext[im.aesgcm.NonceSize():]

	return im.aesgcm.Open(nil, nonce, ct, additionalData)
}
