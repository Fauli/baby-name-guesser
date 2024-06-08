package controller

import "github.com/gorilla/sessions"

// Controller example
type Controller struct {
	Store *sessions.CookieStore
}

// NewController example
func NewController(store *sessions.CookieStore) *Controller {
	return &Controller{
		Store: store,
	}
}

// Message example
type Message struct {
	Message string `json:"message" example:"message"`
}

// HTTPError example
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}
