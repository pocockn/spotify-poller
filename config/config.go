package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
)

type (
	// Config contains config for the application.
	Config struct {
		Database Database
		Poller   Poller
		Spotify  Spotify
	}

	// Poller holds poller specific values
	Poller struct {
		Interval Duration
	}

	// Database holds database values in our config.
	Database struct {
		Host           string
		DatabaseName   string
		Port           string
		Password       string
		MaxConnections int
		Username       string
		URL            string
	}

	// Spotify holds Spotify API values.
	Spotify struct {
		ClientID     string
		ClientSecret string
		PlaylistID   string
		TokenURL     string
	}
)

// NewConfig creates a new config struct.
func NewConfig() Config {
	var config Config
	if _, err := toml.DecodeFile(config.generatePath(), &config); err != nil {
		fmt.Println(err)
	}

	config.Spotify.ClientSecret = os.Getenv("SPOTIFY_CLIENT_SECRET")

	return config
}

func (c Config) environment() string {
	environment := "development"

	if os.Getenv("ENV") != "" {
		environment = os.Getenv("ENV")
	}

	return environment
}

func (c Config) generatePath() string {
	if os.Getenv("ENV") == "test" {
		return "development.toml"
	}

	return fmt.Sprintf("config/%s.toml", c.environment())
}

func (c Config) isDev() bool {
	return os.Getenv("ENV") == "development"
}
