package http

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go_web/internal/common/config"
	oredis "gopkg.in/go-oauth2/redis.v3"
	"gopkg.in/oauth2.v3/generates"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
	"log"
)

func New() {
	//gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.LoadHTMLGlob("static/*")
	engine.Use(gin.Logger(), gin.Recovery())

	initRouter(engine)
	err := engine.Run(":9096")
	if err != nil {
		log.Fatal("用户服务启动失败：" + err.Error())
	}
	log.Println("Server is running at 9096 port")
}

//路由初始化
func initRouter(engine *gin.Engine) {
	srv := initOAuthServer()
	s := &OAuthAction{
		Srv: srv,
	}
	engine.GET("/login", s.LoginPage)
	engine.POST("/login", s.Login)
	engine.GET("/auth", s.AuthPage)
	engine.GET("/token", s.Token)
	engine.GET("/test", s.Test)
	engine.POST("/token", s.Token)
	engine.GET("/authorize", s.Authorize)
	engine.POST("/authorize", s.Authorize)
	engine.POST("/client-register", s.RegisterClient)
}

//初始化oauth2处理流程
func initOAuthServer() *server.Server {
	conf := config.GetInstance()
	redisConf := conf.Redis
	manager := manage.NewDefaultManager()
	// token store
	manager.MapTokenStorage(oredis.NewRedisStore(&redis.Options{
		Addr: redisConf.Addr,
		DB: redisConf.Db,
	}))

	// generate jwt access token
	manager.MapAccessGenerate(generates.NewJWTAccessGenerate([]byte("00000000"), jwt.SigningMethodHS512))

	clientStore := store.NewClientStore()
	clientStore.Set("222222", &models.Client{
		ID:     "222222",
		Secret: "22222222",
		Domain: "http://localhost:9094",
	})
	manager.MapClientStorage(clientStore)

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)
	srv.SetPasswordAuthorizationHandler(PasswordAuthorizationHandler)
	srv.SetAuthorizeScopeHandler(AuthorizeScopeHandler)
	srv.SetUserAuthorizationHandler(UserAuthorizeHandler)
	return srv
}
