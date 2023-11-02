package main

import (
	"fmt"
	"net/http"
)

func main() {

	// http.Handle("/", http.FileServer(http.Dir("./img")))
	http.Handle("/recources/", http.StripPrefix("/recources/", http.FileServer(http.Dir("img"))))

	http.HandleFunc("/",home)
	http.HandleFunc("/contact",contact)
	http.HandleFunc("/about",about)
	
	http.ListenAndServe(":8080",nil)
}

func home (w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, `WELCOME to our HOME page`)
}
func contact (w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, `WELCOME to CONTACT US page`)
}
func about (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "WELCOME to ABOUT US page <img src=\"makka.jpg\" />")
}