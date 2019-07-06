package controllers

import (
	"encoding/json"
	"github.com/ugniusin/instago/src/infrastructure/gallery"
	"log"
	"net/http"
)

type GalleryController struct {
	ClientId string
	ClientSecret string
	RedirectUri string
}

func NewGalleryController(
	ClientId string,
	ClientSecret string,
	RedirectUri string,
) *GalleryController {

	return &GalleryController{
		ClientId,
		ClientSecret,
		RedirectUri,
	}
}

func (controller *GalleryController) Authorise (w http.ResponseWriter, r *http.Request) {

	var authorizeUrl = "https://api.instagram.com/oauth/authorize/?client_id=" +
		controller.ClientId +
		"&redirect_uri=" +
		controller.RedirectUri +
		"&response_type=code"

	http.Redirect(w, r, authorizeUrl, 301)
}

func (controller *GalleryController) Redirect (w http.ResponseWriter, req *http.Request)  {
	keys, ok := req.URL.Query()["code"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	// Query()["key"] will return an array of items,
	// we only want the single item.
	code := string(keys[0])

	GalleryClient := gallery.NewGalleryClient()
	accessToken := GalleryClient.GetAccessToken(controller.ClientId, controller.ClientSecret, controller.RedirectUri, code)

	userDto := GalleryClient.GetUserDetails(accessToken)

	js, err := json.Marshal(userDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
