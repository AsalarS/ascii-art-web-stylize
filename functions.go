package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

// Function to create ErrorPageData for the error
func createErrorPageData(statusCode int, err error) ErrorPageData {
	return ErrorPageData{
		Title:      fmt.Sprintf("%d", statusCode),
		StatusCode: statusCode,
		Message:    err.Error(),
	}
}

// Function to handle errors and return appropriate HTTP response
func handleError(w http.ResponseWriter, statusCode int, err error) {
	errorData := createErrorPageData(statusCode, err)
	w.WriteHeader(statusCode)
	errorPageTemplate.Execute(w, errorData)
}

// returnAscii generates an ASCII art representation of the input string using the specified art type.
// It returns the ASCII art string and any error encountered during the process.
func returnAscii(input string, artType string) (string, error) {
	if input == "" {
		return "", errors.New("no input")
	}
	Lines := splitString(input)
	fileLines := readFile(artType)
	if fileLines == nil {
		return "", errors.New("banner file not found")
	} else if len(fileLines) != 855 {
		return "", errors.New("file corrupted")
	}
	var asciiString string
	for index, l := range Lines {
		if index <= 1 && l == "" {
			continue
		} else if l == "" {
			asciiString += "\n"
			continue
		}
		if index > 0 {
			asciiString += "\n"
		}
		for i := 0; i <= 7; i++ {
			for _, c := range l {
				if int(c) < 32 || int(c) > 126 {
					return "", errors.New("invalid character")
				}
				defer func() {
					if r := recover(); r != nil {
						return
					}
				}()
				asciiString += fileLines[getFirstLine(int(c))+i]
			}
			if i < 7 {
				asciiString += "\n"
			}
		}
	}
	return asciiString, nil
}

func splitString(input string) []string {
	input = strings.ReplaceAll(input, "\\n", "\n")
	// Split the input string based on newline character
	lines := strings.Split(input, "\n")

	result := append([]string{}, lines...)

	return result
}

// Function that reads the Ascii-Fonts Files and returns an array of strings for each line.
func readFile(filename string) []string {
	if !fileExists("Ascii-Fonts/" + filename + ".txt") {
		return nil
	}
	file, err := os.Open("Ascii-Fonts/" + filename + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

// Function to check if a file exists or not
func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

// Function to get the first line of the character in the Ascii-Fonts files
func getFirstLine(ascii int) int {
	line := ((ascii - 32) * 9) + 1
	return line
}
