// Package nrqldroprules provides a programmatic API for interacting configuring New Relc NRQL Drop Rules.
package nrqldroprules

import (
	"github.com/cegeka-everesst/newrelic-client-go/internal/http"
	"github.com/cegeka-everesst/newrelic-client-go/pkg/config"
	"github.com/cegeka-everesst/newrelic-client-go/pkg/logging"
)

// NrqlDropRules is used to interact with New Relic accounts.
type Nrqldroprules struct {
	client http.Client
	logger logging.Logger
}

// New returns a new client for interacting with New Relic accounts.
func New(config config.Config) Nrqldroprules {
	return Nrqldroprules{
		client: http.NewClient(config),
		logger: config.GetLogger(),
	}
}
