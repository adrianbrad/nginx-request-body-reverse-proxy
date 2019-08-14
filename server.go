package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", http.HandlerFunc(helloHandler)))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "Hello from %s\nreceived body: %s\n", os.Getenv("WHO"), b)
}
