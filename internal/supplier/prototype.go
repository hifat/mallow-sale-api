package supplierModule

import "time"

type Prototype struct {
	ID        string     `fake:"{uuid}" json:"id"`
	Name      string     `fake:"{name}" json:"name"`
	ImgUrl    string     `fake:"{url}" json:"imgUrl"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}
