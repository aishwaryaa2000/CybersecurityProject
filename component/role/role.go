package role

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Role struct {
	designation string
	bell_lvl    int
	biba_lvl    int
}

var allRoles []*Role

func ChooseRole() string {
	var i int
	fmt.Println("---Company Roles---")
	CompanyRole()
	for i, singleRole := range allRoles {
		fmt.Println(i+1, ".", singleRole.designation)
	}
	fmt.Println("Choose your Role in the company:")
	fmt.Scanln(&i)
	for j, singleRole := range allRoles {
		if j == (i - 1) {
			return singleRole.designation
		}
	}
	return "Invalid Role"
}

func AssignLevels(des string) (int, int) {
	for _, singleRole := range allRoles {
		if singleRole.designation == des {
			return singleRole.bell_lvl, singleRole.biba_lvl
		}
	}
	return 0, 0

}

func CompanyRole() {
	allRoles = nil
	f, err := os.Open("role.txt")
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
		store(strings.TrimSpace(data))
	}
}

func store(data string) {
	dataSlice := strings.Split(data, ",")
	var temp Role
	temp.designation = dataSlice[0]
	temp.bell_lvl, _ = strconv.Atoi(dataSlice[1])
	temp.biba_lvl, _ = strconv.Atoi(dataSlice[2])
	allRoles = append(allRoles, &temp)
}
