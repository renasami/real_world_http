package main

import (
	"net/http"
	"net/http/httputil"
)

func handler(w *http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpResponse(r, true)
}
