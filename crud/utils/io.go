package utils

import (
	"bufio"
	"crud/modules/user"
	"fmt"
	"strconv"
	"strings"
)

func ReadFromStdIn(reader *bufio.Reader) string {
	readText, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return strings.TrimRight(readText, "\r\n")
}

func AddUser(reader *bufio.Reader) {
	fmt.Print("Ender name: ")
	readText := ReadFromStdIn(reader)
	u := user.UserEntity{}
	u.New()
	u.SetName(readText)
	fmt.Print("Ender age: ")
	readText = ReadFromStdIn(reader)
	age, err := strconv.Atoi(readText)
	if err != nil {
		fmt.Println("Enter number")
		return
	}
	u.SetAge(age)
	u.Append()
	fmt.Printf("New user added:\n%v\n", u)
}

func PrintUser() {
	for i := 0; i < len(user.Users); i++ {
		fmt.Println(user.Users[i])
	}
}

func RemoveUser(reader *bufio.Reader) {
	fmt.Print("To remove user enter ID: ")
	readText := ReadFromStdIn(reader)
	id, err := strconv.Atoi(readText)
	if err != nil {
		fmt.Println("Enter number")
		return
	}
	for i := 0; i < len(user.Users); i++ {
		u := user.Users[i]
		if u.ID == id {
			user.Users = removeIndex(user.Users, i)
			fmt.Printf("Removed user ID: %d\n", u.ID)
		}
	}
}

func removeIndex[T comparable](s []T, index int) []T {
	return append(s[:index], s[index+1:]...)
}
