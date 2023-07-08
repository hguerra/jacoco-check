package internal

import (
	"log"

	"github.com/hguerra/jacoco-check/pkg/config"
)

func NewParser(cfg *config.Config, args []string) {
	env := cfg.GetActiveEnv()
	log.Printf("Active env %s", env)

	isDevelopment := cfg.IsDevelopment()
	log.Printf("Is development %v %v", isDevelopment, args)
}
