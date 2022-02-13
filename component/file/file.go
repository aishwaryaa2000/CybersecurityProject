package file

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"security/component/encrypt"
	"strconv"
	"strings"
)

type File struct {
	name     string
	bell_lvl int
	biba_lvl int
}

var allFiles []*File

func Read() {
	allFiles = nil
	f, err := os.Open("files.txt")
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
	var temp File
	temp.name = dataSlice[0]
	temp.bell_lvl, _ = strconv.Atoi(dataSlice[1])
	temp.biba_lvl, _ = strconv.Atoi(dataSlice[2])
	allFiles = append(allFiles, &temp)
}

var read []*File
var write []*File

func ReadAble(bell, biba int) int {
	read = nil
	for _, val := range allFiles {
		if int(val.bell_lvl) <= bell && val.biba_lvl >= biba {
			fmt.Println(val.name)
			read = append(read, val)
		}
	}
	return len(read)
}

func WriteAble(bell, biba int) int {
	write = nil
	for _, val := range allFiles {
		if val.bell_lvl >= bell && val.biba_lvl <= biba {
			fmt.Println(val.name)
			write = append(write, val)
		}
	}
	return len(write)
}

func ReadFile(name string) error {
	for _, val := range read {
		if val.name == name {
			readFile(name)
			return nil
		}
	}
	return errors.New("file not in the given list")
}

func readFile(name string) {
	data, _ := ioutil.ReadFile("files/" + name)
	if len(data) > 0 {
		for _, val := range strings.Split(string(data), "line") {
			if len(val) > 0 {
				plainText := string(encrypt.DecryptFile([]byte(strings.TrimSpace(val))))
				fmt.Println(plainText)
			}
		}
		return
	}
	fmt.Println("Please Note: File is empty")
}

func WriteFile(name string, phrase string) error {
	for _, val := range write {
		if val.name == name {
			f, err := os.OpenFile("files/"+name, os.O_APPEND, 0644)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			cipherText := encrypt.EncryptFile([]byte(phrase))
			_, err = f.WriteString("line" + string(cipherText) + "\n")
			if err != nil {
				log.Panic(err)
			}
			return nil
		}
	}
	return errors.New("File not in the given list")
}
