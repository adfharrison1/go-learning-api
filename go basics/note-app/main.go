package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/adfharrison1/go-learning-api/note-app/note"
	"github.com/adfharrison1/go-learning-api/note-app/todo"
)

type saver interface {
	Save() error
}

type handleable interface {
	saver
	Display()
}

func main() {

	title, content := getNoteData()
	todotext := getTodoData()

	newNote, error := note.New(title, content)

	handleError(error)

	newTodo, error := todo.New(todotext)

	handleError(error)

	error = saveData(newNote)

	handleError(error)

	error = handleData(newTodo)

	handleError(error)

}

func handleError(error error) {
	if error != nil {
		fmt.Println(error.Error())
		return
	}
}

func handleData(data handleable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	error := data.Save()

	if error != nil {
		fmt.Println("Error saving file")
		return error
	}

	fmt.Println("Note saved successfully!")

	return nil
}

func getTodoData() string {

	content := getUserInput("Todo text: ")

	return content
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")

	content := getUserInput("Note content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Println(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
