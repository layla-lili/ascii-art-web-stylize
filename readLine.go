package asciiart

import (
	"bufio"
	"log"
	"os"
	"strings"
)

/*
this function used to read a single line of ASCII art from the file.
It takes the file path and line number as arguments and returns the line of text as a string.
*/

func ReadLine(text string, file *os.File) string {
	result := ""
	for i := 0; i < 8; i++ {
		var sb strings.Builder
		// Loop through the letters in the text string
		for _, letter := range text {
			// Set the file pointer to the beginning of the file
			_, err := file.Seek(0, 0)
			if err != nil {
				log.Fatal(err)
			}
			// Create a new scanner for the file
			scanner := bufio.NewScanner(file)
			currentLine := 1
			// Loop through the file
			for scanner.Scan() {
				//check if the current line is the correct line to read
				if currentLine == 2+int(letter-' ')*9+i {
					//Append the text to the string builder
					sb.WriteString(scanner.Text())
				}
				currentLine++
			}
		}
		//print the final string
		//fmt.Println(sb.String())
		result = result + sb.String() + "\n"
	}
	return result
}