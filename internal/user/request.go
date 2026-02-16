package userModule

import "golang.org/x/crypto/bcrypt"

type Request struct {
	Password string `json:"password"`
}

func (r *Request) HashPassword() error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(r.Password), 10)
	if err != nil {
		return err
	}

	r.Password = string(hashed)

	return nil
}
