package bdcrypto

import (
	"crypto/cipher"
	"crypto/des"
	"errors"
)

// DESCBCEncrypt3 实现3DES加密, CBC模式
func DESCBCEncrypt3(plaintext, key, iv []byte) (ciphertext []byte, err error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}

	defer func() {
		if rerr := recover(); rerr != nil {
			err = errors.New(rerr.(string))
		}
	}()

	bs := block.BlockSize()
	plaintext = PKCS5Padding(plaintext, bs)
	blockMode := cipher.NewCBCEncrypter(block, iv)

	ciphertext = make([]byte, len(plaintext))
	blockMode.CryptBlocks(ciphertext, plaintext)
	return ciphertext, nil
}

// DESCBCDecrypt3 实现3DES解密, CBC模式
func DESCBCDecrypt3(ciphertext, key, iv []byte) (plaintext []byte, err error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}

	defer func() {
		if rerr := recover(); rerr != nil {
			err = errors.New(rerr.(string))
		}
	}()

	blockMode := cipher.NewCBCDecrypter(block, iv)

	plaintext = make([]byte, len(ciphertext))
	blockMode.CryptBlocks(plaintext, ciphertext)
	plaintext = PKCS5UnPadding(plaintext)
	return plaintext, nil
}
