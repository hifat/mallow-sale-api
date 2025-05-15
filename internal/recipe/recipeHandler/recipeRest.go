package recipeHandler

import (
	"net/http"

	"github.com/hifat/cost-calculator-api/internal/recipe"
	"github.com/hifat/cost-calculator-api/internal/recipe/recipeService"
	"github.com/hifat/cost-calculator-api/pkg/utils/handlerUtils.go"
	core "github.com/hifat/goroger-core"
)

type recipeRest struct {
	recipeSrv recipeService.IRecipeService
}

func NewRest(recipeSrv recipeService.IRecipeService) *recipeRest {
	return &recipeRest{recipeSrv}
}

func (h *recipeRest) Create(c core.IHttpCtx) {
	req := recipe.RecipeReq{}
	if err := c.BodyParser(&req); err != nil {
		handlerUtils.ResponseBadRequest(c, err)

		return
	}

	res, err := h.recipeSrv.Create(c.Context(), req)
	if err != nil {
		handlerUtils.ResponseErr(c, err)

		return
	}

	c.JSON(http.StatusCreated, map[string]any{
		"item": res,
	})
}

func (h *recipeRest) Find(c core.IHttpCtx) {
	res, err := h.recipeSrv.Find(c.Context())
	if err != nil {
		handlerUtils.ResponseErr(c, err)

		return
	}

	handlerUtils.ResponseItems(c, res)
}

func (h *recipeRest) FindByID(c core.IHttpCtx) {
	recipeID := c.Param("recipeID")

	res, err := h.recipeSrv.FindByID(c.Context(), recipeID)
	if err != nil {
		handlerUtils.ResponseErr(c, err)
		return
	}

	handlerUtils.ResponseItem(c, res)
}

func (h *recipeRest) Update(c core.IHttpCtx) {
	recipeID := c.Param("recipeID")

	req := recipe.UpdateRecipeReq{}
	if err := c.BodyParser(&req); err != nil {
		handlerUtils.ResponseBadRequest(c, err)

		return
	}

	err := h.recipeSrv.Update(c.Context(), recipeID, req)
	if err != nil {
		handlerUtils.ResponseErr(c, err)

		return
	}

	handlerUtils.ResponseOK(c)
}

func (h *recipeRest) Delete(c core.IHttpCtx) {
	recipeID := c.Param("recipeID")

	err := h.recipeSrv.Delete(c.Context(), recipeID)
	if err != nil {
		handlerUtils.ResponseBadRequest(c, err)

		return
	}

	handlerUtils.ResponseOK(c)
}
