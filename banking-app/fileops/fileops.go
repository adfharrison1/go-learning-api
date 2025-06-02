package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(value float64, filename string) {
	valueString := fmt.Sprint(value)
	os.WriteFile(filename, []byte(valueString), 0644)
}

func ReadFloatFromFile(filename string) (float64, error) {
	value, readErr := os.ReadFile(filename)
	if readErr != nil {
		return 0, errors.New("failed to find file")
	}
	floatValue, floatErr := strconv.ParseFloat(string(value), 64)
	if floatErr != nil {
		return 0, errors.New("failed to parse stored file")
	}
	return floatValue, nil
}
