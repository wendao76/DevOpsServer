package model

import (
    "github.com/dgrijalva/jwt-go"
    "gopkg.in/oauth2.v3/errors"
    "time"
)

//JWT内容主体
type JWTAccessClaimsLocal struct {
    StandardClaims jwt.StandardClaims
    Uid uint `json:uid`
    Username string `json:username`
}

//jwt传递其他数据字段
type ClaimsLocal interface {
    //获取用户ID
    GetUid() uint
    //获取用户帐号
    GetUsername() string
    //获取过期时间
    GetExpiresAt() int64
}

func (a *JWTAccessClaimsLocal) GetUid() uint {
    return a.Uid
}

func (a *JWTAccessClaimsLocal) GetUsername() string{
    return a.Username
}

func (a *JWTAccessClaimsLocal) GetExpiresAt() int64 {
    return a.StandardClaims.ExpiresAt
}

//token可用性校验
func (a *JWTAccessClaimsLocal) Valid() error {
    if time.Unix(a.StandardClaims.ExpiresAt, 0).Before(time.Now()) {
        return errors.ErrInvalidAccessToken
    }
    return nil
}

