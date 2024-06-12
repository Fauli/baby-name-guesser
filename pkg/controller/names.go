package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	n "sbebe.ch/baby-name-guesser/pkg/names"
	"sbebe.ch/baby-name-guesser/pkg/utils"
)

// GetAllBabyNames godoc
//
//	@Summary		Show all names
//	@Description	Return all of the stored names
//	@Tags			names
//	@Accept			json
//	@Produce		json
//	@Success		200					{object}	[]string
//	@Failure		400					{object}	HTTPError
//	@Failure		404					{object}	HTTPError
//	@Failure		500					{object}	HTTPError
//	@Router			/names [get]
func (c *Controller) GetAllBabyNames(ctx *gin.Context) {
	utils.Logger.Debug("GetAllBabyNames called")
	names, err := n.GetNames()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"names": names,
	})
}

// AddNewBabyNames godoc
//
//	@Summary		Add a list of new baby names
//	@Description	Add a list of new baby names to the store
//	@Tags			names
//	@Accept			json
//	@Produce		json
//	@Param			names	body	[]string	true	"names"
//	@Success		200					{object}	[]string
//	@Failure		400					{object}	HTTPError
//	@Failure		404					{object}	HTTPError
//	@Failure		500					{object}	HTTPError
//	@Router			/names [post]
func (c *Controller) AddBabyNames(ctx *gin.Context) {
	utils.Logger.Debug("AddBabyNames called")

	var names []string

	if err := ctx.BindJSON(&names); err != nil {
		ctx.JSON(http.StatusInternalServerError, HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	names, err := n.AddNames(names)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"names": names,
	})
}

// DeleteBabyName godoc
//
//	@Summary		Delete a baby name
//	@Description	Delete a specific baby name from the store
//	@Tags			names
//	@Accept			json
//	@Produce		json
//	@Param			name	path	string	true	"name"
//	@Success		200					{object}	Message
//	@Failure		400					{object}	HTTPError
//	@Failure		404					{object}	HTTPError
//	@Failure		500					{object}	HTTPError
//	@Router			/names/{name} [delete]
func (c *Controller) DeleteBabyName(ctx *gin.Context) {
	utils.Logger.Debug("DeleteBabyName called")

	name := ctx.Param("name")
	utils.Logger.Sugar().Infof("Deleting name: %v\n", name)

	err := n.DeleteName(name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, Message{Message: "Name deleted"})
}
