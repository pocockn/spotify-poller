package spotify

import (
	"context"
	"github.com/pkg/errors"
	"github.com/pocockn/spotify-poller/config"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

// Client is a wrapper around the SpotifyClient interface.
type Client struct {
	SpotifyClient
}

// Client is an interface that allows us to mock the spotify client in testing.
type SpotifyClient interface {
	GetPlaylist(playlistID spotify.ID) (*spotify.FullPlaylist, error)
}

// NewClient creates a new Spotify client.
func NewClient(config config.Spotify) (SpotifyClient, error) {
	oauthConfig := &clientcredentials.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		TokenURL:     spotify.TokenURL,
	}

	token, err := oauthConfig.Token(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "problem creating oauth token")
	}

	client := spotify.Authenticator{}.NewClient(token)

	return &client, nil
}

// GetPlaylist takes a spotify client and a playlist ID, it takes a client so we can mock
// a fake client in the tests.
func (c *Client) GetPlaylist(playlistID spotify.ID) (*spotify.FullPlaylist, error) {
	result, err := c.SpotifyClient.GetPlaylist(playlistID)
	if err != nil {
		return nil, errors.Wrap(err, "problem getting playlist")
	}

	return result, nil
}
