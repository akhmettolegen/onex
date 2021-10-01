package handlers

import (
	"github.com/akhmettolegen/onex/internal/manager"
	"github.com/akhmettolegen/onex/pkg/application"
)

// Handler model
type Handler struct {
	App application.Application
	Manager *manager.Manager
}

// Get - Handler initializer
func Get(app application.Application) *Handler {
	manager, _ := manager.Get(&app)

	return &Handler{
		App:     app,
		Manager: manager,
	}
}


