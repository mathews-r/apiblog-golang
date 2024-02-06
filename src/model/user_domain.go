package model

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/mathews-r/golang/src/configs/rest_err"
)

func NewUserDomain(
	email, password, name string,
	age int,
) UserDomainInterface {
	return &UserDomain{
		email, password, name, age,
	}
}

type UserDomain struct {
	Email    string
	Password string
	Name     string
	Age      int
}

func (ud *UserDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}

type UserDomainInterface interface {
	CreateUser() *rest_err.RestErr
	// UpdateUser(string) *rest_err.RestErr
	// FindUser(string) (*UserDomain, *rest_err.RestErr)
	// DeleteUser(string) *rest_err.RestErr
}
