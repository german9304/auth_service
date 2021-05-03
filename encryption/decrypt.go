package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

// Decrypt takes a base64 text and encrypts it
func Decrypt(text string) string {
	ciphertext, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		logrus.Fatalf("dcrypt first error: %s\n", err)
	}
	block, err := aes.NewCipher([]byte(os.Getenv("encryptionKey")))
	if err != nil {
		logrus.Fatalf("dcrypt second error: %s\n", err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		logrus.Fatalf("dcrypt third error: %s\n", err)
	}

	nonceSize := aesgcm.NonceSize()
	if len(ciphertext) < nonceSize {
		fmt.Println(err)
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		logrus.Fatalf("dcrypt four error: %s\n", err)
	}

	return string(plaintext)
}
