package main

import (
	"crypto/des"
	"crypto/cipher"
	"encoding/base64"
	"bytes"
	"fmt"
)

var APP_KEY  = []byte("5be461e8386d572bc835d02059ab8973")
var iv         = []byte("01234567")

// 填充
func PKCS5Padding(text []byte, blockSize int) []byte {
	padding := blockSize - len(text)%blockSize
	addText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(text, addText...)
}

// 反填充
func PKCS5UnPadding(text []byte) []byte {
	length := len(text)
	cut := int(text[length-1])
	return text[:(length - cut)]
}


// 加密
// app_key长度必须24位，否则引发panic
func Encrypt(data string) string {
	block, _ := des.NewTripleDESCipher(APP_KEY[:24])
	eData := PKCS5Padding([]byte(data), block.BlockSize())
	rst := make([]byte, len(eData))
	blockMode := cipher.NewCBCEncrypter(block, iv)
	blockMode.CryptBlocks(rst, eData)
	return base64.StdEncoding.EncodeToString(rst)
}

// 解密
func Decrypt(data string) ([]byte, error) {
	eData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	rst := make([]byte, len(eData))
	block, err := des.NewTripleDESCipher(APP_KEY[:24])
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	blockMode.CryptBlocks(rst, eData)
	rst = PKCS5UnPadding(rst)
	return rst, nil
}

func main() {
	var i string;
	i = "123"
	fmt.Println(Encrypt(i))
	c := "LnTwkcJTqKQXkmo7JK4WlIhlPX1WmWzveJV8knvH5J5BaurNdwbyHHpQ87/H lhq9cOrqygaxMSY6YZiggqam+YvQdwTTeu16ImkyNjh7esXU1s7FZ7eCNUIq WY6HetQV/9XqgH89Nn3yFL7ZysZcnQybHuhMWDSIjtAA4qsJmEfwvp1HCkD8 1SBZ6ZHISLnQM/sUE3M4SC53IGOl2jcKvvc6Q3d/nWPVKAcINfDY1CKbtdhU zLkmjOk/o0Z+3d+CqkKsYS418qZkHwt+0C8FJJd6OEwxP5LX"
	a, err := Decrypt(c)
	if err != nil{
		fmt.Println(a)
	}
}
