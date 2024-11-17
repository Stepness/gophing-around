package main

import (
	"api/handlers"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Best API started!")

	l := log.New(os.Stdout, "api", log.LstdFlags)

	mux := http.NewServeMux()
	ph := handlers.NewProducts(l)

	mux.Handle("GET /", ph)

	s := http.Server{
		Addr:     ":8080",
		Handler:  mux,
		ErrorLog: l,
	}

	_ = s.ListenAndServe()
}
