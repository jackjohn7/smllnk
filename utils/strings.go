package utils

import (
	crand "crypto/rand"
	"encoding/base64"
	"math/rand"
)

const SessionIdLength int = 64 // length in bytes

var alphabet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

/*
Cryptographically INSECURE. Used for generating IDs that are not sensitive
*/
func RandomString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(b)
}

/*
Cryptographically secure way of generating bytes for use as secrets
*/
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := crand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

/*
Generates a base-64 encoded session id
*/
func GenerateSessionId() (string, error) {
	bytes, err := GenerateRandomBytes(SessionIdLength)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
}

/*
Generates a base-64 encoded session id
*/
func GenerateMagicLinkId() (string, error) {
	bytes, err := GenerateRandomBytes(22)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
}
