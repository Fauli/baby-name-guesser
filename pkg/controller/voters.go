package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sbebe.ch/baby-name-guesser/pkg/utils"
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
	utils.Logger.Debug("AddNewVoter called")

	var voter v.VoterFull
	voter.HasVoted = false

	if err := ctx.BindJSON(&voter); err != nil {
		ctx.JSON(http.StatusInternalServerError, HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	result, err := v.AddVoter(voter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	// TODO [franz] clean this up
	login, err := v.LoginVoter(voter.Email, voter.Password)
	if err != nil || !login {
		ctx.JSON(http.StatusForbidden, HTTPError{
			Code:    http.StatusForbidden,
			Message: "Login failed",
		})
		return
	}

	session, err := c.Store.Get(ctx.Request, "session")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	name, lastName, err := v.GetNameAndLastname(voter.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	session.Values["authenticated"] = true
	session.Values["email"] = voter.Email
	session.Values["name"] = name
	session.Values["last_name"] = lastName
	session.Save(ctx.Request, ctx.Writer)

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
	utils.Logger.Debug("GetVoter called")

	email := ctx.Param("email")
	voter, err := v.GetVoterByEmail(email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
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
	utils.Logger.Debug("DeleteVoter called")

	session, err := c.Store.Get(ctx.Request, "session")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	// Only admin can delete voters
	loggedInMail := session.Values["email"]
	if loggedInMail != utils.GetAdminEmail() {
		ctx.JSON(http.StatusForbidden, HTTPError{
			Code:    http.StatusForbidden,
			Message: "You are not authorized to view this resource",
		})
		return
	}

	email := ctx.Param("email")
	err = v.DeleteVoterByEmail(email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
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
	utils.Logger.Debug("LoginVoter called")

	var voter v.VoterLogin
	if err := ctx.BindJSON(&voter); err != nil {
		ctx.JSON(http.StatusInternalServerError, HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	result, err := v.LoginVoter(voter.Email, voter.Password)
	if err != nil || !result {
		ctx.JSON(http.StatusForbidden, HTTPError{
			Code:    http.StatusForbidden,
			Message: "Login failed",
		})
		return
	}

	session, err := c.Store.Get(ctx.Request, "session")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	name, lastName, err := v.GetNameAndLastname(voter.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	session.Values["authenticated"] = true
	session.Values["email"] = voter.Email
	session.Values["name"] = name
	session.Values["last_name"] = lastName
	session.Save(ctx.Request, ctx.Writer)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Voter logged in successfully",
	})
}

// LogoutVoter godoc
//
//	@Summary		LogoutVoter a voter
//	@Description	LogoutVoter a voter
//	@Tags			voter
//	@Accept			json
//	@Produce		json
//	@Success		200					{string}	Message
//	@Failure		400					{object}	HTTPError
//	@Failure		404					{object}	HTTPError
//	@Failure		500					{object}	HTTPError
//	@Router			/voters/logout [post]
func (c *Controller) LogoutVoter(ctx *gin.Context) {
	utils.Logger.Debug("LogoutVoter called")

	session, err := c.Store.Get(ctx.Request, "session")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	user := session.Values["email"]
	utils.Logger.Sugar().Info("Logging out user %s\n", user)

	session.Values["authenticated"] = false
	session.Values["email"] = ""
	session.Save(ctx.Request, ctx.Writer)

	ctx.JSON(http.StatusOK, Message{
		Message: "Voter logged out successfully",
	})
}
