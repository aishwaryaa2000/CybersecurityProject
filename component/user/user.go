package user

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
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
	publicKey  *rsa.PublicKey
	mailFiles  []*string
}

var UserList []*User

func CreateUser(name, username, password, role string) {
	var bell, biba int
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	publicKey := &privateKey.PublicKey
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
	fmt.Println(name, username, role, hashPassword)
}

func NewUser(name, userName, password, role string, bellLevel, bibaLevel int, privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey, files []*string) *User {
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
		fmt.Println(" Designation:", val.role)
		fmt.Println("Public Key : ", val.publicKey)
		// fmt.Println("Private Key : ",val.privateKey)
		fmt.Println()

	}
}

func ListUserName() {
	fmt.Println("All username")
	for _, val := range UserList {
		fmt.Println(val.username)

	}
}

func GetMailFiles(username string) []*string {
	for _, val := range UserList {
		if username == val.username {
			return val.mailFiles
		}
	}
	return nil
}

func GetPublicPrivateKey(username string) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	for _, val := range UserList {
		if username == val.username {
			return val.privateKey, val.publicKey, nil
		}
	}
	err := errors.New("Username not found")
	return nil, nil, err
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
	// reader := bufio.NewReader(f)
	// for {
	// 	data, err := reader.ReadString('\n')
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	storeData(strings.TrimSpace(data))
	// }
	data, err := ioutil.ReadFile("users.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	storeData(string(data))
}

func storeData(data string) {
	//dataSlice := make([]string, 0)
	data1 := strings.Split(data, "fmt")
	for _, records := range data1 {
		records = strings.TrimSpace(records)
		if records == "" {
			break
		}
		dataSlice := strings.Split(records, "line")
		var tempUser User
		tempUser.name = dataSlice[0]
		tempUser.username = dataSlice[1]
		tempUser.password = dataSlice[2]
		tempUser.role = dataSlice[3]
		tempUser.bellLevel, _ = strconv.Atoi(dataSlice[4])
		tempUser.bibaLevel, _ = strconv.Atoi(dataSlice[5])
		tempUser.privateKey, _ = ParseRsaPrivateKeyFromPemStr((dataSlice[6]))
		tempUser.publicKey, _ = ParseRsaPublicKeyFromPemStr((dataSlice[7]))
		if dataSlice[8] == "" {
			tempUser.mailFiles = nil

		} else {
			//var fileName string
			for _, fileName := range dataSlice[8] {
				files := strconv.QuoteRune(fileName)
				tempUser.mailFiles = append(tempUser.mailFiles, &files)
			}
		}

		UserList = append(UserList, &tempUser)
		//fmt.Println(UserList)
	}
}
func WriteData() {
	f, err := os.OpenFile("users.txt", os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var data string
	for _, val := range UserList {
		privKey := ExportRsaPrivateKeyAsPemStr(val.privateKey)
		pubKey, _ := ExportRsaPublicKeyAsPemStr(val.publicKey)
		data = val.name + "line" + val.username + "line" + val.password + "line" + val.role + "line" + fmt.Sprintf("%d", (int(val.bellLevel))) + "line" + fmt.Sprintf("%d", (int(val.bibaLevel))) + "line" + string(privKey) + "line" + string(pubKey) + "line" //Write Public & Private RSA Key
		for _, val1 := range val.mailFiles {
			data += *val1 + "line"
		}
		f.WriteString(data + "fmt\n")
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

func ExportRsaPrivateKeyAsPemStr(privkey *rsa.PrivateKey) string {
	privkey_bytes := x509.MarshalPKCS1PrivateKey(privkey)
	privkey_pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privkey_bytes,
		},
	)
	return string(privkey_pem)
}

func ExportRsaPublicKeyAsPemStr(pubkey *rsa.PublicKey) (string, error) {
	pubkey_bytes, err := x509.MarshalPKIXPublicKey(pubkey)
	if err != nil {
		return "", err
	}
	pubkey_pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubkey_bytes,
		},
	)

	return string(pubkey_pem), nil
}

func ParseRsaPrivateKeyFromPemStr(privPEM string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privPEM))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return priv, nil
}

func ParseRsaPublicKeyFromPemStr(pubPEM string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pubPEM))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	switch pub := pub.(type) {
	case *rsa.PublicKey:
		return pub, nil
	default:
		break // fall through
	}
	return nil, errors.New("Key type is not RSA")
}

func (u *User) AppendFiles(fileName string) {
	u.mailFiles = append(u.mailFiles, &fileName)
	for _, singleFile := range u.mailFiles {
		fmt.Println(*singleFile)

	}

}

//Display All Mails received to user.
// func ReadMail() {

// }
