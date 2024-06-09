package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	n "sbebe.ch/baby-name-guesser/pkg/names"
	"sbebe.ch/baby-name-guesser/pkg/utils"
	v "sbebe.ch/baby-name-guesser/pkg/votes"
)

// AddNewVoter godoc
//
//	@Summary		Add a single new voter
//	@Description	Add a single new voter, used to register to the voting system
//	@Tags			voting
//	@Accept			json
//	@Produce		json
//	@Param			vote	body	n.Names	true	"vote"
//	@Success		200					{object}	v.Vote
//	@Failure		400					{object}	HTTPError
//	@Failure		404					{object}	HTTPError
//	@Failure		500					{object}	HTTPError
//	@Router			/votes [post]
func (c *Controller) AddVotes(ctx *gin.Context) {

	fmt.Println("Adding a new vote")
	var votedNames n.Names

	if err := ctx.BindJSON(&votedNames); err != nil {
		ctx.JSON(http.StatusInternalServerError, HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	session, err := c.Store.Get(ctx.Request, "session")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	email := session.Values["email"]
	fmt.Printf("Adding votes for %s\n", email)
	result, err := v.AddVotes(session.Values["email"].(string), votedNames.Names)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

// GetAllVotes godoc
//
//	@Summary		Get all votes for the names
//	@Description	Get all votes for the names aggregated
//	@Tags			voting
//	@Accept			json
//	@Produce		json
//	@Success		200					{object}	map[string]int
//	@Failure		400					{object}	HTTPError
//	@Failure		404					{object}	HTTPError
//	@Failure		500					{object}	HTTPError
//	@Router			/votes [get]
func (c *Controller) GetAllVotes(ctx *gin.Context) {

	result, err := v.GetVotes()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

// GetVotesPerVoters godoc
//
//	@Summary		Get all votes for the voters
//	@Description	Get all votes for the voters aggregated
//	@Tags			voting
//	@Accept			json
//	@Produce		json
//	@Success		200					{object}	[]v.Vote
//	@Failure		400					{object}	HTTPError
//	@Failure		404					{object}	HTTPError
//	@Failure		500					{object}	HTTPError
//	@Router			/votes/voters [get]
func (c *Controller) GetVotesPerVoters(ctx *gin.Context) {

	session, err := c.Store.Get(ctx.Request, "session")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	// Only admin can view all votes
	email := session.Values["email"]
	if email != utils.GetAdminEmail() {
		ctx.JSON(http.StatusForbidden, HTTPError{
			Code:    http.StatusForbidden,
			Message: "You are not authorized to view this resource",
		})
		return
	}

	result, err := v.GetVotesForVoters()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

// GetTopVotes godoc
//
//	@Summary		Get the top votes
//	@Description	Get the top votes for the names
//	@Tags			voting
//	@Accept			json
//	@Produce		json
//	@Success		200					{object}	map[string]int
//	@Failure		400					{object}	HTTPError
//	@Failure		404					{object}	HTTPError
//	@Failure		500					{object}	HTTPError
//	@Router			/votes/top [get]
func (c *Controller) GetTopVotes(ctx *gin.Context) {

	result, err := v.GetTopVotes()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
