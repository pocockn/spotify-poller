package poller

import (
	"github.com/sirupsen/logrus"
	"time"
)

type (
	// HandlerFunc is a function that is run each time the Spotify API is polled.
	HandlerFunc func() error

	// Poller will poll the API based on the interval passed in.
	Poller struct {
		HandlerFunc HandlerFunc
		interval    *time.Ticker
	}
)

// NewPoller creates a new Poller struct.
func NewPoller(handlerFunc HandlerFunc, interval *time.Ticker) Poller {
	return Poller{
		HandlerFunc: handlerFunc,
		interval:    interval,
	}
}

// Starts take the poller.
func (p *Poller) Start() error {
	logrus.Info("polling initialized. Status: Running")

	for {
		select {
		case <-p.interval.C:
			err := p.HandlerFunc()
			if err != nil {
				return err
			}
		}
	}
}
