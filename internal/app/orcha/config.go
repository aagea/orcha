package orcha

import "github.com/rs/zerolog/log"

// Config structure with all the options required by the service and service components.
type Config struct {
	Version string
	Commit  string
}

// IsValid checks if the configuration options are valid.
func (c *Config) IsValid() error {
	return nil
}

// Print the configuration using the application logger.
func (c *Config) Print() {
	// Use logger to print the configuration
	log.Info().Str("version", c.Version).Str("commit", c.Commit).Msg("Orcha config")
}
