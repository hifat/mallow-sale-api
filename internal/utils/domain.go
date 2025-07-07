package utilsModule

type QueryReq struct {
	Page   int    `form:"page" json:"page" default:"1"`
	Limit  int    `form:"limit" json:"limit" default:"10"`
	Search string `form:"search" json:"search" default:""`
	Sort   string `form:"sort" json:"sort" default:"createdAt"`
	Order  string `form:"order" json:"order" default:"desc"`
	Fields string `form:"fields" json:"fields"`
}
