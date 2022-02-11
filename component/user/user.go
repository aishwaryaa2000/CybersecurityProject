package user

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"security/component/encrypt"
	"security/component/file"
)

var UserList []*User

func CreateUser(fname, lname, username, password, role string) {
	var bell, biba int
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	publicKey := privateKey.PublicKey
	salt := fname + lname
	saltHash := encrypt.CreateHash(salt)
	newPassword := password + saltHash
	hashPassword := encrypt.CreateHash(newPassword)
	if role == "hr" {
		bell = 3
		biba = 2
	} else if role == "sales" {
		bell = 2
		biba = 2
	} else if role == "dev" {
		bell = 2
		biba = 3
	} else if role == "intern" {
		bell = 1
		biba = 1
	} else if role == "head" {
		bell = 3
		biba = 3
	} else {
		bell = 1
		biba = 1
	}
	newUser := NewUser(fname+lname, username, hashPassword, role, bell, biba, privateKey, publicKey, nil)
	UserList = append(UserList, newUser)
	fmt.Println("User successfull added")

}

func NewUser(name, userName, password, role string, bellLevel, bibaLevel int, privateKey *rsa.PrivateKey, publicKey rsa.PublicKey, files []*file.File) *User {
	var user = &User{
		name:       name,
		username:   userName,
		password:   password,
		role:       role,
		bellLevel:  bellLevel,
		bibaLevel:  bibaLevel,
		privateKey: privateKey,
		publicKey:  publicKey,
		files:      files,
	}
	return user
}

type User struct {
	name       string
	username   string
	password   string
	role       string
	bellLevel  int
	bibaLevel  int
	privateKey *rsa.PrivateKey
	publicKey  rsa.PublicKey
	files      []*file.File
}
