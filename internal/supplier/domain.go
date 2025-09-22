package supplierModule

import "time"

type Request struct {
	Name   string `fake:"{name}" validate:"required" json:"name"`
	ImgUrl string `fake:"{url}" json:"imgUrl"`
}

type Prototype struct {
	ID        string     `fake:"{uuid}" json:"id"`
	Name      string     `fake:"{name}" json:"name"`
	ImgUrl    string     `fake:"{url}" json:"imgUrl"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

type Response struct {
	Prototype
}
