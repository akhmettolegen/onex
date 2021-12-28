package manager

import "github.com/akhmettolegen/texert/pkg/application"

type Manager struct {
	App *application.Application
}

// Get - creates new Manager instance
func Get(app *application.Application) (*Manager, error) {
	return &Manager{App: app}, nil
}
