package u

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// getDerivedKey 获取派生密钥
func getDerivedKey(password string, salt []byte, count int) ([]byte, []byte) {
	key := md5.Sum([]byte(password + string(salt)))
	for i := 0; i < count - 1; i++ {
		key = md5.Sum(key[:])
	}
	return key[:8], key[8:]
}

// PBEEncrypt 加密
func PBEEncrypt(password string, iterations int, plainText string, salt []byte) (string, error) {
	padNum := byte(8 - len(plainText) % 8)
	for i := byte(0); i < padNum; i++ {
		plainText += string(padNum)
	}

	dk, iv := getDerivedKey(password, salt, iterations)

	block,err := des.NewCipher(dk)

	if err != nil {
		return "", err
	}

	enc := cipher.NewCBCEncrypter(block, iv)
	encrypted := make([]byte, len(plainText))
	enc.CryptBlocks(encrypted, []byte(plainText))

	return hex.EncodeToString(encrypted), nil
}

// PBEDecrypt 解密
func PBEDecrypt(password string, iterations int, cipherText string, salt []byte) (string, error) {
	msgBytes, err := hex.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	dk, iv := getDerivedKey(password, salt, iterations)
	block, err := des.NewCipher(dk)

	if err != nil {
		return "", err
	}

	dec := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(msgBytes))
	dec.CryptBlocks(decrypted, msgBytes)

	decryptedString := strings.TrimRight(string(decrypted), "\x01\x02\x03\x04\x05\x06\x07\x08")
	return decryptedString, nil
}
