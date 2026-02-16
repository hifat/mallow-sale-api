package authModule

import userModule "github.com/hifat/mallow-sale-api/internal/user"

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
