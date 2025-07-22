package model

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"

	"github.com/Alan-Gomes1/go-api/src/configuration/logger"
)

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8
	GetJSONValue() (string, error)
	SetID(id string)
	EncryptPassword()
}

type userDomain struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Age      int8   `json:"age"`
}

func NewUserDomain(email, password, name string, age int8) UserDomainInterface {
	return &userDomain{
		Email: email,
		Password: password,
		Name: name,
		Age:  age,
	}
}

func (u *userDomain) GetEmail() string {
	return u.Email
}

func (u *userDomain) GetPassword() string {
	return u.Password
}

func (u *userDomain) GetName() string {
	return u.Name
}

func (u *userDomain) GetAge() int8 {
	return u.Age
}

func (u *userDomain) GetJSONValue() (string, error) {
	b, err := json.Marshal(u)
	if err != nil {
		logger.Error("Error trying to marshal user", err)
		return "", err
	}
	return string(b), nil
}

func (u *userDomain) SetID(id string) {
	u.ID = id
}

func (u *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(u.Password))
	u.Password = hex.EncodeToString(hash.Sum(nil))
}
