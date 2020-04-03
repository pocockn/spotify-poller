package handler_test

import (
	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/pocockn/recs-api/models"
	handler2 "github.com/pocockn/spotify-poller/internals/handler"
	mock_spotify "github.com/pocockn/spotify-poller/internals/store/mock"
	"github.com/stretchr/testify/assert"
	spotify_api "github.com/zmb3/spotify"
	"testing"
)

type mockSpotifyClient struct{}

func (m *mockSpotifyClient) GetPlaylist(playlistID spotify_api.ID) (*spotify_api.FullPlaylist, error) {
	return &spotify_api.FullPlaylist{
		Tracks: spotify_api.PlaylistTrackPage{
			Tracks: tracks,
		},
	}, nil
}

var (
	tracks = []spotify_api.PlaylistTrack{
		spotify_api.PlaylistTrack{
			Track: spotify_api.FullTrack{
				SimpleTrack: spotify_api.SimpleTrack{Name: "Banger", ID: "1234"},
			},
		},
		spotify_api.PlaylistTrack{
			Track: spotify_api.FullTrack{
				SimpleTrack: spotify_api.SimpleTrack{Name: "Melter", ID: "12344"},
			},
		},
		spotify_api.PlaylistTrack{
			Track: spotify_api.FullTrack{
				SimpleTrack: spotify_api.SimpleTrack{Name: "Face Crusher", ID: "14234"},
			},
		}}
)

func TestHandlerCreatesNewRecs(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockStore := mock_spotify.NewMockStore(controller)

	mockStore.EXPECT().
		FetchAll().
		Return(
			models.Recs{
				models.Rec{
					Model:     gorm.Model{ID: 1},
					Rating:    2,
					Review:    "Meh",
					SpotifyID: "1234",
					Title:     "Hello",
				},
				models.Rec{
					Model:     gorm.Model{ID: 1},
					Rating:    2,
					Review:    "Meh",
					SpotifyID: "12344",
					Title:     "Hello",
				},
			},
			nil,
		)

	newRec := models.Rec{
		SpotifyID: "14234",
		Title:     "Face Crusher",
	}

	mockStore.EXPECT().Create(&newRec).Return(nil)

	client := &mockSpotifyClient{}

	handler := handler2.NewHandler(
		client,
		"XDR",
		mockStore,
	)

	assert.NoError(t, handler.Spotify())
}

func TestHandlerOnlyAddsRecsNotInDB(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockStore := mock_spotify.NewMockStore(controller)

	mockStore.EXPECT().
		FetchAll().
		Return(
			models.Recs{
				models.Rec{
					Model:     gorm.Model{ID: 1},
					Rating:    2,
					Review:    "Meh",
					SpotifyID: "1234",
					Title:     "Hello",
				},
				models.Rec{
					Model:     gorm.Model{ID: 1},
					Rating:    2,
					Review:    "Meh",
					SpotifyID: "12344",
					Title:     "Hello",
				},
				models.Rec{
					Model:     gorm.Model{ID: 1},
					Rating:    2,
					Review:    "Meh",
					SpotifyID: "14234",
					Title:     "Hello",
				},
			},
			nil,
		)

	client := &mockSpotifyClient{}

	handler := handler2.NewHandler(
		client,
		"XDR",
		mockStore,
	)

	assert.NoError(t, handler.Spotify())
}

func TestHandlerErrorOnFetch(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockStore := mock_spotify.NewMockStore(controller)

	mockStore.EXPECT().
		FetchAll().
		Return(
			models.Recs{
				models.Rec{
					Model:     gorm.Model{ID: 1},
					Rating:    2,
					Review:    "Meh",
					SpotifyID: "1234",
					Title:     "Hello",
				},
				models.Rec{
					Model:     gorm.Model{ID: 1},
					Rating:    2,
					Review:    "Meh",
					SpotifyID: "12344",
					Title:     "Hello",
				},
				models.Rec{
					Model:     gorm.Model{ID: 1},
					Rating:    2,
					Review:    "Meh",
					SpotifyID: "14234",
					Title:     "Hello",
				},
			},
			errors.New("Fetch error"),
		)

	client := &mockSpotifyClient{}

	handler := handler2.NewHandler(
		client,
		"XDR",
		mockStore,
	)

	assert.Error(t, handler.Spotify())
}

func TestHandlerErrorOnCreate(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockStore := mock_spotify.NewMockStore(controller)

	mockStore.EXPECT().
		FetchAll().
		Return(
			models.Recs{
				models.Rec{
					Model:     gorm.Model{ID: 1},
					Rating:    2,
					Review:    "Meh",
					SpotifyID: "1234",
					Title:     "Hello",
				},
				models.Rec{
					Model:     gorm.Model{ID: 1},
					Rating:    2,
					Review:    "Meh",
					SpotifyID: "12344",
					Title:     "Hello",
				},
			},
			nil,
		)

	newRec := models.Rec{
		SpotifyID: "14234",
		Title:     "Face Crusher",
	}

	mockStore.EXPECT().Create(&newRec).Return(errors.New("Create error."))

	client := &mockSpotifyClient{}

	handler := handler2.NewHandler(
		client,
		"XDR",
		mockStore,
	)

	assert.Error(t, handler.Spotify())
}
