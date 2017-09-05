package main

import (
	"fmt"
	//"os"
	//"path/filepath"
	//"encoding/base64"
	//"regexp"
	"encoding/base64"
	"crypto/des"
	"crypto/cipher"
	"bytes"
)

var (
	iv         = []byte("01234567")
	APP_KEY	= "7de1a69814ce94eff83bb3ae9dde3ad5"
)

func byteToString(p []byte) string{
	for i:= 0; i < len(p); i++{
		if p[i] == 0{
			return string(p[0:i])
		}
	}
	return string(p)
}

func main() {
	//fmt.Println("hello world")
	//pwd, _ := os.Getwd()
	//fmt.Printf("this is %s\n", pwd)
	//WorkDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	//fmt.Println(WorkDir, err)
	//var data = "MTExMTEx"
	//data1 := []byte("111111")
	//eData := base64.StdEncoding.EncodeToString(data1)
	//fmt.Println(eData)

	tmpo1 := "tC6j%2FEH2P23Sav7C2CPdKiAuOwXcfVhJQVo1" +
		"ICg%2FOLFas%2BC1R3wKwMPH9Yuw+0uvVUZwI7Q8goi6" +
		"tuuSHILvlXGA%2BRviZR%2FmZ59oTAV5kvvhCRS2DC9z9Dl7%2B+KIElyOycV5lKzo" +
		"jhHOpXaSkVQsKQVMc18T52LPgnI6fCMUmA%2FV3jZCEtY0jv+XwKwma2iMR4CFhXRNv%2ByGQ0%3D"
	data, err := Decrypt(tmpo1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(byteToString(data))
}

func Decrypt(data string) ([]byte, error) {
	eData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	rst := make([]byte, len(eData))
	APP_KEY1 := []byte(APP_KEY)
	block, err := des.NewTripleDESCipher(APP_KEY1[:24])
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	blockMode.CryptBlocks(rst, eData)
	rst = PKCS5UnPadding(rst)
	return rst, nil
}

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