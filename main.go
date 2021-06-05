package main // project should start from here!

import (
	"fmt"
	"net/http" // these are packages ! creating web server , handling web request and so on!
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html") // text/plain
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome to___ my super Awesome Tutorial</h1>")
	} else if r.URL.Path == "/Contact" {
		fmt.Fprint(w, "<p>Email Us : </p> <a href=\"facebook.com\" >Contact</a>  ") // linking ..a page
	} else {
		w.WriteHeader(http.StatusNotFound) 
		fmt.Fprint(w, "<h3>Bad Request !</h3>")
	}
	fmt.Println(r.URL.Path)
}

func main() {
	mux := &http.ServeMux{}
	mux.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", mux)

}
