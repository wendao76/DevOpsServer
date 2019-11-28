package http

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-session/session"
	"gopkg.in/oauth2.v3/server"
	"net/http"
	"net/url"
)

type OAuthService struct {
	Srv *server.Server
}

//获取token
func (s *OAuthService) Token(ctx *gin.Context) {
	w := ctx.Writer
	r := ctx.Request
	err := s.Srv.HandleTokenRequest(w, r)
	if err != nil {
		errors(ctx, 500, "token获取失败")
	}
}

//获取授权
func (s *OAuthService) Authorize(ctx *gin.Context) {
	w := ctx.Writer
	r := ctx.Request
	store, err := session.Start(nil, w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var form url.Values
	if v, ok := store.Get("ReturnUri"); ok {
		form = v.(url.Values)
	}
	r.Form = form

	store.Delete("ReturnUri")
	store.Save()

	err = s.Srv.HandleAuthorizeRequest(w, r)
	if err != nil {
		errors(ctx, http.StatusBadRequest, err.Error())
	}
}

//TODO 客户端注册
func (s *OAuthService) RegisterClient(ctx *gin.Context) {

}

func (s *OAuthService) Test(ctx *gin.Context) {
	w := ctx.Writer
	r := ctx.Request
	token, err := s.Srv.ValidationBearerToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data := map[string]interface{}{
		"expires_in": int64(token.GetAccessCreateAt().Add(token.GetAccessExpiresIn()).Sub(time.Now()).Seconds()),
		"client_id":  token.GetClientID(),
		"user_id":    token.GetUserID(),
	}
	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	e.Encode(data)
}

//跳转到登录页面
func (s *OAuthService) LoginPage(ctx *gin.Context) {
	w := ctx.Writer
	r := ctx.Request
	_, err := session.Start(nil, w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	html(ctx, 200, "login.html")
}

//TODO 登录操作
func (s *OAuthService) Login(ctx *gin.Context) {
	type LoginModel struct {
		Username string `form:"username" json:"username" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
	}
	w := ctx.Writer
	r := ctx.Request
	store, err := session.Start(nil, w, r)
	if err != nil {
		log.Fatal(err.Error())
	}

	var loginData LoginModel

	err = ctx.Bind(&loginData)
	if err != nil {
		log.Fatal(err.Error())
	}
	store.Set("LoggedInUserID", loginData.Username)
	store.Save()

	ctx.Redirect(http.StatusFound, "/auth")
}

//授权页面
func (s *OAuthService) AuthPage(ctx *gin.Context) {
	w := ctx.Writer
	r := ctx.Request
	store, err := session.Start(nil, w, r)
	fmt.Println(store)
	if err != nil {
		errors(ctx, 500, err.Error())
	}

	if _, ok := store.Get("LoggedInUserID"); !ok {
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusFound)
		return
	}
	html(ctx, http.StatusOK, "auth.html")
}
