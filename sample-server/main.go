package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", helloHandler)

	fmt.Println("Starting webserver at %s", ":8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Panic(err)
	}

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Haaris\n"))
}
