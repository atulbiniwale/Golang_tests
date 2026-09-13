package note

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	NoteTitle   string
	NoteContent string
	NoteCreated time.Time
}

func (n Note) Display() {
	fmt.Printf("************************************\n Your note titled: %v has content: %v and is created at %v\n", n.NoteTitle, n.NoteContent, n.NoteCreated)
}

func (n Note) Save() error {
	filename := strings.ReplaceAll(n.NoteTitle, " ", "_")
	filename = strings.ToLower(filename) + ".json"

	json, err := json.Marshal(n)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, json, 0644)

}

func NewNote(title, content string) Note {

	return Note{
		NoteTitle:   title,
		NoteContent: content,
		NoteCreated: time.Now(),
	}

}
