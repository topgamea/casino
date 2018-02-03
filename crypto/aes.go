package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

const (
//ivDefValue = "0102030405060708"
)

func AesEncrypt(plaintext []byte, key []byte) ([]byte, error) {
	lt := len(key)
	if lt != 32 {
		return nil, errors.New("token invalid")
	}
	key[0] = '0'
	key[lt-1] = '1'
	block, err := aes.NewCipher(key[0:16])
	if err != nil {
		return nil, errors.New("invalid decrypt key")
	}
	blockSize := block.BlockSize()
	plaintext = PKCS5Padding(plaintext, blockSize)

	blockMode := cipher.NewCBCEncrypter(block, key[16:])

	ciphertext := make([]byte, len(plaintext))
	blockMode.CryptBlocks(ciphertext, plaintext)

	dl := base64.StdEncoding.EncodedLen(len(ciphertext))
	dst := make([]byte, dl)
	base64.StdEncoding.Encode(dst, ciphertext)

	return dst, nil
}

func AesDecrypt(ciphertext []byte, key []byte) ([]byte, error) {
	lt := len(key)
	if lt != 32 {
		return nil, errors.New("token invalid")
	}
	key[0] = '0'
	key[lt-1] = '1'

	dl := base64.StdEncoding.DecodedLen(len(ciphertext))
	dafter64 := make([]byte, dl)
	n, err := base64.StdEncoding.Decode(dafter64, ciphertext)
	if err != nil {
		return nil, err
	}
	ciphertext = dafter64[0:n]

	block, err := aes.NewCipher(key[0:16])
	if err != nil {
		return nil, errors.New("invalid decrypt key")
	}

	blockSize := block.BlockSize()

	if len(ciphertext) < blockSize {
		return nil, errors.New("ciphertext too short")
	}

	//iv := []byte(ivDefValue)
	if len(ciphertext)%blockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}

	blockModel := cipher.NewCBCDecrypter(block, key[16:])

	plaintext := make([]byte, len(ciphertext))
	blockModel.CryptBlocks(plaintext, ciphertext)
	plaintext = PKCS5UnPadding(plaintext)

	return plaintext, nil
}

func PKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}
