package main

import (
	spotify_poller "github.com/pocockn/mono-repo/pkg/poller"
	"github.com/pocockn/mono-repo/services/spotify-poller/config"
	"github.com/pocockn/mono-repo/services/spotify-poller/internals/database"
	"github.com/pocockn/mono-repo/services/spotify-poller/internals/handler"
	"github.com/pocockn/mono-repo/services/spotify-poller/internals/spotify"
	"github.com/pocockn/mono-repo/services/spotify-poller/internals/store"
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

	h := handler.NewHandler(
		client,
		pollerConfig.Spotify.PlaylistID,
		store.NewStore(connection),
	)

	poller := spotify_poller.NewPoller(
		h.Spotify,
		time.NewTicker(pollerConfig.Poller.Interval.Duration),
	)

	errChan := poller.Start()
	for err := range errChan {
		if err != nil {
			logrus.Fatalf("fatal error within poller: %s", err)
		}
	}
}
