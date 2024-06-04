package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	v "sbebe.ch/baby-name-guesser/pkg/votes"
)

// AddNewVoter godoc
//
//	@Summary		Add a single new voter
//	@Description	Add a single new voter, used to register to the voting system
//	@Tags			voting
//	@Accept			json
//	@Produce		json
//	@Param			vote	body	v.Vote	true	"vote"
//	@Success		200					{object}	v.Vote
//	@Failure		400					{object}	HTTPError
//	@Failure		404					{object}	HTTPError
//	@Failure		500					{object}	HTTPError
//	@Router			/votes [post]
func (c *Controller) AddVotes(ctx *gin.Context) {

	fmt.Println("Adding a new vote")
	var voter v.Vote

	if err := ctx.BindJSON(&voter); err != nil {
		// DO SOMETHING WITH THE ERROR
		fmt.Println(err)
	}

	result, err := v.AddVotes(voter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
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
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
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

	result, err := v.GetVotesForVoters()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
