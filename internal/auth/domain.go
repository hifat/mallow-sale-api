package authModule

import (
	"github.com/google/uuid"
	userModule "github.com/hifat/mallow-sale-api/internal/user"
)

type SigninReq struct {
	Username string `fake:"{username}" json:"username"`
	Password []byte `fake:"{password}" json:"password"`
}

type AuthRes struct {
	userModule.Prototype

	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type Passport struct {
	AuthID uuid.UUID `json:"authID"`
	User   *AuthRes  `json:"user"`
}
