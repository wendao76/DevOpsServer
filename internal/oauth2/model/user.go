package model


type OAuthUser struct {
	Uid uint
    Username string
	ExpiresAt int64
}

func (a *OAuthUser) GetUid() uint {
	return a.Uid
}

func (a *OAuthUser) GetUsername() string{
	return a.Username
}

func (a *OAuthUser) GetExpiresAt() int64 {
	return a.ExpiresAt
}
