package poller

import (
	"github.com/sirupsen/logrus"
	"time"
)

type (
	// HandlerFunc is a function that is run each time the API is polled.
	HandlerFunc func() error

	// Poller will poll the API based on the interval passed in.
	Poller struct {
		HandlerFunc HandlerFunc
		interval    *time.Ticker
		Errs        chan error
		done        chan bool
	}
)

// NewPoller creates a new Poller struct.
func NewPoller(handlerFunc HandlerFunc, interval *time.Ticker) Poller {
	return Poller{
		HandlerFunc: handlerFunc,
		interval:    interval,
		Errs:        make(chan error),
		done:        make(chan bool, 0),
	}
}

// Start starts take the poller.
func (p *Poller) Start() <-chan error {
	errc := make(chan error, 1)

	go func() {
		logrus.Info("polling initialized. Status: Running")
		defer close(errc)
		for {
			select {
			case <-p.interval.C:
				err := p.HandlerFunc()
				if err != nil {
					errc <- err
					return
				}
			case <-p.done:
				logrus.Info("polling shutting down. Status: stopped.")
				return
			}
		}
	}()

	return errc
}

// Stop the poller.
func (p *Poller) Stop() {
	p.done <- true
}
