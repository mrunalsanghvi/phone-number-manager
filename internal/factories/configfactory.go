package factories

import (
	"phone-number-manager/internal/config"
)

// NewConfig loads the application configuration.
func NewConfig(path string) (*config.Config, error) {
	return config.LoadConfig(path)
}
