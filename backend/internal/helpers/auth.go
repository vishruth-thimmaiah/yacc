package helpers

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	math_rand "math/rand"
	"strings"

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

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandUsername() string {
	result := make([]byte, 20)
	for i := range result {
		result[i] = letters[math_rand.Intn(len(letters))]
	}
	return string(result)
}

func MatchFileType(filetype string) bool {
	typeof := strings.Split(filetype, "/")[0]
	for _, filetypes := range []string{"image", "video", "audio"} {
		if typeof == filetypes {
			return true
		}
	}
	return false
}
