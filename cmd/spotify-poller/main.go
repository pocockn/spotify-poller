package main

import (
	"github.com/pocockn/spotify-poller/config"
	"github.com/pocockn/spotify-poller/internals/database"
	"github.com/pocockn/spotify-poller/internals/handler"
	"github.com/pocockn/spotify-poller/internals/spotify"
	"github.com/pocockn/spotify-poller/internals/store"
	spotify_poller "github.com/pocockn/spotify-poller/pkg/poller"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	pollerConfig := config.NewConfig()
	connection := database.NewConnection(pollerConfig)

	client, err := spotify.NewClient(pollerConfig.Spotify)
	if err != nil {
		logrus.Fatal(err)
	}

	handler := handler.NewHandler(
		client,
		pollerConfig.Spotify.PlaylistID,
		store.NewStore(connection),
	)

	poller := spotify_poller.NewPoller(
		handler.Spotify,
		time.NewTicker(pollerConfig.Poller.Interval.Duration),
	)

	errChan := poller.Start()
	for err := range errChan {
		if err != nil {
			logrus.Fatalf("fatal error: %s", err)
		}
	}
}
