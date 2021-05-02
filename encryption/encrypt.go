package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

// Encrypts text and returns a base64 encoded text
func Encrypt(text string) string {
	plaintext := []byte(text)

	block, err := aes.NewCipher([]byte(os.Getenv("encryptionKey")))
	if err != nil {
		logrus.Fatalf("first error: %s\n", err)
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		logrus.Fatalf("second error: %s\n", err)
	}
	ciphertext := aesgcm.Seal(nonce, nonce, plaintext, nil)
	return base64.StdEncoding.EncodeToString(ciphertext)
}
