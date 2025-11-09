package userModule

import "golang.org/x/crypto/bcrypt"

type Prototype struct {
	ID       string `fake:"{uuid}" json:"id"`
	Name     string `fake:"{name}" json:"name"`
	Username string `json:"username"`

	Password string `bson:"password" json:"-"`
}

type Response struct {
	Prototype
}

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
