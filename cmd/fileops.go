package main

import (
	"os"
	"strconv"
)

// loadCount reads the count value from a file
func loadCount(filename string) (int, error) {
	data, err := os.ReadFile(filename) // Using os.ReadFile
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(string(data))
}

// saveCount writes the count value to a file
func saveCount(filename string, count int) error {
	return os.WriteFile(filename, []byte(strconv.Itoa(count)), 0644) // Using os.WriteFile
}
