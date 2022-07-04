package logs

import (
	"github.com/rs/zerolog"
	"os"
)

// Logger is the global zerolog logger which we can customise in our constructors.
// Usage - logs.Logger.Info().Msgf("%d attempt at connecting to the DB", i)
var Logger zerolog.Logger

// New sets the global log level and creates the logger, directing logs to Stderr.
// It expects the version number and service name to attach additional config to the logs.
func New(opts ...NewFuncOption) {
	logLevel := zerolog.InfoLevel
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	zerolog.SetGlobalLevel(logLevel)

	for _, o := range opts {
		logger = o(&logger)
	}

	Logger = logger
}

// NewFuncOption is a functional option for the New function.
type NewFuncOption func(logger *zerolog.Logger) zerolog.Logger

// WithDebug sets the log level to debug. Used in all environments except production.
func WithDebug() NewFuncOption {
	return func(zl *zerolog.Logger) zerolog.Logger {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		return *zl
	}
}

// WithVersion sets the version of the service in the logs.
func WithVersion(version string) NewFuncOption {
	return func(zl *zerolog.Logger) zerolog.Logger {
		return zl.With().Str("version", version).Logger()
	}
}

// WithService adds the service name to the logs.
func WithService(service string) NewFuncOption {
	return func(zl *zerolog.Logger) zerolog.Logger {
		return zl.With().Str("service", service).Logger()
	}
}
