package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	n "sbebe.ch/baby-name-guesser/pkg/names"
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
	names, err := n.GetNames()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
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

	fmt.Println("AddBabyNames called")
	var names []string

	if err := ctx.BindJSON(&names); err != nil {
		// DO SOMETHING WITH THE ERROR
		fmt.Println(err)
	}

	names, err := n.AddNames(names)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
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
	name := ctx.Param("name")
	fmt.Printf("Deleting name: %v\n", name)
	err := n.DeleteName(name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, Message{Message: "Name deleted"})
}
