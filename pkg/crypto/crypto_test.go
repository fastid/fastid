package crypto

import (
	"github.com/stretchr/testify/require"
	"testing"
)

const cipherKey = "b5226304d593078fb00c8e63f1649420ee0f12d2308c2fca42f0a73ec7e35c88"

func TestCryptoCipher(t *testing.T) {
	cipher, err := GenerateCipher()
	require.NoError(t, err)
	require.NotEmpty(t, cipher)
}

func TestCryptEncryptDecrypt(t *testing.T) {
	crypto := New(cipherKey)
	cryptString, err := crypto.EncryptBase64("Hello word!")
	require.NoError(t, err)
	require.NotEmpty(t, cryptString)

	deCryptString, err := crypto.DecryptBase64(cryptString)
	require.NoError(t, err)
	require.Equal(t, deCryptString, "Hello word!")
}
