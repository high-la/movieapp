package main

import (
	"log"
	"net/http"

	"github.com/high-la/movieapp/metadata/internal/controller/metadata"
	httphandler "github.com/high-la/movieapp/metadata/internal/handler/http"
	"github.com/high-la/movieapp/metadata/internal/repository/memory"
)

func main() {

	log.Println("starting the movie metadata service")

	repo := memory.New()
	ctrl := metadata.New(repo)
	h := httphandler.New(ctrl)

	http.Handle("/metadata", http.HandlerFunc(h.GetMetadata))

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
	}
}
