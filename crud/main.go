package main

import (
	"bufio"
	"crud/utils"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var text string

	for text != "exit" {
		fmt.Println("To add new user type: add\nTo remove user type: remove\nTo exit from app type: exit\nTo display all users type: print")
		fmt.Print("> ")
		text = utils.ReadFromStdIn(reader)

		switch text {
		case "add":
			utils.AddUser(reader)
		case "print":
			utils.PrintUser()
		case "remove":
			utils.RemoveUser(reader)
		}
	}
}
