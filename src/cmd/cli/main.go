package main

import (
	"net/http"

	src "bz.openstep/gcpserverless"
)

func main() {
	http.HandleFunc("/hello", src.HelloWorld)
	http.ListenAndServe(":8080", nil)
}
