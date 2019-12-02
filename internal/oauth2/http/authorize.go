package http

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-session/session"
	"go_web/internal/oauth2/model"
	"go_web/internal/oauth2/service"
	"gopkg.in/oauth2.v3/server"
	"log"
	"net/http"
	"net/url"
	"time"
)

type OAuthAction struct {
	Srv *server.Server
}

//获取token
func (s *OAuthAction) Token(ctx *gin.Context) {
	w := ctx.Writer
	r := ctx.Request
	err := s.Srv.HandleTokenRequest(w, r)
	if err != nil {
		errors(ctx, 500, "token获取失败")
	}
}

//获取授权
func (s *OAuthAction) Authorize(ctx *gin.Context) {
	log.Println("Authorize")
	w := ctx.Writer
	r := ctx.Request
	store, err := session.Start(nil, w, r)
	if err != nil {
		errors(ctx, http.StatusInternalServerError, err.Error())
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
func (s *OAuthAction) RegisterClient(ctx *gin.Context) {
}


//跳转到登录页面
func (s *OAuthAction) LoginPage(ctx *gin.Context) {
	w := ctx.Writer
	r := ctx.Request
	_, err := session.Start(nil, w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	html(ctx, 200, "login.html")
}

//登录操作
// TODO 帐号密码校验
func (s *OAuthAction) Login(ctx *gin.Context) {
	log.Println("POST.Login")
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
func (s *OAuthAction) AuthPage(ctx *gin.Context) {
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

//自测
func (s *OAuthAction) Test(ctx *gin.Context) {
	oauthSrv := service.OauthService

	standardClaims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
	}
	user := &model.OAuthUser {
	    Username: "wendao",
	}
	user.Uid = 1000
	log.Println("tokenStr")
	token, err := oauthSrv.GenJWTAccessToken(user, standardClaims)
	if err != nil {
		fmt.Printf("GenJWTAccessToken error:" , err.Error())
	}
	log.Println(token)


	claims, err2 := oauthSrv.ParseJWTAccessToken(token)
	if err2 != nil {
		log.Println("parse error:" + err2.Error())
		//log.Fatal(err.Error())
	}
	log.Println("username:", claims.GetUsername())
	result(ctx,nil, http.StatusOK)
}

