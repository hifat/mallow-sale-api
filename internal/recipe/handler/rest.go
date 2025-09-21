package recipeHandler

import (
	"github.com/gin-gonic/gin"
	recipeModule "github.com/hifat/mallow-sale-api/internal/recipe"
	recipeService "github.com/hifat/mallow-sale-api/internal/recipe/service"
	"github.com/hifat/mallow-sale-api/pkg/handling"
)

type Rest struct {
	recipeService recipeService.IService
}

func NewRest(recipeService recipeService.IService) *Rest {
	return &Rest{recipeService: recipeService}
}

// @Summary 	Create Recipe
// @Tags 		recipe
// @Accept 		json
// @Produce 	json
// @Param 		recipe body recipeModule.Request true "Created recipe data"
// @Success 	201 {object} handling.ResponseItem[recipeModule.Request]
// @Failure 	400 {object} handling.ErrorResponse
// @Failure 	404 {object} handling.ErrorResponse
// @Failure 	500 {object} handling.ErrorResponse
// @Router 		/recipes [post]
func (r *Rest) Create(c *gin.Context) {
	var req recipeModule.Request
	if err := c.ShouldBindJSON(&req); err != nil {
		handling.ResponseFormErr(c, err)
		return
	}

	res, err := r.recipeService.Create(c.Request.Context(), &req)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseSuccess(c, *res)
}

// @Summary 	Find Recipes
// @Tags 		recipe
// @Accept 		json
// @Produce 	json
// @Param 		query query recipeModule.QueryReq false "Query parameters"
// @Success 	200 {object} handling.ResponseItems[recipeModule.Response]
// @Failure 	500 {object} handling.ErrorResponse
// @Router 		/recipes [get]
func (r *Rest) Find(c *gin.Context) {
	var query recipeModule.QueryReq
	if err := c.ShouldBindQuery(&query); err != nil {
		handling.ResponseFormErr(c, err)
		return
	}

	res, err := r.recipeService.Find(c.Request.Context(), &query)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseSuccess(c, *res)
}

// @Summary 	Find Recipe by ID
// @Tags 		recipe
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "recipeID"
// @Success 	200 {object} handling.ResponseItem[recipeModule.Response]
// @Failure 	400 {object} handling.ErrorResponse
// @Failure 	404 {object} handling.ErrorResponse
// @Failure 	500 {object} handling.ErrorResponse
// @Router 		/recipes/{id} [get]
func (r *Rest) FindByID(c *gin.Context) {
	id := c.Param("id")

	res, err := r.recipeService.FindByID(c.Request.Context(), id)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseSuccess(c, *res)
}

// @Summary 	Update Recipe by ID
// @Tags 		recipe
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "recipe ID"
// @Param 		recipe body recipeModule.Request true "Updated recipe data"
// @Success 	200 {object} handling.ResponseItem[recipeModule.Request]
// @Failure 	400 {object} handling.ErrorResponse
// @Failure 	404 {object} handling.ErrorResponse
// @Failure 	500 {object} handling.ErrorResponse
// @Router 		/recipes/{id} [put]
func (r *Rest) UpdateByID(c *gin.Context) {
	id := c.Param("id")

	var req recipeModule.Request
	if err := c.ShouldBindJSON(&req); err != nil {
		handling.ResponseFormErr(c, err)
		return
	}

	res, err := r.recipeService.UpdateByID(c.Request.Context(), id, &req)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseSuccess(c, *res)
}

// @Summary 	Delete Recipe by ID
// @Tags 		recipe
// @Accept 		json
// @Produce 	json
// @Param 		id 	path string true "Recipe ID"
// @Success 	200 {object} handling.SuccessResponse
// @Failure 	404 {object} handling.ErrorResponse
// @Failure 	500 {object} handling.ErrorResponse
// @Router 		/recipes/{id} [delete]
func (r *Rest) DeleteByID(c *gin.Context) {
	id := c.Param("id")

	res, err := r.recipeService.DeleteByID(c.Request.Context(), id)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseSuccess(c, res)
}

// @Summary      Batch update recipe order no
// @Tags         recipe
// @Accept       json
// @Produce      json
// @Param        body body []recipeModule.UpdateOrderNoRequest true "Array of recipe id and orderNo"
// @Success      200 {object} handling.SuccessResponse
// @Failure      400 {object} handling.ErrorResponse
// @Failure      500 {object} handling.ErrorResponse
// @Router       /recipes/order-no [patch]
func (r *Rest) PatchNoBatch(c *gin.Context) {
	var req []recipeModule.UpdateOrderNoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handling.ResponseFormErr(c, err)
		return
	}

	err := r.recipeService.UpdateNoBatch(c.Request.Context(), req)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseSuccess(c, nil)
}
