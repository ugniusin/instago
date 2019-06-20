package controllers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type GalleryController struct {
}

func NewGalleryController() *GalleryController {
	return &GalleryController{}
}

func (c *GalleryController) Authorise(w http.ResponseWriter, r *http.Request) {

	var clientId = ""
	var redirectUri = "http://localhost:8090/redirect"

	var url = "https://api.instagram.com/oauth/authorize/?client_id=" +
		clientId +
		"&redirect_uri=" +
		redirectUri +
		"&response_type=code"

	http.Redirect(w, r, url, 301)
}

func (c *GalleryController) Redirect(w http.ResponseWriter, req *http.Request)  {
	fmt.Fprintf(w, "hello\n")

	keys, ok := req.URL.Query()["code"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	// Query()["key"] will return an array of items,
	// we only want the single item.
	code := string(keys[0])

	getAccessToken(code)


	fmt.Fprintf(w, "byebye\n")

}

func getAccessToken(code string) {
	url := "https://api.instagram.com/oauth/access_token"
	fmt.Println("URL:>", url)



	var jsonStr = []byte(`{"client_id":"","client_secret":"","grant_type":"authorization_code","redirect_uri":"http://localhost:8090/redirect","code": ` + code + `}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")


	//body, _ := ioutil.ReadAll(req.Body)
	//fmt.Println("response Body:", string(body))



	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

}