package poller

import (
	"github.com/sirupsen/logrus"
	"time"
)

type (
	// HandlerFunc is a function that is run each time allotted interval is up.
	HandlerFunc func() error

	// Poller will execute the HandlerFunc every x number of minutes/seconds/hours
	// this is based on the interval passed in.
	Poller struct {
		HandlerFunc HandlerFunc
		Errs        chan error
		done        chan bool
		interval    *time.Ticker
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
	go func() {
		logrus.Info("polling initialized. Status: Running")
		defer close(p.Errs)
		for {
			select {
			case <-p.interval.C:
				err := p.HandlerFunc()
				if err != nil {
					p.Errs <- err
					return
				}
			case <-p.done:
				logrus.Info("polling shutting down. Status: stopped.")
				return
			}
		}
	}()

	return p.Errs
}

// Stop the poller.
func (p *Poller) Stop() {
	p.done <- true
}
