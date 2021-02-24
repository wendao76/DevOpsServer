package model

type IAuthUser interface {
	GetUsername() string
	GetUserId() string
	GetUserType() string
}
//认证授权用户
type AuthUser struct {
	UserId string
	Username string
	UserType int
}

func (authUser * AuthUser) GetUsername() string{
	return authUser.Username
}

func (authUser * AuthUser) GetUserId() string{
	return authUser.UserId
}

func (authUser * AuthUser) GetUserType() string{
	return authUser.UserId
}


