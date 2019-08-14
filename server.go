package main

import (
	"net/http"
	"fmt"
	"os"
	"log"
	"io/ioutil"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b,_ := ioutil.ReadAll(r.Body)
		fmt.Fprintf(w,"Hello from %s\nreceived body: %s\n", os.Getenv("WHO"), b)
	})))
}