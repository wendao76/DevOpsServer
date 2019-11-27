package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/go-session/session"
	"go_web/internal/common/config"
	oredis "gopkg.in/go-oauth2/redis.v3"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
	"log"
	"net/http"
	"net/url"
)

func New() {
	//gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())

	srv := initOAuthService()
	initRouter(srv, engine)
	err := engine.Run(":9096")
	if err != nil {
		log.Fatal("用户服务启动失败!")
	}
}

//初始化oauth2处理流程
func initOAuthService() *server.Server {
	manager := manage.NewDefaultManager()
	conf := config.GetInstance()
	redisConf := conf.Redis
	// token memory store
	manager.MapTokenStorage(oredis.NewRedisStore(
		&redis.Options{
			Addr: redisConf.Addr,
			DB: redisConf.Db,
		}, "oauth2-cli-"))

	// client memory store
	clientStore := store.NewClientStore()
	clientStore.Set("000000", &models.Client{
		ID:     "000000",
		Secret: "999999",
		Domain: "http://localhost:9096",
	})
	manager.MapClientStorage(clientStore)

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)
	srv.SetPasswordAuthorizationHandler(func(username, password string) (userID string, err error) {
		fmt.Println("密码校验")
		if username == "test" && password == "test" {
			userID = "test"
		}
		return
	})
	srv.SetAuthorizeScopeHandler(func(w http.ResponseWriter, r *http.Request) (scope string, err error) {
		fmt.Println("authorizeScopeHandler")
		return
	})
	srv.SetUserAuthorizationHandler(userAuthorizeHandler)
	return srv
}

func initRouter(srv *server.Server, engine *gin.Engine) {
	//身份教养（password|client_credentials）
	engine.GET("/token", func(ctx *gin.Context) {
		w := ctx.Writer
		r := ctx.Request
		err := srv.HandleTokenRequest(w, r)
		if err != nil {
		    errors(ctx, 500, "token获取失败")
		}
	})

	engine.POST("/authorize", func(ctx *gin.Context) {
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
		fmt.Println("session处理")
		err = srv.HandleAuthorizeRequest(w, r)
		if err != nil {
			errors(ctx, 500, "身份校验失败:" + err.Error())
		}
		ctx.JSON(200, gin.H{})
	})
}

func userAuthorizeHandler(w http.ResponseWriter, r *http.Request) (userID string, err error) {
	store, err := session.Start(nil, w, r)
	if err != nil {
		return
	}

	uid, ok := store.Get("LoggedInUserID")
	fmt.Printf("userAuthorizeHandler:uid:%s" , uid)
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



