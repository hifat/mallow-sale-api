package authModule

type SigninReq struct {
	Username string `fake:"{username}" json:"username"`
	Password string `fake:"{password}" json:"password"`
}
