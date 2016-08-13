package main

import (
	"net/url"
	"fmt"
	"log"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	Env_load()
	var param  HttpParam
	param = HttpParam {
		ClientId: os.Getenv("CLIENT_ID"),
		ResponseType: "code",
		Scope: "sleep",
		RedirectURI: os.Getenv("REDIRECT_URI"),
		Expires: "31536000",

	}
	var authUrl AuthUrl
	authUrl = AuthUrl{
		Host: "www.fitbit.com",
		Path: "/oauth2/authorize",

	}
	url := &url.URL{}
	url.Scheme = "https"
	url.Host = authUrl.Host
	url.Path = authUrl.Path
	query := url.Query()
	query.Set("response_type", "code")
	query.Set("client_id", param.ClientId)
	query.Set("redirect_uri", param.RedirectURI)
	query.Set("scope",param.Scope)
	query.Set("expires_in", param.Expires)
	url.RawQuery = query.Encode()
	fmt.Println(url)

}

type HttpParam struct {
	ClientId string
	ResponseType string
	Scope	string
	RedirectURI string
	Expires string
}

type AuthUrl struct {
	Host string
	Path string
}

func Env_load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}