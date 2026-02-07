package internal

import (
	"crypto/aes"
	"crypto/cipher"
)

func AESCBCDecrypt(ciphertext, key, iv []byte) ([]byte, error) {
	var block, err = aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	var blockSize = block.BlockSize()
	iv = iv[:blockSize]

	var dst = make([]byte, len(ciphertext))

	var mode = cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(dst, ciphertext)
	dst = PKCS7Unpad(dst, blockSize)
	return dst, nil
}

func PKCS7Unpad(data []byte, blockSize int) []byte {
	length := len(data)
	if length == 0 {
		return nil
	}

	unpadding := int(data[length-1])

	if unpadding == 0 || unpadding > blockSize || unpadding > length {
		return nil
	}

	for i := 0; i < unpadding; i++ {
		if data[length-1-i] != byte(unpadding) {
			return nil
		}
	}

	return data[:(length - unpadding)]
}
