package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpResponse(r, true)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(dump))
	fmt.Fprintf(w, "<html><body>hello</body></html>")
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("start http listening :18443")
	err := http.ListenAndServe(":18443", "server.crt", "server.key", nil)
	log.Println(err)
}
