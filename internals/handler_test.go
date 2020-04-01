package internals_test

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/pocockn/spotify-poller/internals"
	mock_spotify "github.com/pocockn/spotify-poller/internals/store/mock"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
	"testing"
)

var ()

func TestHandler(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockStore := mock_spotify.NewMockStore(controller)

	config := &clientcredentials.Config{
		ClientID:     "www",
		ClientSecret: "eeee",
		TokenURL:     "spotify",
	}

	token, _ := config.Token(context.Background())

	handler := internals.NewHandler(
		spotify.Authenticator{}.NewClient(token),
		mockStore,
	)

	// Need to mock the recs and spotify calls and assert what comes back.
}

// Test for error.
