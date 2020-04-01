package main

import (
	"github.com/pocockn/spotify-poller/config"
	"github.com/pocockn/spotify-poller/internals"
	"github.com/pocockn/spotify-poller/internals/database"
	"github.com/pocockn/spotify-poller/internals/spotify"
	"github.com/pocockn/spotify-poller/internals/store"
	spotify_poller "github.com/pocockn/spotify-poller/pkg/poller"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	config := config.NewConfig()
	connection := database.NewConnection(config)

	client, err := spotify.NewClient(config.Spotify)
	if err != nil {
		logrus.Fatal(err)
	}

	handler := internals.NewHandler(
		client,
		store.NewStore(connection),
	)

	poller := spotify_poller.NewPoller(
		handler.Spotify,
		time.NewTicker(config.Poller.Interval.Duration),
	)

	err = poller.Start()
	if err != nil {
		logrus.Fatalf("fatal error: %s", err)
	}
}
