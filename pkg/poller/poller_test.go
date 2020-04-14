package poller_test

import (
	"github.com/pkg/errors"
	"github.com/pocockn/mono-repo/pkg/poller"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	called = false
)

func testHandlerFunc() error {
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

	poller.Start()
	poller.Stop()

	for err := range errChan {
		assert.NoError(t, err)
	}
}

func TestPollerError(t *testing.T) {
	poller := poller.NewPoller(
		testHandlerFuncError,
		time.NewTicker(1*time.Millisecond),
	)

	errChan := poller.Start()
	poller.Stop()

	for err := range errChan {
		assert.Error(t, err)
	}
}
