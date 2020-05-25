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
	fmt.Fprintf(w, "Code.education Rocks!\n40 + 2 = %v", soma(40, 2))
}