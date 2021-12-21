package u

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

// HmacSHA1 hmac sha1
func HmacSHA1(salt string, data string) string {
	hash := hmac.New(sha1.New, []byte(salt))
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}

// HmacSHA256 hmac sha256
func HmacSHA256(salt string, data string) string {
	hash := hmac.New(sha256.New, []byte(salt))
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}

// HmacSHA512 hmac sha512
func HmacSHA512(salt string, data string) string {
	hash := hmac.New(sha512.New, []byte(salt))
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}
