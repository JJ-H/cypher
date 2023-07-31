package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"cypher/config"
	"encoding/base64"
	"fmt"
	"github.com/fatih/color"
	"io"
)

var (
	CypherAES CypherCrypt
)

func init() {
	cfg := config.Config
	if cfg == nil {
		err := config.InitConfig()
		if err != nil {
			CypherAES = &CypherAes{}
			return
		}
	}
	aes, _ := config.Config.GetSection("crypto")
	key, err := aes.GetKey("key")
	if err != nil {
		color.Red("获取加密密钥失败！")
		return
	}
	CypherAES = &CypherAes{key: []byte(key.String())}
}

type CypherAes struct {
	key []byte
}

func (c *CypherAes) Encrypt(plainText string) (string, error) {
	if !c.isEnabled() {
		return plainText, nil
	}
	data := []byte(plainText)
	block, err := aes.NewCipher(c.key)
	if err != nil {
		return "", err
	}

	// 随机 IV 向量
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// CBC 加密
	cipherText := make([]byte, aes.BlockSize+len(data))
	copy(cipherText[:aes.BlockSize], iv)
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(cipherText[aes.BlockSize:], data)

	encoded := base64.URLEncoding.EncodeToString(cipherText)
	return encoded, nil
}

func (c *CypherAes) Decrypt(cipherText string) (string, error) {
	// 未配置密钥，直接返回原文
	if !c.isEnabled() {
		return cipherText, nil
	}
	data, err := base64.URLEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(c.key)
	if err != nil {
		return "", err
	}

	if len(data) < aes.BlockSize {
		return "", fmt.Errorf("密文太短")
	}

	// 提取 IV 向量
	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]

	// CBC 解密
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(data, data)
	return string(data), nil
}

func (c *CypherAes) isEnabled() bool {
	return c.key != nil
}
