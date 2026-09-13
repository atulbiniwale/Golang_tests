package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"atulb.com/noteapp/note"
)

// func to get user input.
func getNoteTitle() string {
	fmt.Printf("note title:\n")
	reader := bufio.NewReader(os.Stdin)
	titleValue, _ := reader.ReadString('\n')
	titleValue = strings.TrimSuffix(titleValue, "\n")
	titleValue = strings.TrimSuffix(titleValue, "\r")
	return titleValue
}

func getNoteContent() string {
	fmt.Printf("note content:\n")
	reader := bufio.NewReader(os.Stdin)
	contentValue, _ := reader.ReadString('\n')
	contentValue = strings.TrimSuffix(contentValue, "\n")
	contentValue = strings.TrimSuffix(contentValue, "\r")
	return contentValue
}

func main() {
	title := getNoteTitle()
	content := getNoteContent()
	userNote := note.NewNote(title, content)
	userNote.Display()
	userNote.Save()
}
