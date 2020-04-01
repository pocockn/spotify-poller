package spotify

import (
	"context"
	"github.com/pkg/errors"
	"github.com/pocockn/spotify-poller/config"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

// NewClient creates a new Spotify client.
func NewClient(config config.Spotify) (spotify.Client, error) {
	oauthConfig := &clientcredentials.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		TokenURL:     spotify.TokenURL,
	}

	token, err := oauthConfig.Token(context.Background())
	if err != nil {
		return spotify.Client{}, errors.Wrap(err, "problem creating oauth token")
	}

	return spotify.Authenticator{}.NewClient(token), nil
}
