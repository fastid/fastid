package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	b64 "encoding/base64"
	"encoding/hex"
	"io"
)

type Crypto interface {
	Encrypt(stringToEncrypt string) (encryptedBytes []byte, err error)
	EncryptBase64(stringToEncrypt string) (encryptedString string, err error)
	Decrypt(encryptedBytes []byte) (decryptedString string, err error)
	DecryptBase64(encryptedString string) (decryptedString string, err error)
}

type crypto struct {
	signatureKey string
}

func New(signatureKey string) Crypto {
	return &crypto{signatureKey: signatureKey}
}

func GenerateCipher() (string, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(key), nil
}

func (c *crypto) Encrypt(stringToEncrypt string) (encryptedBytes []byte, err error) {

	//Since the key is in string, we need to convert decode it to bytes
	key, _ := hex.DecodeString(c.signatureKey)
	plaintext := []byte(stringToEncrypt)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Create a new GCM - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	//Create a nonce. Nonce should be from GCM
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	//Encrypt the data using aesGCM.Seal
	//Since we don't want to save the nonce somewhere else in this case, we add it as a prefix to the encrypted data. The first nonce argument in Seal is the prefix.
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)

	return ciphertext, nil
}

func (c *crypto) EncryptBase64(stringToEncrypt string) (encryptedString string, err error) {
	encrypt, err := c.Encrypt(stringToEncrypt)
	if err != nil {
		return "", err
	}
	return b64.StdEncoding.EncodeToString([]byte(encrypt)), nil
}

func (c *crypto) Decrypt(encryptedBytes []byte) (decryptedString string, err error) {

	key, err := hex.DecodeString(c.signatureKey)
	if err != nil {
		return "", err
	}

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	//Create a new GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	//Get the nonce size
	nonceSize := aesGCM.NonceSize()

	//Extract the nonce from the encrypted data
	nonce, ciphertext := encryptedBytes[:nonceSize], encryptedBytes[nonceSize:]

	//Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func (c *crypto) DecryptBase64(encryptedString string) (decryptedString string, err error) {
	enc, err := b64.StdEncoding.DecodeString(encryptedString)
	if err != nil {
		return "", err
	}
	decrypt, err := c.Decrypt(enc)
	if err != nil {
		return "", err
	}
	return decrypt, nil

}
