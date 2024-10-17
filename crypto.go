package wxm

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
	dst = PKCS7UnPad(dst)
	return dst, nil
}

func PKCS7UnPad(data []byte) []byte {
	var length = len(data)
	var unpadding = int(data[length-1])
	if length < unpadding {
		return nil
	}
	return data[:(length - unpadding)]
}
