package services

import (
	"bufio"
	"cypher/internal/database"
	"cypher/internal/database/models"
	"cypher/internal/errors"
	"encoding/json"
	"github.com/fatih/color"
	"io/ioutil"
	"os"
	"regexp"
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

func (c CredentialService) FuzzySearch(query string, attribute string) (CredentialList, error) {
	credentialList, err := c.ListCredential()
	if err != nil {
		panic(err)
	}

	var newCredentialList CredentialList
	for _, v := range credentialList {
		var value string
		if attribute == "domain" {
			value = v.Domain
		} else {
			value = v.Note
		}

		matched, err := regexp.MatchString(query, value)
		if err == nil && matched {
			newCredentialList = append(newCredentialList, v)
		}
	}

	return newCredentialList, nil
}

func (c CredentialService) SetCredential(credential models.Credential) (models.Credential, error) {
	credentialList, err := c.ListCredential()
	if err != nil {
		panic(err)
	}
	flag := false
	for i, v := range credentialList {
		if v.Domain == credential.Domain {
			_credential := credentialList[i]
			if credential.Username != "" {
				_credential.Username = credential.Username
			}
			if credential.Password != "" {
				_credential.Password = credential.Password
			}
			if credential.Note != "" {
				_credential.Note = credential.Note
			}

			flag = true
			break
		}
	}
	if !flag {
		if credential.Domain == "" || credential.Password == "" {
			color.Red("请至少输入 domain 和 password 以设置！")
			return models.Credential{}, errors.SetKeyError
		}
		credentialList = append(credentialList, &credential)
	}
	ciphers, err := json.Marshal(credentialList)
	if err != nil {
		panic(err)
	}
	err = writeFile(ciphers)
	if err != nil {
		return models.Credential{}, errors.SetKeyError
	}
	return credential, nil
}

func (c CredentialService) DeleteCypherByDomain(domain string) {
	credentialList, err := c.ListCredential()
	if err != nil {
		panic(err)
	}
	flag := false
	for i, v := range credentialList {
		if v.Domain == domain {
			credentialList = append(credentialList[:i], credentialList[i+1:]...)
			flag = true
			break
		}
	}
	if !flag {
		color.Red("未找到该域名的凭证！")
		return
	}
	ciphers, err := json.Marshal(credentialList)
	if err != nil {
		panic(err)
	}
	err = writeFile(ciphers)
	if err != nil {
		color.Red("写入文件失败！")
	}
}

func (c CredentialService) GetCredentialByDomain(domain string) models.Credential {
	credentialList, err := c.ListCredential()
	if err != nil {
		panic(err)
	}
	for _, v := range credentialList {
		if v.Domain == domain {
			return *v
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
	if _, err = file.Write(content); err != nil {
		return errors.WriteFileError
	}
	defer file.Close()
	return nil
}
