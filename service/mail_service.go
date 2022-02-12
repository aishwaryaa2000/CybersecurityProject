package service

import (
	"bufio"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"security/component/encrypt"
	"security/component/user"
	"strings"
	// "security/component/encrypt"
)

var secret = "key"

func readMails(username string) {
	fmt.Println("-----INBOX-----")
	allMails := user.GetMailFiles(username)
	if allMails == nil {
		fmt.Println("No Mails present in Inbox")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter any listed files name with extension which you wish to read:")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	name = "mailFiles/" + name

	file, err := os.Open(name)
	if err != nil {
		log.Fatalf("Failed to open")

	}
	defer file.Close()
	hmacWithMail, _ := ioutil.ReadFile(name) //entire data in byte format

	privateKey, _, _ := user.GetPublicPrivateKey(username)

	ok, mailEncrypted := checkHmacSame(string(hmacWithMail))

	if ok {
		fmt.Println("Hmac recieved and sent is same")
		plainText := encrypt.DecryptMail(privateKey, mailEncrypted)
		fmt.Println(plainText)
	} else {
		fmt.Println("Hmac recieved and sent is different")
	}
}

func checkHmacSame(hmacWithMail string) (bool, string) {
	hmacRecieved := hmacWithMail[:64]
	msgRecieved := hmacWithMail[64:] //msg is recieved in encrypted format by rsa

	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(msgRecieved))
	hmacCode := hex.EncodeToString(h.Sum(nil))

	if hmacRecieved == hmacCode {
		return true, msgRecieved

	}
	return false, ""

}

func sendMails(username string) {

	reader := bufio.NewReader(os.Stdin)
	user.ListUserName()
begin:
	fmt.Println("To which user do you wish to send the mail?")
	uid, _ := reader.ReadString('\n')
	uid = strings.TrimSpace(uid)

	_, publicKey, err := user.GetPublicPrivateKey(uid)
	if err != nil {
		fmt.Println("Incorrect username")
		goto begin
	}

	fmt.Println("Enter a one word subject of your mail")
	subject, _ := reader.ReadString('\n')
	subject = strings.TrimSpace(subject)

	fmt.Println("Enter the data of your mail")
	data, _ := reader.ReadString('\n')
	data = strings.TrimSpace(data)
	data = "Sender of this mail is : " + username + "\n" + data

	h := hmac.New(sha256.New, []byte(secret))
	mailEncrypted := encrypt.EncryptMail(publicKey, data)
	h.Write([]byte(mailEncrypted))
	hmacCode := hex.EncodeToString(h.Sum(nil)) //hmac code by using key and encrypted mail
	hmacWithMail := hmacCode + mailEncrypted   //appending hmacCode with mail

	//Now,this hmacWithMail is recieved by the reciever
	sendToReciever(subject, hmacWithMail, uid)

}

func sendToReciever(subject, hmacWithMail, uid string) {

	fileName := subject + ".txt"
	var err error
	f, _ := os.Create("mailFiles/" + fileName)
	f.Close()
	f, err = os.OpenFile("mailFiles/"+fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, errr := f.WriteString(hmacWithMail) //Write the content into the mail inbox of the reciever

	if errr != nil {
		log.Fatal(errr)
	}

	//Get the MailFiles slice of the reciever so that we can append the filename
	user.AppendFiles(uid, fileName)

	fmt.Println("Mail has been sent to ", uid)

}
