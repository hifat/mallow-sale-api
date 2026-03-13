package userModule

type Prototype struct {
	ID       string `fake:"{uuid}" json:"id"`
	Name     string `fake:"{name}" json:"name"`
	Username string `json:"username"`

	Password string `bson:"password" json:"-"`
}
