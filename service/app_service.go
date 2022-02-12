package service

import (
	"bufio"
	"fmt"
	"os"
	"security/component/file"
	"security/component/filelog"
	"strings"
)

func AppService(bell, biba int, uid string) {
	appMenu(bell, biba, uid)
}

func appMenu(bell, biba int, uid string) {
	file.Read()
	for {
		fmt.Println("-------MENU-------", "\nEnter 1 to Read File", "\nEnter 2 to Write File", "\nEnter 3 to Read Mails", "\nEnter 4 to Send Mails", "\nEnter 5 to Exit")
		reader := bufio.NewReader(os.Stdin)
		ch, _ := reader.ReadString('\n')
		ch = strings.TrimSpace(ch)
		switch ch {
		case "1":
			readFile(bell, biba, uid)
		case "2":
			writeFile(bell, biba, uid)
		case "3":
			readMails(uid)
		case "4":
			sendMails(uid)
		case "5":
			fmt.Println("Logging Out!!!")
			return
		default:
			fmt.Println("Wrong Choice")
		}
	}
}

func readFile(bell, biba int, uid string) {
	//show all file slice
	fmt.Println("The ReadAble files are")
	cnt := file.ReadAble(bell, biba)
	if cnt <= 0 {
		fmt.Println("No files to read")
		return
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter any listed files name with extension:")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	//if file is in readable list->show its content
	err := file.ReadFile(name)
	if err != nil {
		fmt.Println("Error:", err)
	}
	filelog.WriteFileLog(uid, name, "reads")
}

func writeFile(bell, biba int, uid string) {
	fmt.Println("The Write able files are")
	//get count of all right able files
	cnt := file.WriteAble(bell, biba)
	if cnt <= 0 {
		fmt.Println("No files to write")
		return
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter any listed files name with extension:")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	content, _ := reader.ReadString('\n')
	content = strings.TrimSpace(content)

	//if file is in writeable list->append to the file
	err := file.WriteFile(name, content)
	if err != nil {
		fmt.Println("Error:", err)
	}
	filelog.WriteFileLog(uid, name, "writes")
}
