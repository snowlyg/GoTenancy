// +build !enterprise

package enterprise

import (
	"fmt"

	"GoTenancy/config/application"
)

// New new home app
func New(config *Config) *App {
	return &App{Config: config}
}

// App home app
type App struct {
	Config *Config
}

// Config home config struct
type Config struct {
}

// ConfigureApplication configure application
func (App) ConfigureApplication(application *application.Application) {
	fmt.Println("Enterprise features not enabled...")
}
