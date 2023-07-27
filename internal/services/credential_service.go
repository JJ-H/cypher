package services

import (
	"bufio"
	"cipher_manager/internal/database"
	"cipher_manager/internal/database/models"
	"encoding/json"
	"io/ioutil"
	"os"
)

type CredentialList = models.CredentialList

type CredentialService struct{}

var (
	CredentialSrv = &CredentialService{}
)

func (c CredentialService) ListCredential() (CredentialList, error) {
	reader := bufio.NewReader(database.DB)
	ciphers, err := ioutil.ReadAll(reader)
	var credentialList CredentialList
	if err != nil {
		return credentialList, err
	}
	err = json.Unmarshal(ciphers, &credentialList)
	if err != nil {
		return credentialList, nil
	}
	return credentialList, nil
}

func (c CredentialService) SetCredential(credential models.Credential) models.Credential {
	credentialList, err := c.ListCredential()
	if err != nil {
		panic(err)
	}
	flag := false
	for i, v := range credentialList {
		if v.Domain == credential.Domain {
			credentialList[i] = credential
			flag = true
			break
		}
	}
	if !flag {
		credentialList = append(credentialList, credential)
	}
	ciphers, err := json.Marshal(credentialList)
	if err != nil {
		panic(err)
	}
	err = writeFile(ciphers)
	if err != nil {
		return models.Credential{}
	}
	return credential
}

func (c CredentialService) DeleteCypherByDomain(domain string) {
	credentialList, err := c.ListCredential()
	if err != nil {
		panic(err)
	}
	for i, v := range credentialList {
		if v.Domain == domain {
			credentialList = append(credentialList[:i], credentialList[i+1:]...)
			break
		}
	}
	ciphers, err := json.Marshal(credentialList)
	if err != nil {
		panic(err)
	}
	err = writeFile(ciphers)
}

func (c CredentialService) GetCredentialByDomain(domain string) models.Credential {
	credentialList, err := c.ListCredential()
	if err != nil {
		panic(err)
	}
	for _, v := range credentialList {
		if v.Domain == domain {
			return v
		}
	}
	return models.Credential{}
}

func writeFile(content []byte) error {
	dir, err := os.UserHomeDir()
	var file *os.File
	if err != nil {
		panic(err)
	}
	file, err = os.OpenFile(dir+"/cypher.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	_, err = file.Write(content)
	defer file.Close()
	return err
}
