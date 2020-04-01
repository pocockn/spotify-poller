package poller_test

import (
	"github.com/pkg/errors"
	poller "github.com/pocockn/spotify-poller/pkg/poller"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	called = false
)

func testHandlerFunc() error {
	called = true
	return nil
}

func testHandlerFuncError() error {
	return errors.New("poller error")
}

func TestPoller(t *testing.T) {
	poller := poller.NewPoller(
		testHandlerFunc,
		time.NewTicker(1*time.Millisecond),
	)

	assert.NoError(t, poller.Start())
	assert.Equal(t, true, called)
}

func TestPollerError(t *testing.T) {
	poller := poller.NewPoller(
		testHandlerFuncError,
		time.NewTicker(1*time.Millisecond),
	)

	assert.Error(t, poller.Start())
}
