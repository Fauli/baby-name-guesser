package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sbebe.ch/baby-name-guesser/pkg/utils"
	v "sbebe.ch/baby-name-guesser/pkg/voters"
)

// GetAllBabyNames godoc
//
//	@Summary		Get current logged in user information
//	@Description	Get current logged in user information
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Success		200					{object}	[]string
//	@Failure		400					{object}	HTTPError
//	@Failure		404					{object}	HTTPError
//	@Failure		500					{object}	HTTPError
//	@Router			/me [get]
func (c *Controller) GetUserInformation(ctx *gin.Context) {
	utils.Logger.Debug("GetUserInformation called")

	session, err := c.Store.Get(ctx.Request, "session")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	email := session.Values["email"].(string)

	voter, err := v.GetVoterByEmail(email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	hasVoted, err := v.HasUserVoted(email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	voter.HasVoted = hasVoted

	ctx.JSON(http.StatusOK, voter)
}
