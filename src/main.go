package main

import (
	"net/http"
	"io"
)


func hello(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "hello ")
}

func main(){
	http.HandleFunc("/",hello)
	http.ListenAndServe(":3000", nil)
}