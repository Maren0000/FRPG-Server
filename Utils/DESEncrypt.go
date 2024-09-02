package Utils

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"errors"
	"fmt"
)

const (
	KEY = "FrPg-109"
	IV  = "sX5nJEQL"
)

func DESDecrypt(cipherText []byte) (plainText []byte, err error) {
	DecryptData := make([]byte, len(cipherText))
	block, err := des.NewCipher([]byte(KEY))
	if err != nil {
		return DecryptData, err
	}
	blockMode := cipher.NewCBCDecrypter(block, []byte(IV))

	blockMode.CryptBlocks(DecryptData, cipherText)
	DecryptData, err = pkcs7strip(DecryptData, des.BlockSize)
	if err != nil {
		return DecryptData, err
	}

	return DecryptData, nil
}

func DESEncrypt(plainText []byte) (cipherText []byte) {
	block, err := des.NewCipher([]byte(KEY))
	if err != nil {
		fmt.Println(err)
	}
	blockSize := block.BlockSize()
	origData, err := pkcs7pad(plainText, blockSize)
	if err != nil {
		fmt.Println(err)
	}
	blockMode := cipher.NewCBCEncrypter(block, []byte(IV))
	encrypted := make([]byte, len(origData))
	blockMode.CryptBlocks(encrypted, origData)

	return encrypted
}

// pkcs7strip remove pkcs7 padding
func pkcs7strip(data []byte, blockSize int) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("pkcs7: Data is empty")
	}
	if length%blockSize != 0 {
		return nil, errors.New("pkcs7: Data is not block-aligned")
	}
	padLen := int(data[length-1])
	ref := bytes.Repeat([]byte{byte(padLen)}, padLen)
	if padLen > blockSize || padLen == 0 || !bytes.HasSuffix(data, ref) {
		return nil, errors.New("pkcs7: Invalid padding")
	}
	return data[:length-padLen], nil
}

// pkcs7pad add pkcs7 padding
func pkcs7pad(data []byte, blockSize int) ([]byte, error) {
	if blockSize <= 1 || blockSize >= 256 {
		return nil, fmt.Errorf("pkcs7: Invalid block size %d", blockSize)
	} else {
		padLen := blockSize - len(data)%blockSize
		padding := bytes.Repeat([]byte{byte(padLen)}, padLen)
		return append(data, padding...), nil
	}
}
