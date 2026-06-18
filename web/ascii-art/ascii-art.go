package asciiart

import (
	"strings"
	"os"
	"fmt"
)
func AsciiArt(input string, banners string)string{
	result := ""
	

	content, err := os.ReadFile("banner/" + banners + ".txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		// tpl.ExecuteTemplate(w,"index.html","Error reading file")
	}
	banner := strings.Split(string(content), "\n")
	lines:= strings.Split(input, "\\n")

	if len(lines) > 1 && lines[0] == "" && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	// var test string

	// loop through each line of the input
	for _, line := range lines {

		// if the line is empty, print a blank line and move on
		if line == "" {
			fmt.Println("Input something")
			continue
		}

		// each ASCII art character is 8 rows tall
		// so we print row by row (0 to 7)
		for row := 0; row < 8; row++ {
			

			// for each character, find its art in the banner
			// (int(char) - 32) * 9 + 1 calculates the starting line
			for _, char := range line {
				startIndex := (int(char)-32) *9 + 1
				result += banner[startIndex+row]
			}

			// print this row of all characters combined
			// fmt.Println(result)
			result = result + "\n" 
			fmt.Print(result)
			// tpl.Execute(w,result)

			
		}
	}
	return result
}
