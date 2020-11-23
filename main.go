package main

import (
	"fmt"
	"net/http"
)

func HomeRouterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "Hello, World2!")
}
func main(){
	http.HandleFunc("/", HomeRouterHandler)
	fmt.Println("Listen 8090")
	http.ListenAndServe(":8090", nil)
	for{
	}
	fmt.Println("End")
}