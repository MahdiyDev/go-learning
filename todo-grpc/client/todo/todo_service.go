package todo

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CreateTodo() *TodoMessage {
	reader := bufio.NewReader(os.Stdin)
	var s TodoMessage
	fmt.Print("Enter title: ")
	s.Title = ReadFromStdIn(reader)
	fmt.Print("Enter context: ")
	s.Context = ReadFromStdIn(reader)
	return &s
}

func ReadFromStdIn(reader *bufio.Reader) string {
	readText, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return strings.TrimRight(readText, "\r\n")
}
