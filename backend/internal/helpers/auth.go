package helpers

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/argon2"
)

func Hash(passwd string) (string, error) {

	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(passwd), salt, 1, 64*1024, 4, 32)

	encodedHash := base64.RawStdEncoding.EncodeToString(hash)
	encodedSalt := base64.RawStdEncoding.EncodeToString(salt)

	return encodedHash + encodedSalt, nil
}

func VerifyHash(passwd string, hash string) bool {

	decodedHash, errh := base64.RawStdEncoding.DecodeString(hash[:43])
	decodedSalt, errs := base64.RawStdEncoding.DecodeString(hash[43:])

	if errh != nil && errs != nil {
		return false
	}

	newHash := argon2.IDKey([]byte(passwd), decodedSalt, 1, 64*1024, 4, 32)

	if bytes.Equal(decodedHash, newHash) {
		return true
	}

	return false
}
