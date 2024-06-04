package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	v "sbebe.ch/baby-name-guesser/pkg/voters"
)

// AddNewVoter godoc
//
//	@Summary		Add a single new voter
//	@Description	Add a single new voter, used to register to the voting system
//	@Tags			voter
//	@Accept			json
//	@Produce		json
//	@Param			voter	body	v.VoterFull	true	"names"
//	@Success		200					{object}	v.Voter
//	@Failure		400					{object}	HTTPError
//	@Failure		404					{object}	HTTPError
//	@Failure		500					{object}	HTTPError
//	@Router			/voters [post]
func (c *Controller) AddNewVoter(ctx *gin.Context) {

	fmt.Println("Adding a new voter")
	var voter v.VoterFull

	if err := ctx.BindJSON(&voter); err != nil {
		// DO SOMETHING WITH THE ERROR
		fmt.Println(err)
	}

	result, err := v.AddVoter(voter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

// GetVoter godoc
//
//	@Summary		Get a single voter by email
//	@Description	Get a single voter by email
//	@Tags			voter
//	@Accept			json
//	@Produce		json
//	@Param			email	path	string	true	"Email of the voter"
//	@Success		200					{object}	v.Voter
//	@Failure		400					{object}	HTTPError
//	@Failure		404					{object}	HTTPError
//	@Failure		500					{object}	HTTPError
//	@Router			/voters/{email} [get]
func (c *Controller) GetVoter(ctx *gin.Context) {
	email := ctx.Param("email")
	voter, err := v.GetVoterByEmail(email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Voter not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, voter)
}

// DeleteVoter godoc
//
//	@Summary		Delete a single voter by email
//	@Description	Delete a single voter by email
//	@Tags			voter
//	@Accept			json
//	@Produce		json
//	@Param			email	path	string	true	"Email of the voter"
//	@Success		200					{string}	string
//	@Failure		400					{object}	HTTPError
//	@Failure		404					{object}	HTTPError
//	@Failure		500					{object}	HTTPError
//	@Router			/voters/{email} [delete]
func (c *Controller) DeleteVoter(ctx *gin.Context) {
	email := ctx.Param("email")
	err := v.DeleteVoterByEmail(email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Voter not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Voter deleted successfully",
	})
}

// LoginVoter godoc
//
//	@Summary		Login a voter
//	@Description	Login a voter
//	@Tags			voter
//	@Accept			json
//	@Produce		json
//	@Param			voter	body	v.VoterLogin	true	"names"
//	@Success		200					{string}	string
//	@Failure		400					{object}	HTTPError
//	@Failure		404					{object}	HTTPError
//	@Failure		500					{object}	HTTPError
//	@Router			/voters/login [post]
func (c *Controller) LoginVoter(ctx *gin.Context) {
	var voter v.VoterLogin
	if err := ctx.BindJSON(&voter); err != nil {
		// DO SOMETHING WITH THE ERROR
		fmt.Println(err)
	}
	result, err := v.LoginVoter(voter.Email, voter.Password)
	if err != nil || !result {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Login failed",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Voter logged in successfully",
	})
}
