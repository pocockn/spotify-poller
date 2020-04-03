package internals

import (
	"github.com/pocockn/recs-api/models"
	internal_spotify "github.com/pocockn/spotify-poller/internals/spotify"
	"github.com/sirupsen/logrus"
	"github.com/zmb3/spotify"
)

type (
	// Handler holds the dependencies the handler function needs.
	Handler struct {
		client internal_spotify.SpotifyClient
		store  Storer
	}
)

// NewHandler creates a new handler struct.
func NewHandler(client internal_spotify.SpotifyClient, store Storer) Handler {
	return Handler{
		client: client,
		store:  store,
	}
}

// Spotify queries Spotify and checks if the supplied playlist ID
// has new songs, if it finds new songs it will add them to our database.
func (h *Handler) Spotify() error {
	logrus.Debugf("performing API request for playlist #%s", "01VmpQKq19m0CjP1bo1eMO")
	playlist, _ := h.client.GetPlaylist("01VmpQKq19m0CjP1bo1eMO")

	recs, err := h.fetchRecs()
	if err != nil {
		return err
	}

	err = h.addNewRecs(playlist, recs)
	if err != nil {
		return err
	}

	return nil
}

func (h *Handler) addNewRecs(playlist *spotify.FullPlaylist, existingRecs models.Recs) error {
	if len(existingRecs) == len(playlist.Tracks.Tracks) {
		return nil
	}

	for _, t := range playlist.Tracks.Tracks {
		if h.found(existingRecs, t.Track.ID.String()) {
			logrus.Debugf("track : %s with ID : %s already in database, skipping.", t.Track.Name, t.Track.ID.String())
			continue
		}

		logrus.Infof("adding new song %s from Spotify", t.Track.Name)
		rec := models.Rec{
			Rating:    0,
			Review:    "",
			SpotifyID: t.Track.ID.String(),
			Title:     t.Track.Name,
		}

		err := h.store.Create(&rec)
		if err != nil {
			return err
		}

		continue
	}

	return nil
}

func (h *Handler) found(slice models.Recs, val string) bool {
	for _, item := range slice {
		if item.SpotifyID == val {
			return true
		}
	}
	return false
}

func (h *Handler) fetchRecs() (models.Recs, error) {
	recs, err := h.store.FetchAll()
	if err != nil {
		return nil, err
	}

	return recs, nil
}
