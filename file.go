package main

import (
	"bufio"
	"fmt"
	"os"
)

// WriteObjectsToFile writes strings derived from a list of objects to a file
func WriteObjectsToFile[T any](filename string, objects []T, toString func(T) string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, obj := range objects {
		line := toString(obj) + "\n"
		_, err := writer.WriteString(line)
		if err != nil {
			return fmt.Errorf("error writing to file: %w", err)
		}
	}
	if err := writer.Flush(); err != nil {
		return fmt.Errorf("error flushing writer: %w", err)
	}
	return nil
}
