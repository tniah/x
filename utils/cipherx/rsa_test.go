package cipherx_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/tniah/x/utils/cipherx"
	"testing"
)

func TestGenerateRSAPemKeyPair(t *testing.T) {
	keyPair, err := cipherx.GenerateRSAPemKeyPair(1024)
	assert.NoError(t, err)
	fmt.Println(string(keyPair[0]))
	fmt.Println(string(keyPair[1]))
}
