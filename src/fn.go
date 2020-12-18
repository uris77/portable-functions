package src

import (
	"fmt"
	"net/http"
	"os"
)

// HelloWorld writes "Hello, World!" to the HTTP response.
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	msg := os.Getenv("MSG")
	fmt.Fprint(w, fmt.Sprintf("Hello, World from %s!\n", msg))
}
