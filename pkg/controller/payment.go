package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"sbebe.ch/baby-name-guesser/pkg/payment"
	"sbebe.ch/baby-name-guesser/pkg/voters"
)

// GetAllBabyNames godoc
//
//	@Summary		Get the payment information
//	@Description	Get the payment information on how to transfer the bet money
//	@Tags			payments
//	@Accept			json
//	@Produce		json
//	@Success		200					{object}	payment.Payment
//	@Failure		400					{object}	HTTPError
//	@Failure		404					{object}	HTTPError
//	@Failure		500					{object}	HTTPError
//	@Router			/payments [get]
func (c *Controller) GetPayment(ctx *gin.Context) {

	payment := payment.GetPaymentInfo()
	ctx.JSON(http.StatusOK, payment)
}

// TODO: [franz] refactor this to /payments/{email}

// GetTopVotes godoc
//
//	@Summary		Set the payment status for the voter to true
//	@Description	Set the payment status for the voter to true
//	@Tags			payments
//	@Accept			json
//	@Produce		json
//	@Param			email	path	string	true	"Email of the voter"
//	@Success		200					{object}	Message
//	@Failure		400					{object}	HTTPError
//	@Failure		404					{object}	HTTPError
//	@Failure		500					{object}	HTTPError
//	@Router			/payments/{email}  [post]
func (c *Controller) PayForVotes(ctx *gin.Context) {

	voter := ctx.Param("email")

	err := voters.PayForVotes(voter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	message := Message{
		Message: fmt.Sprintf("Payment for %s has been processed", voter),
	}
	ctx.JSON(http.StatusOK, message)
}

// GetTopVotes godoc
//
//	@Summary		Gets the payment status for the voters votes
//	@Description	Gets the payment status for the voters votes
//	@Tags			payments
//	@Accept			json
//	@Produce		json
//	@Param			email	path	string	true	"Email of the voter"
//	@Success		200					{object}	bool
//	@Failure		400					{object}	HTTPError
//	@Failure		404					{object}	HTTPError
//	@Failure		500					{object}	HTTPError
//	@Router			/payments/{email}  [get]
func (c *Controller) HasUserPaid(ctx *gin.Context) {

	voter := ctx.Param("email")

	hasPaid, err := voters.HasUserPaid(voter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, hasPaid)
}
