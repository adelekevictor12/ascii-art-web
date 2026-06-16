package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

type test struct{
	result string

}
var tpl  *template.Template
func main(){
	tpl,_ = template.ParseFiles("index.html")
	http.HandleFunc("/index",indexhandler)
	http.HandleFunc("/ascii_art",arthandler)
	http.ListenAndServe(":8080",nil)

}

func indexhandler(w http.ResponseWriter, r *http.Request){
	tpl.Execute(w,"")
	


}

func arthandler (w http.ResponseWriter, r *http.Request){
	var s test
	input := r.FormValue("input_val")

	content, err := os.ReadFile(r.FormValue("selector") + ".txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		tpl.ExecuteTemplate(w,"index.html","Error reading file")
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
			// result := ""

			// for each character, find its art in the banner
			// (int(char) - 32) * 9 + 1 calculates the starting line
			for _, char := range line {
				startIndex := (int(char)-32) *9 + 1
				s.result += banner[startIndex+row]
			}

			// print this row of all characters combined
			// fmt.Println(result)
			s.result = s.result + "\n" 
			fmt.Print(s.result)
			// tpl.Execute(w,result)

			
		}
	}
	fmt.Print(s.result)
	
	tpl.ExecuteTemplate(w,"index.html",s.result)
}
