package cipherx_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/tniah/x/utils/cipherx"
	"testing"
)

func TestEncryptAndDecrypt(t *testing.T) {
	key := []byte("PuyOd7UI1JyHI45a")
	aead, err := cipherx.New(key)
	assert.NoError(t, err)
	for k, plaintext := range []string{
		"this is my secret text",
		"-----BEGIN PUBLIC KEY-----\nMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCCm1e7FSSD3ZUSh/0gUlvGdxMz\nQySr4Ezx5w7Ygx0sZGclJltB7FIKHRlsXzDy0kVKNsMFyGDwxariGjRtgj/Hwew6\n4YyS/hfZgaPhiEUndAyeW2RW3VpvA5KcTjy3GiKIBTIy+TxIci1zeXfXeU+/vIzf\n+a8hnTF2P8sgfkww1QIDAQAB\n-----END PUBLIC KEY-----",
		"-----BEGIN RSA PRIVATE KEY-----\nMIICWwIBAAKBgQCCm1e7FSSD3ZUSh/0gUlvGdxMzQySr4Ezx5w7Ygx0sZGclJltB\n7FIKHRlsXzDy0kVKNsMFyGDwxariGjRtgj/Hwew64YyS/hfZgaPhiEUndAyeW2RW\n3VpvA5KcTjy3GiKIBTIy+TxIci1zeXfXeU+/vIzf+a8hnTF2P8sgfkww1QIDAQAB\nAoGAJrxFx8Gcg9N6+/UDGMv0VidItYJrZOJwT6pUl9hDFcBtavI2TJX3OvKocKDG\n1q2QSVN2gceNILuvU8Gr3PKtUXH5GXTYuZYJ3lUyhm0Pa7MkgQ3/doGXDKubPX+a\nrK56rPfn3SpX1aWY8TXacAi7qERtmUL+XWmeo4qa+8lvjn0CQQD+j2kUN9IGnyvC\nvzBcVcR6W6dyoCT3yVDkE23b7RbnpjikVxC3qhX6hpcPLrTfTLieJYT87czWgczr\n/Fr/y3ubAkEAg1h0URbEORUGG3+wxM9ihSdalDQMKW2vVc0PCUxea6D4uGQb1tkz\naPVp9shker+er8wnwinV9IGuIajjH3HkTwJAdq+usnqENgoogRhbF/H1NYdePxdj\npRP73xsf8ZZNQ5xAdH8TkE6BCNmPvMhuFF7VBQdBRhwpkSnbvXtfgjwBWQJAP8ad\nhBo34TeyJXwVCxtfzSPUuY2kMiGON21AVdV9K2mYG4CQe/wvGFHByBB5qZiNpvLM\ng1zpBLZLJRDqZ4RXxQJABnkjQqkQAs9SjdB/dGMYf5A7X4L5PUfW0v1C8PbIez5k\nuMDsDJ+0iMTP8ayl7MzrxXxnnHOalHHVGeT0TCmx7Q==\n-----END RSA PRIVATE KEY-----",
	} {
		ciphertext, err := aead.Encrypt([]byte(plaintext))
		assert.NoError(t, err, "Test case %d", k+1)
		decrypted, err := aead.Decrypt(ciphertext)
		assert.NoError(t, err, "Test case %d", k+1)
		assert.Equal(t, plaintext, string(decrypted), "Test case %d", k+1)
	}
}

func TestEncryptAndDecryptWithAdditionalData(t *testing.T) {
	key := []byte("o9LWDwz8MXKZlU1PBEMBDWklPBHBkqfs")
	aad := []byte("this is my context")
	aead, err := cipherx.New(key)
	assert.NoError(t, err)
	for k, plaintext := range []string{
		"this is my secret text",
		"-----BEGIN PUBLIC KEY-----\nMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCCm1e7FSSD3ZUSh/0gUlvGdxMz\nQySr4Ezx5w7Ygx0sZGclJltB7FIKHRlsXzDy0kVKNsMFyGDwxariGjRtgj/Hwew6\n4YyS/hfZgaPhiEUndAyeW2RW3VpvA5KcTjy3GiKIBTIy+TxIci1zeXfXeU+/vIzf\n+a8hnTF2P8sgfkww1QIDAQAB\n-----END PUBLIC KEY-----",
		"-----BEGIN RSA PRIVATE KEY-----\nMIICWwIBAAKBgQCCm1e7FSSD3ZUSh/0gUlvGdxMzQySr4Ezx5w7Ygx0sZGclJltB\n7FIKHRlsXzDy0kVKNsMFyGDwxariGjRtgj/Hwew64YyS/hfZgaPhiEUndAyeW2RW\n3VpvA5KcTjy3GiKIBTIy+TxIci1zeXfXeU+/vIzf+a8hnTF2P8sgfkww1QIDAQAB\nAoGAJrxFx8Gcg9N6+/UDGMv0VidItYJrZOJwT6pUl9hDFcBtavI2TJX3OvKocKDG\n1q2QSVN2gceNILuvU8Gr3PKtUXH5GXTYuZYJ3lUyhm0Pa7MkgQ3/doGXDKubPX+a\nrK56rPfn3SpX1aWY8TXacAi7qERtmUL+XWmeo4qa+8lvjn0CQQD+j2kUN9IGnyvC\nvzBcVcR6W6dyoCT3yVDkE23b7RbnpjikVxC3qhX6hpcPLrTfTLieJYT87czWgczr\n/Fr/y3ubAkEAg1h0URbEORUGG3+wxM9ihSdalDQMKW2vVc0PCUxea6D4uGQb1tkz\naPVp9shker+er8wnwinV9IGuIajjH3HkTwJAdq+usnqENgoogRhbF/H1NYdePxdj\npRP73xsf8ZZNQ5xAdH8TkE6BCNmPvMhuFF7VBQdBRhwpkSnbvXtfgjwBWQJAP8ad\nhBo34TeyJXwVCxtfzSPUuY2kMiGON21AVdV9K2mYG4CQe/wvGFHByBB5qZiNpvLM\ng1zpBLZLJRDqZ4RXxQJABnkjQqkQAs9SjdB/dGMYf5A7X4L5PUfW0v1C8PbIez5k\nuMDsDJ+0iMTP8ayl7MzrxXxnnHOalHHVGeT0TCmx7Q==\n-----END RSA PRIVATE KEY-----",
		"07DijfhAvt+L+OQTsKb7Ue4kJCJae9lpHQlFF9nTEQQ",
	} {
		ciphertext, err := aead.Encrypt([]byte(plaintext), aad)
		assert.NoError(t, err, "Test case %d", k+1)
		decrypted, err := aead.Decrypt(ciphertext, aad)
		assert.NoError(t, err, "Test case %d", k+1)
		assert.Equal(t, plaintext, string(decrypted), "Test case %d", k+1)
	}
}
