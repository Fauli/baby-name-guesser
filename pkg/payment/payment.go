package payment

import "sbebe.ch/baby-name-guesser/pkg/utils"

type Payment struct {
	PayPal  string `json:"paypal"`
	Twint   string `json:"twint"`
	Revolut string `json:"revolut"`
	IBAN    string `json:"iban"`
}

// GetPaymentInfo returns the payment information for voters.
func GetPaymentInfo() Payment {

	return Payment{
		PayPal:  utils.GetPaypal(),
		Twint:   utils.GetMobile(),
		Revolut: utils.GetMobile(),
		IBAN:    utils.GetIBAN(),
	}

}
