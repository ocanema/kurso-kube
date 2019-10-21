package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

// HelloServer function: writes response for client
func HelloServer(rw http.ResponseWriter, req *http.Request) {

	hostname, error := os.Hostname()
	if error != nil {
		panic(error)
	}

	fmt.Fprint(rw, "Hello from ", hostname)

}
