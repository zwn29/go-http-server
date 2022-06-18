package main

import (
	"encoding/json"
	"log"
	"net/http"
)


type Book struct{
	Title string `json:"title"`
	Author string `json:"author"`
	Pages int `json:"pages"`
}

func Hello(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"message\":\"Hello World\"}"))
}

func GetBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	book := Book{
		Title:"The Gunslinger",
		Author: "Stephen King",
		Pages: 340,
	}
	json.NewEncoder(w).Encode(book)
}

func Html(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<h1 style='color:blue'>this is go server</h1>"))
}


func SimpleText(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is a simple text request"))
}

func root(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<h1 style='color:green;text-align:center;font-family:Arial'>welcome to go</h1>"))
}

func main() {
	http.HandleFunc("/",root)
	http.HandleFunc("/book",GetBook)
	http.HandleFunc("/hello",Html)
	http.HandleFunc("/text",SimpleText)
	log.Fatal(http.ListenAndServe(":3000",nil))
}
