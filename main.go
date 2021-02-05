package main

import (
	"fmt"
	"log"
	"net/http"
)

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	log.Println("running main handler")
	_, _ = w.Write([]byte(`{"message": "hello delve"}`))
}

func main() {
	config, err := loadConfig()
	if err != nil {
		panic("failed to lod config:" + err.Error())
	}

	fmt.Println("The config is ", config)
	s := &server{}
	http.Handle("/", s)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
