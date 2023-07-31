package crypto

import (
	"fmt"
	"testing"
)

var (
	key       = []byte("1234567890123456")
	plainText = "hello world"
)

func TestCypherAes_Encrypt(t *testing.T) {
	aes := &CypherAes{key: key}
	cipherText, err := aes.Encrypt(plainText)
	if err != nil {
		t.Error(err)
	}
	t.Log(cipherText)
}

func TestCypherAes_Decrypt(t *testing.T) {
	aes := &CypherAes{key: key}
	cipherText, err := aes.Encrypt(plainText)
	if err != nil {
		t.Error(err)
	}
	t.Log(cipherText)
	plainText, err = aes.Decrypt(cipherText)
	if err != nil {
		t.Error(err)
	}
	t.Log(plainText)
}

func TestCrypto(t *testing.T) {
	aes := &CypherAes{key: key}
	encryptedText, err := aes.Encrypt(plainText)
	if err != nil {
		t.Error(err)
	}
	t.Log(encryptedText)

	decryptedText, err := aes.Decrypt(encryptedText)

	if decryptedText != plainText {
		t.Log("加解密正常")
	} else {
		fmt.Println("原密码:" + plainText + "\n 加密后:" + encryptedText + "\n 解密后:" + decryptedText)
		t.Error("加解密失败")
	}

}
