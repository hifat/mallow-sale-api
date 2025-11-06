package userModule

import "golang.org/x/crypto/bcrypt"

type Prototype struct {
	Name     string `fake:"{name}" json:"name"`
	Username string `json:"username"`

	Password []byte `bson:"password" json:"-"`
}

type Response struct {
	Prototype
}

type Request struct {
	Password []byte `json:"password"`
}

func (r *Request) HashPassword() error {
	hashed, err := bcrypt.GenerateFromPassword(r.Password, 10)
	if err != nil {
		return err
	}

	r.Password = hashed

	return nil
}
