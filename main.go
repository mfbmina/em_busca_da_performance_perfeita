package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

// func main() {
// 	log.Println("booting on localhost:8080")

// 	r := chi.NewRouter()
// 	r.Mount("/debug", middleware.Profiler())

// 	log.Fatal(http.ListenAndServe(":8080", r))
// }

func main() {
	log.Println("booting on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
