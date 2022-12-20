package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"time"
)

func createStateOauthCookie(w http.ResponseWriter) string {
	var expire = time.Now().Add(24 * time.Hour)

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expire}
	http.SetCookie(w, &cookie)

	return state
}
