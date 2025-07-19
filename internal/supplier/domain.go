package supplierModule

import "time"

type Request struct {
	Name   string `validate:"required" json:"name"`
	ImgUrl string `json:"imgUrl"`
}

type Prototype struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	ImgUrl    string     `json:"imgUrl"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

type Response struct {
	Prototype
}
