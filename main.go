package main

import "net/http"

func main() {
	serveMux := http.ServeMux{}
	server := http.Server{
		Addr:    ":8080",
		Handler: &serveMux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
