package service

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func AppService() {
	appMenu()
}

func appMenu() {
	for {
		fmt.Println("-------MENU-------", "\nEnter 1 to Create Text File", "\nEnter 2 to Read File", "\nEnter 3 to Write File", "\nEnter 4 to Read Mails", "\nEnter 5 to Send Mails", "\nEnter 6 to Exit")
		reader := bufio.NewReader(os.Stdin)
		ch, _ := reader.ReadString('\n')
		ch = strings.TrimSpace(ch)
		switch ch {
		case "1":
			createFile()
		case "2":
			readFile()
		case "3":
			writeFile()
		case "4":
			readMails()
		case "5":
			sendMails()
		case "6":
			fmt.Println("Logging Out!!!")
			return
		default:
			fmt.Println("Wrong Choice")
		}
	}
}

func createFile() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter File Name:")
	fName, _ := reader.ReadString('\n')
	fName = strings.TrimSpace(fName)

	var bell, biba int
	var err error
	for {
		fmt.Println("Enter Bell La Level:")
		temp, _ := reader.ReadString('\n')
		temp = strings.TrimSpace(temp)
		bell, err = strconv.Atoi(temp)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		if bell <= 0 || bell > 5 {
			fmt.Println("Error:Please enter Bell La Level from 1 to 5")
			continue
		}
		break
	}

	for {
		fmt.Println("Enter Biba Level:")
		temp, _ := reader.ReadString('\n')
		temp = strings.TrimSpace(temp)
		biba, err = strconv.Atoi(temp)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		if biba <= 0 || biba > 5 {
			fmt.Println("Error:Please enter Biba Level from 1 to 5")
			continue
		}
		break
	}

	//create new file
	fmt.Println(fName, bell, biba)
}

func readFile() {
	//show all file slice
	fmt.Println("The ReadAble files are")
	//show count of all readable file
	// if cnt <= 0 {
	// 	fmt.Println("No files to read")
	// 	continue
	// }
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter any listed files name with extension:")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	//if file is in readable list->show its content
	// err := files.ReadFile(name)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }
}

func writeFile() {
	fmt.Println("The Write able files are")
	//get count of all right able files
	// cnt := files.WriteAble(bell, biba)
	// if cnt <= 0 {
	// 	fmt.Println("No files to write")
	// 	continue
	// }
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter any listed files name with extension:")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	content, _ := reader.ReadString('\n')
	content = strings.TrimSpace(content)

	//if file is in writeable list->append to the file
	// err := files.WriteFile(name, phrase)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }
}
