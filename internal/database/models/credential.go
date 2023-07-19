package models

type Credential struct {
	ID       int64  `json:"id"`
	Domain   string `json:"domain"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type CredentialList []Credential

func (c Credential) ToString() string {
	return c.Domain + "\t" + c.Username + "\t" + c.Password
}
