package main

import (
	"github.com/ugniusin/instago/framework/config"
	galleryControllers "github.com/ugniusin/instago/src/application/gallery/controllers"
	"net/http"
)

func main() {
	config := config.GetConfigs("./config.json")

	galleryController := galleryControllers.NewGalleryController(
		config.Instagram["ClientId"],
		config.Instagram["ClientSecret"],
		config.Instagram["RedirectUri"],
	)

	http.HandleFunc("/auth", galleryController.Authorise)
	http.HandleFunc("/redirect", galleryController.Redirect)

	http.ListenAndServe(":8090", nil)
}
