package services

import (
	"cipher_manager/internal/database"
	"cipher_manager/internal/database/models"
	"fmt"
	"testing"
)

func init() {
	database.DB = database.InitDB()
}

func TestSetCredential(t *testing.T) {
	c := models.Credential{
		ID:       1,
		Domain:   "baidu.com",
		Username: "admin",
		Password: "666666",
	}
	CredentialSrv.SetCredential(c)
}

func TestListCredential(t *testing.T) {
	credentialList, err := CredentialSrv.ListCredential()
	fmt.Println(credentialList)
	if err != nil {
		t.Error("list credential error: ", err)
	}
	t.Log("credential list: ", credentialList)
}

func TestGetCredential(t *testing.T) {
	domain := "gitee.com"
	cypher := CredentialSrv.GetCredentialByDomain(domain)
	t.Log("cypher: ", cypher.ToString())
}

func TestDeleteCypher(t *testing.T) {
	domain := "gitee.com"
	CredentialSrv.DeleteCypherByDomain(domain)
	if CredentialSrv.GetCredentialByDomain(domain).Domain != "" {
		t.Error("delete cypher error")
	}
}
