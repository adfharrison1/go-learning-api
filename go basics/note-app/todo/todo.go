package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("invalid input")
	}
	return Todo{
		content,
	}, nil
}

func (todo Todo) Display() {
	fmt.Println(todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"
	fileContent, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, fileContent, 0644)
}
