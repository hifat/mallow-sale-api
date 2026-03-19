package authModule

type SigninReq struct {
	Username  string           `fake:"{username}" validate:"required_if=LoginType INTERNAL" json:"username"`
	Password  string           `fake:"{password}" validate:"required_if=LoginType INTERNAL" json:"password"`
	Token     string           `fake:"{token}" validate:"required_if=LoginType GOOGLE" json:"token"`
	LoginType EnumCodeAuthType `fake:"{loginType}" validate:"required" json:"loginType"`
}
