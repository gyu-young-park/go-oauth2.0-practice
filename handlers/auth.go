package handlers

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gyu-young-park/go-oauth2.0-practice/config"
)

func oauthGoogleLoginHandler(w http.ResponseWriter, r *http.Request) {
	oauthState := createStateOauthCookie(w)
	redirectURL := config.ConfigMap.GoogleOauthConfig.AuthCodeURL(oauthState)
	fmt.Println(config.ConfigMap.GoogleOauthConfig.ClientID)
	http.Redirect(w, r, redirectURL, http.StatusTemporaryRedirect)
}

func oauthGoogleCallbackHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("oauthGoogleCallbackHandler start")
	oauthState, _ := r.Cookie("oauthstate")

	if r.FormValue("state") != oauthState.Value {
		log.Println("Invalid oauth google state")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	data, err := getUserDataFromGoogle(r.FormValue("code"))
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	fmt.Fprintf(w, "UserInfo: %s\n", data)
}

func getUserDataFromGoogle(code string) ([]byte, error) {
	token, err := config.ConfigMap.GoogleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange error: %s", err)
	}
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
