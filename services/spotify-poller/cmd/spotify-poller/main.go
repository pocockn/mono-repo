package main

import (
	"fmt"
	"time"

	"github.com/pocockn/mono-repo/pkg/logs"
	"github.com/pocockn/mono-repo/pkg/poller"
	"github.com/pocockn/mono-repo/services/spotify-poller/config"
	"github.com/pocockn/mono-repo/services/spotify-poller/internals/database"
	"github.com/pocockn/mono-repo/services/spotify-poller/internals/handler"
	"github.com/pocockn/mono-repo/services/spotify-poller/internals/spotify"
	"github.com/pocockn/mono-repo/services/spotify-poller/internals/store"
)

func main() {
	logs.New(logs.WithDebug(), logs.WithService("spotify-poller"), logs.WithVersion("0.1.0"))
	pollerConfig := config.NewConfig()
	connection := database.NewConnection(pollerConfig)

	client, err := spotify.NewClient(pollerConfig.Spotify)
	if err != nil {
		logs.Logger.Fatal().Err(err).Send()
	}

	h := handler.NewHandler(
		client,
		pollerConfig.Spotify.PlaylistID,
		store.NewStore(connection),
	)

	p := poller.NewPoller(
		h.Spotify,
		time.NewTicker(pollerConfig.Poller.Interval.Duration),
	)

	errChan := p.Start()
	for err := range errChan {
		if err != nil {
			logs.Logger.Fatal().Err(fmt.Errorf("fatal error within poller %w", err)).Send()
		}
	}
}
