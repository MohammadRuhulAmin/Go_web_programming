package main // project should start from here!

import (
	"fmt"
	"net/http" // these are packages ! creating web server , handling web request and so on!
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html") // text/plain
	fmt.Fprint(w, "<h1>Welcome to___ my super Awesome Tutorial</h1>")
}

func main() {

	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)

}
