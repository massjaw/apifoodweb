package util

import (
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io"

	"github.com/dgryski/go-camellia"
	"github.com/spf13/viper"
)

// Fungsi untuk mengenkripsi plaintext menggunakan Camellia
func EncryptCamellia(plaintext string) (string, error) {

	// Membuat cipher Camellia dengan kunci
	c, err := camellia.New([]byte(getSalt()))
	if err != nil {
		SystemLog("Encrypt CAMELLIA", "Error create cipher block", err).Error()
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		SystemLog("Encrypt CAMELLIA", "Error create new GCM", err).Error()
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		SystemLog("Encrypt CAMELLIA", "Error in ReadFull", err).Error()
		return "", err
	}

	value := gcm.Seal(nonce, nonce, []byte(plaintext), nil)

	// Encode ciphertext menjadi string base64
	return base64.StdEncoding.EncodeToString(value), nil
}

// Fungsi untuk mendekripsi ciphertext menggunakan Camellia
func DecryptCamellia(encryptedString string) (string, error) {

	// Decode ciphertext dari base64
	ciphertext, err := base64.StdEncoding.DecodeString(encryptedString)
	if err != nil {
		SystemLog("Decrypt CAMELLIA", "Error when turn secret to cipher", err).Error()
		return "", err
	}

	// Membuat cipher Camellia dengan kunci
	c, err := camellia.New([]byte(getSalt()))
	if err != nil {
		SystemLog("Encrypt CAMELLIA", "Error create cipher block", err).Error()
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		SystemLog("Encrypt CAMELLIA", "Error create new GCM", err).Error()
		return "", err
	}
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		SystemLog("Encrypt CAMELLIA", "Error in nonce validation", err).Error()
		return "", errors.New("Error in ValidateNonceSize")
	}
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		SystemLog("Encrypt CAMELLIA", "Error while open GCM", err).Error()
		return "", err
	}

	return string(plaintext), nil
}

func getSalt() string {

	encSalt := viper.GetString("Encryption.enc_salt")

	saltBase64, err := base64.StdEncoding.DecodeString(encSalt)
	if err != nil {
		SystemLog("getSalt Encryption", "Error in decode base64", err).Error()
		return ""
	}

	bytedSalt, err := hex.DecodeString(reverseString(string(saltBase64)))
	if err != nil {
		SystemLog("getSalt Encryption", "Error in decode hex", err).Error()
		return ""
	}

	hexBase64, err := base64.StdEncoding.DecodeString(string(bytedSalt))
	if err != nil {
		SystemLog("getSalt Encryption", "Error in decode base64", err).Error()
		return ""
	}

	return string(hexBase64)
}

func reverseString(s string) string {

	runes := []rune(s)
	reversed := make([]rune, len(runes))

	for i, r := range runes {
		reversed[len(runes)-1-i] = r
	}

	return string(reversed)
}
