package main

import (
	"github.com/ugniusin/instago/application/gallery/controllers"
	"net/http"
)

func main() {
	galleryController := controllers.NewGalleryController()

	http.HandleFunc("/auth", galleryController.Authorise)
	http.HandleFunc("/redirect", galleryController.Redirect)

	http.ListenAndServe(":8090", nil)
}
