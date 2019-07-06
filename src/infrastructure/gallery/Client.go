package gallery

import (
	"encoding/json"
	"github.com/ugniusin/instago/src/domain/gallery/dto"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Client struct {
	ClientId string
	ClientSecret string
	RedirectUri string
}

func NewGalleryClient() *Client {
	return &Client{}
}

func (client *Client) GetAccessToken (clientId string, clientSecret string, redirectUri string, code string) string {
	accessTokenUrl := "https://api.instagram.com/oauth/access_token"

	resp, err := http.PostForm(accessTokenUrl,
		url.Values{
			"client_id": {clientId},
			"client_secret": {clientSecret},
			"grant_type": {"authorization_code"},
			"redirect_uri": {redirectUri},
			"code": {code},
		})

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var responseBody map[string]interface{}

	if err := json.Unmarshal(body, &responseBody); err != nil {
		panic(err.Error())
	}

	return responseBody["access_token"].(string)
}

func (client *Client) GetUserDetails (accessToken string) dto.User {
	resp, err := http.Get("https://api.instagram.com/v1/users/self/?access_token=" + accessToken)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var responseBody map[string]interface{}

	if err := json.Unmarshal(body, &responseBody); err != nil {
		panic(err.Error())
	}

	data := responseBody["data"].(map[string]interface{})

	return dto.User{
		Username: data["username"].(string),
		FullName: data["full_name"].(string),
		ProfilePicture: data["profile_picture"].(string),
	}
}
