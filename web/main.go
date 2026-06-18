package main

import (
	
	"html/template"
	"net/http"

	"ascii-art/ascii-art"
)


var tpl  *template.Template
func main(){
	tpl,_ = template.ParseFiles("template/index.html")
	http.HandleFunc("/index",indexhandler)
	http.HandleFunc("/ascii-art",arthandler)
	http.ListenAndServe(":8080",nil)

}

func indexhandler(w http.ResponseWriter, r *http.Request){
	tpl.Execute(w,"")
	


}

func arthandler (w http.ResponseWriter, r *http.Request){
	
	input := r.FormValue("input_val")
	banners := r.FormValue("selector")

	
	tpl.ExecuteTemplate(w,"index.html",asciiart.AsciiArt(input,banners))
}
