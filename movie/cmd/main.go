package main

import (
	"log"
	"net/http"

	"github.com/high-la/movieapp/movie/internal/controller/movie"
	metadataGateway "github.com/high-la/movieapp/movie/internal/gateway/metadata/http"
	ratingGateway "github.com/high-la/movieapp/movie/internal/gateway/rating/http"
	httphandler "github.com/high-la/movieapp/movie/internal/handler/http"
)

func main() {

	log.Println("Starting the movie service")

	metadataGateway := metadataGateway.New("localhost:8081")
	ratingGateway := ratingGateway.New("localhost:8082")

	ctrl := movie.New(ratingGateway, metadataGateway)
	h := httphandler.New(ctrl)

	http.Handle("/movie", http.HandlerFunc(h.GetMovieDetails))
	if err := http.ListenAndServe(":8083", nil); err != nil {
		panic(err)
	}
}
