package main

import (
	"net/http"
	"strings"
)

var contentTypeSetter = func(h http.Handler) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		if strings.HasSuffix(req.URL.Path, ".wasm") {
			resp.Header().Set("content-type", "application/wasm")
		}
		h.ServeHTTP(resp, req)
	})
}

func main() {
	fs := http.FileServer(http.Dir("./html"))
	http.ListenAndServe(":8080", contentTypeSetter(fs))
}
