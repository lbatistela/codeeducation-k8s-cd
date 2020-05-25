package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", SumServer)
	http.ListenAndServe(":8000", nil)
}

func SumServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Code.education Rocks!\n5 + 5 = %v", soma(5, 5))
}