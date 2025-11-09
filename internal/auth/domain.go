package authModule

import (
	"github.com/google/uuid"
	userModule "github.com/hifat/mallow-sale-api/internal/user"
)

type SigninReq struct {
	Username string `fake:"{username}" json:"username"`
	Password string `fake:"{password}" json:"password"`
}

type AuthRes struct {
	userModule.Prototype

	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func (a *AuthRes) SetAccessToken(t string) {
	a.AccessToken = t
}

func (a *AuthRes) GetAccessToken() string {
	return a.AccessToken
}

func (a *AuthRes) SetRefreshToken(t string) {
	a.RefreshToken = t
}

func (a *AuthRes) GetUserID() string {
	return a.ID
}

type Passport struct {
	AuthID uuid.UUID `json:"authID"`
	User   *AuthRes  `json:"user"`
}

func (a *Passport) SetAccessToken(t string) {
	if a.User != nil {
		a.User.SetAccessToken(t)
	}
}

func (a *Passport) GetAccessToken() string {
	if a.User != nil {
		a.User.GetAccessToken()
	}

	return ""
}

func (a *Passport) SetRefreshToken(t string) {
	if a.User != nil {
		a.User.SetRefreshToken(t)
	}
}

func (a *Passport) GetUserID() string {
	if a.User != nil {
		a.User.GetAccessToken()
	}

	return ""
}
