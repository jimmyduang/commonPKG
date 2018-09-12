package common

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"strings"
)

/**
* 定义结构体
 */
type AES struct {
	keyStr string //加密的key
	IV     string //加密的IV
	ty     string //补位规则
}

/**
* 创建并实例化一个aes结构体
* @param string keyStr 密钥
* @param string IV
* @param string ty 补位算法
* @return *AES 返回结构体(对象)
 */
func SetAES(keyStr, IV, ty string) *AES {
	c := new(AES)
	c.keyStr = keyStr
	c.IV = IV
	c.ty = ty
	if len(c.IV) > 0 {
		if len(c.IV) != 16 {
			fmt.Println("加密长度IV必须16位")
		}
	}
	if len(c.ty) < 1 {
		c.ty = "pkcs5"
	}

	return c
}

/**
* 数据加密
* @param string	str 需要加密的字符串
* @return strintg 返回加密后的字符串
 */
func (c *AES) AesEncryptString(str string) string {
	if len(str) < 1 {
		return ""
	}
	key := make([]byte, 16)     //设置加密数组
	copy(key, []byte(c.keyStr)) //合并数组补位

	res := ""
	result, err := c.AesEncrypt([]byte(str), key)
	if err == nil {
		res = base64.StdEncoding.EncodeToString(result)
	}
	return res
}

/**
* 数据加密,返回[]byte
 */
func (c *AES) AesEncrypt(origData, key []byte) ([]byte, error) {
	if len(origData) < 1 || len(key) < 1 {
		return nil, nil
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	if c.ty == "pkcs5" {
		origData = c.PKCS5Padding(origData, blockSize)
	} else {
		origData = c.ZeroPadding(origData, block.BlockSize())
	}
	IV := key[:blockSize]
	if len(c.IV) > 0 {
		IV = []byte(c.IV)
	}
	blockMode := cipher.NewCBCEncrypter(block, IV)
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

/**
* 数据解密
* @param string	str 需要解密的字符串
* @return strintg 返回解密后的字符串
 */
func (c *AES) AesDecryptString(str string) string {
	if len(str) < 1 {
		return ""
	}
	key := make([]byte, 16)     //设置加密数组
	copy(key, []byte(c.keyStr)) //合并数组补位

	res := ""
	str = strings.Replace(str, "%2B", "+", -1)
	result, _ := base64.StdEncoding.DecodeString(str) //将字符串变成[]byte

	if len(result) > 0 {
		origData, err := c.AesDecrypt(result, key)
		if err == nil {
			res = string(origData)
		}
	}
	return res
}

/**
*数据解密，返回[]byte
 */
func (c *AES) AesDecrypt(crypted, key []byte) ([]byte, error) {
	if len(crypted) < 1 || len(key) < 1 {
		return nil, nil
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := recover(); err != nil {
			base64.StdEncoding.EncodeToString(crypted)
		}
	}()

	blockSize := block.BlockSize()

	IV := key[:blockSize]
	if len(c.IV) > 0 {
		IV = []byte(c.IV)
	}

	blockMode := cipher.NewCBCDecrypter(block, IV)
	origData := make([]byte, len(crypted))
	// origData := crypted

	blockMode.CryptBlocks(origData, crypted)
	if c.ty == "pkcs5" {
		origData = c.PKCS5UnPadding(origData)
	} else {
		origData = c.ZeroUnPadding(origData)
	}
	return origData, nil
}

/**
* Zero补位算法
 */
func (c *AES) ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func (c *AES) ZeroUnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

/**
* PKCS5补位算法
 */
func (c *AES) PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func (c *AES) PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
