package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	engine := gee.New()
	engine.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	})
	engine.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})
	engine.Run(":8080")
}
