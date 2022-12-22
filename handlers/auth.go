package handlers

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gyu-young-park/go-oauth2.0-practice/config"
	"golang.org/x/oauth2"
)

func oauthGoogleLoginHandler(w http.ResponseWriter, r *http.Request) {
	oauthState := createStateOauthCookie(w)
	//To get refresh token, Set access_type = offline
	redirectURL := config.ConfigMap.GoogleOauthConfig.AuthCodeURL(oauthState, oauth2.SetAuthURLParam("access_type", "offline"))
	fmt.Println(config.ConfigMap.GoogleOauthConfig.ClientID)
	http.Redirect(w, r, redirectURL, http.StatusTemporaryRedirect)
}

func oauthGoogleCallbackHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("oauthGoogleCallbackHandler start")
	oauthState, _ := r.Cookie("oauthstate")
	fmt.Println("oauthstate: ", oauthState)

	if r.FormValue("state") != oauthState.Value {
		log.Println("Invalid oauth google state")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	code := r.FormValue("code")
	token, err := getTokenFromGoogle(code)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	data, err := getUserDataFromGoogle(token)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	fmt.Fprintf(w, "UserInfo: %s\n %+v", data, prettyPrint(token))
}

func getTokenFromGoogle(code string) (*oauth2.Token, error) {
	token, err := config.ConfigMap.GoogleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange error: %s", err)
	}
	return token, err
}

func getUserDataFromGoogle(token *oauth2.Token) ([]byte, error) {
	res, err := http.Get(config.OauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer res.Body.Close()
	contents, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed read res: %s", err)
	}
	return contents, nil
}
