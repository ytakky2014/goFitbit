package main

import (
	"fmt"
	"net/http"
	_"net/url"
	_"encoding/base64"
	_"strings"
	_"io/ioutil"
	"net/url"
	"encoding/base64"
	"strings"
	"io/ioutil"
	"encoding/json"
	"time"
	"log"
	"os"
	"github.com/joho/godotenv"
)

type String  string

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// envからIDなどを取得
	Env_load()
	r.ParseForm()
    code,_  := r.Form["code"];
	codeStr := code[0]

	client := &http.Client{}
	var param  HttpParam
	param = HttpParam {
		ClientId: os.Getenv("CLIENT_ID"),
		ResponseType: "code",
		Scope: "sleep",
		RedirectURI: os.Getenv("REDIRECT_URI"),
		Expires: "31536000",
	}

	values := url.Values{}
	values.Set("code", codeStr)
	values.Add("grant_type", "authorization_code")
	values.Add("clientId", param.ClientId)
	values.Add("redirect_uri", param.RedirectURI)

	key := os.Getenv("SECRET_KEY")
	keyEnc := base64.StdEncoding.EncodeToString([]byte(key))



	req, _:= http.NewRequest("POST", "https://api.fitbit.com/oauth2/token", strings.NewReader(values.Encode()))
	req.Header.Set("Authorization", " Basic " + keyEnc)
	req.Header.Set("content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	bodyStr:=string(body)
	jsonBytes := ([]byte)(bodyStr)
	//fmt.Fprint(w,bodyStr)
	decodeBody := new(FitbitJson)
	err = json.Unmarshal(jsonBytes, decodeBody)
//	fmt.Println(decodeBody)
//	fmt.Println(decodeBody.AccessToken)
	now := time.Now().Add(-24 * time.Hour).Format("2006-01-02")

	fitbitApiURL := "https://api.fitbit.com/1/user/" + string(decodeBody.UserId) + "/sleep/date/" + now + ".json";
	client = &http.Client{}
	req, _= http.NewRequest("GET", fitbitApiURL, nil)

	req.Header.Set("Authorization", " Bearer " + decodeBody.AccessToken)
	res, _ = client.Do(req)


	body, _ = ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	fmt.Fprintln(w, string(body))


}

func main() {
	http.HandleFunc("/", ServeHTTP)
	http.ListenAndServe("localhost:8080", nil)
}

type HttpParam struct {
	ClientId string
	ResponseType string
	Scope	string
	RedirectURI string
	Expires string
}

type FitbitJson struct {
	AccessToken string `json:"access_token"`
	ExpiresIn int	`json:"expires_in"`
	RefreshToken String `json:"refresh_token"`
	Scope String `json:"scope"`
	TokenType String `json:"token_type"`
	UserId String `json:"user_id"`
}


func Env_load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}