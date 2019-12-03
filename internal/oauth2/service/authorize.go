package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go_web/internal/oauth2/model"
	"gopkg.in/oauth2.v3/generates"
	"log"
)

var (
	OauthService *oauthService
)
var (
	gen      *generates.JWTAccessGenerate
	tokenKey = "00000000"
)

func init() {
	OauthService = &oauthService{}
	gen = generates.NewJWTAccessGenerate([]byte(tokenKey), jwt.SigningMethodHS512)
}

type oauthService struct {
}

//生成jwt access token
func (s *oauthService) GenJWTAccessToken(data model.ClaimsLocal, sClaims jwt.StandardClaims) (string, error) {
	claims := &model.JWTAccessClaimsLocal{
		Username:       data.GetUsername(),
		Uid:            data.GetUid(),
		StandardClaims: sClaims,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte("00000000"))
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

//解析 jwt对象
func (s *oauthService) ParseJWTAccessToken(accessToken string) (model.ClaimsLocal, error) {
	newClaims := &model.JWTAccessClaimsLocal{}
	token, err := jwt.ParseWithClaims(accessToken, newClaims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("parse error")
		}
		return []byte(tokenKey), nil
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	claimsResult, _ := token.Claims.(*model.JWTAccessClaimsLocal)
	return claimsResult, nil
}
