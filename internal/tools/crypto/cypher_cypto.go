package crypto

type CypherCrypt interface {
	Encrypt(plainText string) (string, error)
	Decrypt(cipherText string) (string, error)
}
