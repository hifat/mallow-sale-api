package supplierModule

type Request struct {
	Name   string `fake:"{name}" validate:"required" json:"name"`
	ImgUrl string `fake:"{url}" json:"imgUrl"`
}
