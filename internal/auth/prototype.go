package authModule

import "github.com/google/uuid"

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
