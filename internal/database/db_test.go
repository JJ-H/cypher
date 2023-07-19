package database

import (
	"io/ioutil"
	"testing"
)

func TestDb(t *testing.T) {
	_db := InitDB()
	content, err := ioutil.ReadAll(_db)
	if err != nil {
		t.Error("read db error: ", err)
	}
	t.Log("content", string(content))
}
