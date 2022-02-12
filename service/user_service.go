package service

import (
	"bufio"
	"fmt"
	"os"
	"security/component/user"
	"security/component/filelog"
	"strings"
)

func UserService() {
	userMenu()
}

func userMenu() {
	user.ReadData()
	for {
		fmt.Println("-------MENU-------", "\nEnter 1 to Register", "\nEnter 2 to Login", "\nEnter 3 to list all Users", "\nEnter 4 to Logout")
		reader := bufio.NewReader(os.Stdin)
		ch, _ := reader.ReadString('\n')
		ch = strings.TrimSpace(ch)
		switch ch {
		case "1":
			register()
		case "2":
			login()
		case "3":
			user.ListAllUser()
		case "4":
			user.WriteData()
			fmt.Println("Exiting!!!")
			return
		default:
			fmt.Println("Wrong Choice")
		}
	}
}

func register() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Name:")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Println("Enter Designation Name:")
	des, _ := reader.ReadString('\n')
	des = strings.TrimSpace(des)

	fmt.Println("Enter User Name:")
	uname, _ := reader.ReadString('\n')
	uname = strings.TrimSpace(uname)

	fmt.Println("Enter Password:")
	pass, _ := reader.ReadString('\n')
	pass = strings.TrimSpace(pass)

	//Create New User
	user.CreateUser(name, uname, pass, des)
	filelog.WriteUserLog(uname," signed in")


}

func login() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter User Id:")
	uid, _ := reader.ReadString('\n')
	uid = strings.TrimSpace(uid)

	fmt.Println("Enter Password:")
	pass, _ := reader.ReadString('\n')
	pass = strings.TrimSpace(pass)

	//Authenticate User then call
	check, bell, biba := user.CheckUser(uid, pass)
	if check {
		filelog.WriteUserLog(uid," logged in")
		AppService(bell, biba, uid)
		filelog.WriteUserLog(uid," logged out")
		return
	}
	fmt.Println("Wrong Credentials")
}
