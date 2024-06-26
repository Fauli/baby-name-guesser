package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	n "sbebe.ch/baby-name-guesser/pkg/names"
	"sbebe.ch/baby-name-guesser/pkg/utils"
	"sbebe.ch/baby-name-guesser/pkg/voters"
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
	utils.Logger.Debug("AddVotes called")

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
	utils.Logger.Sugar().Infof("Adding votes for %s", email)
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
	utils.Logger.Debug("GetAllVotes called")

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
//	@Router			/votes/me [get]
func (c *Controller) GetPersonalVotes(ctx *gin.Context) {
	utils.Logger.Debug("GetPersonalVotes called")

	session, err := c.Store.Get(ctx.Request, "session")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	// Only admin can view others votes
	email := session.Values["email"].(string)
	result, err := v.GetVotesForVoter(email)
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
	utils.Logger.Debug("GetVotesPerVoters called")

	session, err := c.Store.Get(ctx.Request, "session")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	// Only admin can view others votes
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
//	@Success		200					{object}	[]string
//	@Failure		400					{object}	HTTPError
//	@Failure		404					{object}	HTTPError
//	@Failure		500					{object}	HTTPError
//	@Router			/votes/top [get]
func (c *Controller) GetTopVotes(ctx *gin.Context) {
	utils.Logger.Debug("GetTopVotes called")

	session, err := c.Store.Get(ctx.Request, "session")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	email := session.Values["email"]
	if email == "" {
		ctx.JSON(http.StatusForbidden, HTTPError{
			Code:    http.StatusForbidden,
			Message: "You are not authorized to view this resource",
		})
		return
	}
	hasVoted, err := voters.HasUserVoted(email.(string))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	// Only voters that have alredy voted can view the top votes
	if !hasVoted && email != utils.GetAdminEmail() {
		ctx.JSON(http.StatusForbidden, HTTPError{
			Code:    http.StatusForbidden,
			Message: "You are not authorized to view this resource",
		})
		return
	}

	var topVoted []string
	result, err := v.GetTopVotes()
	// only return the key values in the same order  as in the results slide
	for _, name := range result {
		topVoted = append(topVoted, name.Key)
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, topVoted)
}
