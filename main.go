package main

import (
	"fmt"
	"net/http"
)

func main() {
	var port = ":8080"

	fmt.Println("Server listening on port:", port)
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(port, nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
