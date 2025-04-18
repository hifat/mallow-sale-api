package recipeHandler

import (
	"context"
	"net/http"

	"github.com/hifat/cost-calculator-api/internal/recipe"
	"github.com/hifat/cost-calculator-api/internal/recipe/recipeService"
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
		c.AbortWithJSON(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
		})

		return
	}

	res, err := h.recipeSrv.Create(context.Background(), req)
	if err != nil {
		c.AbortWithJSON(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, map[string]any{
		"item": res,
	})
}

func (h *recipeRest) Find(c core.IHttpCtx) {
	res, err := h.recipeSrv.Find(context.Background())
	if err != nil {
		c.AbortWithJSON(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"items": res,
	})
}

func (h *recipeRest) FindByID(c core.IHttpCtx) {
	recipeID := c.Param("recipeID")

	res, err := h.recipeSrv.FindByID(context.Background(), recipeID)
	if err != nil {
		c.AbortWithJSON(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"item": res,
	})
}

func (h *recipeRest) Update(c core.IHttpCtx) {
	recipeID := c.Param("recipeID")

	req := recipe.RecipeReq{}
	if err := c.BodyParser(&req); err != nil {
		c.AbortWithJSON(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
		})

		return
	}

	err := h.recipeSrv.Update(context.Background(), recipeID, req)
	if err != nil {
		c.AbortWithJSON(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"message": "ok",
	})
}

func (h *recipeRest) Delete(c core.IHttpCtx) {
	recipeID := c.Param("recipeID")

	err := h.recipeSrv.Delete(context.Background(), recipeID)
	if err != nil {
		c.AbortWithJSON(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"message": "ok",
	})
}
