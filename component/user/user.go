package user

import (
	"bufio"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"io"
	"log"
	"os"
	"security/component/encrypt"
	"strconv"
	"strings"
)

type User struct {
	name       string
	username   string
	password   string
	role       string
	bellLevel  int
	bibaLevel  int
	privateKey *rsa.PrivateKey
	publicKey  rsa.PublicKey
	mailFiles  []string
}

var UserList []*User

func CreateUser(name, username, password, role string) {
	var bell, biba int
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	publicKey := privateKey.PublicKey
	saltHash := encrypt.CreateHash(name)
	newPassword := password + saltHash
	hashPassword := encrypt.CreateHash(newPassword)
	switch role {
	case "hr":
		bell = 3
		biba = 2
	case "sales":
		bell = 2
		biba = 2
	case "dev":
		bell = 2
		biba = 3
	case "intern":
		bell = 1
		biba = 1
	case "head":
		bell = 3
		biba = 3
	default:
		bell = 1
		biba = 1
	}
	newUser := NewUser(name, username, hashPassword, role, bell, biba, privateKey, publicKey, nil)
	UserList = append(UserList, newUser)
	fmt.Println("User successfull added")
}

func NewUser(name, userName, password, role string, bellLevel, bibaLevel int, privateKey *rsa.PrivateKey, publicKey rsa.PublicKey, files []string) *User {
	var user = &User{
		name:       name,
		username:   userName,
		password:   password,
		role:       role,
		bellLevel:  bellLevel,
		bibaLevel:  bibaLevel,
		privateKey: privateKey,
		publicKey:  publicKey,
		mailFiles:  files,
	}
	return user
}

func ListAllUser() {
	for _, val := range UserList {
		fmt.Print("First Name:", val.name)
		fmt.Print(" Userid:", val.username)
		fmt.Print(" Password:", val.password)
		fmt.Print(" Designation:", val.role)
		fmt.Println()
	}
}

func ReadData() {
	var _, err = os.Stat("users.txt")
	if err != nil {
		f, _ := os.Create("users.txt")
		f.Close()
	}
	f, err := os.Open("users.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	for {
		data, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		storeData(strings.TrimSpace(data))
	}
}

func storeData(data string) {
	dataSlice := strings.Split(data, ",")
	var tempUser User
	tempUser.name = dataSlice[0]
	tempUser.username = dataSlice[1]
	tempUser.password = dataSlice[2]
	tempUser.role = dataSlice[3]
	tempUser.bellLevel, _ = strconv.Atoi(dataSlice[4])
	tempUser.bibaLevel, _ = strconv.Atoi(dataSlice[5])
	UserList = append(UserList, &tempUser)
}

func WriteData() {
	f, err := os.OpenFile("users.txt", os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var data string
	for _, val := range UserList {
		data = val.name + "," + "," + val.username + "," + val.password + "," + val.role + "," + fmt.Sprintf("%d", (int(val.bellLevel))) + "," + fmt.Sprintf("%d", (int(val.bibaLevel))) + "," //Write Public & Private RSA Key
		for _, val1 := range val.mailFiles {
			data += val1 + ","
		}
		f.WriteString(data + "\n")
	}
}

func CheckUser(userid, pass string) (bool, int, int) {
	for _, val := range UserList {
		checkpass := encrypt.CreateHash(pass + encrypt.CreateHash(val.name))

		if checkpass == val.password && val.username == userid {
			return true, val.bellLevel, val.bibaLevel
		}
	}
	return false, 0, 0
}

//Display All Mails received to user.
// func ReadMail() {

// }
