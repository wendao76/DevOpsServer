package http

import (
	"log"
	"github.com/go-session/session"
	"net/http"
)

//帐号密码校验器
func PasswordAuthorizationHandler(username, password string) (userID string, err error) {
	log.Println("PasswordAuthorizationHandler")
	if username == "test2" && password == "test2" {
		userID = "testUserID"
	}
	return
}

//TODO 未实现
func AuthorizeScopeHandler(w http.ResponseWriter, r *http.Request) (scope string, err error) {
	log.Println("AuthorizeScopeHandler")
	return
}

func UserAuthorizeHandler(w http.ResponseWriter, r *http.Request) (userID string, err error) {
	log.Println("UserAuthorizeHandler")
	store, err := session.Start(nil, w, r)
	if err != nil {
		return
	}
	uid, ok := store.Get("LoggedInUserID")
	if !ok {
		if r.Form == nil {
			r.ParseForm()
		}

		store.Set("ReturnUri", r.Form)
		store.Save()

		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusFound)
		return
	}

	userID = uid.(string)
	store.Delete("LoggedInUserID")
	store.Save()
	return
}
