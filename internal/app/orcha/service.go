package orcha

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"
)

// Service structure in charge of launching the application.
type Service struct {
	cfg Config
}

// NewService creates a new service with a given configuration
func NewService(cfg Config) *Service {
	return &Service{
		cfg: cfg,
	}
}

// Run method starting the internal components and launching the service
func (s *Service) Run() {
	if err := s.cfg.IsValid(); err != nil {
		log.Fatal().Err(err).Msg("invalid configuration options")
	}
	s.cfg.Print()
	s.registerShutdownListener()
	// Substitute ticker loop with proper code
	for now := range time.Tick(time.Minute) {
		fmt.Println(now, "alive")
	}
}

func (s *Service) registerShutdownListener() {
	osChannel := make(chan os.Signal)
	signal.Notify(osChannel, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-osChannel
		s.Shutdown()
		os.Exit(1)
	}()
}

// Shutdown code
func (s *Service) Shutdown() {
	log.Warn().Msg("shutting down service")
}
