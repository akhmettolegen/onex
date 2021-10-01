package handlers

import (
	"github.com/akhmettolegen/onex/pkg/application"
)

// Handler model
type Handler struct {
	App application.Application
}

// Get - Handler initializer
func Get(app application.Application) *Handler {
	var handler Handler
	handler.App = app
	return &handler
}
