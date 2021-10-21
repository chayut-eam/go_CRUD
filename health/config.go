package health

import (
	"context"
	"net/http"
	"time"

	"crud/db"
	log "crud/logger"
	"crud/model"

	"github.com/alexliesenfeld/health"
)

var (
	checker health.Checker
)

func Init(config model.HealthCheckConfig) {
	checker = health.NewChecker(
		health.WithCacheDuration(time.Millisecond*config.CacheDuration),
		health.WithPeriodicCheck(time.Millisecond*config.RefreshInterval, time.Millisecond*config.InitialDelay, mongoDBHealthCheck()),
		health.WithStatusListener(func(ctx context.Context, state health.CheckerState) {
			log.Logger().Infof("Overall system health status has changed to %v.", state.Status)
		}),
	)
}

func mongoDBHealthCheck() health.Check {
	return health.Check{
		Name: "MongoDB",
		Check: func(ctx context.Context) error {
			return db.Ping()
		},
		StatusListener: func(ctx context.Context, name string, state health.CheckState) {
			log.Logger().Infof("Status of %v has changed to %v.", name, state.Status)
		},
	}
}

func HealthCheckHandler() http.HandlerFunc {
	return health.NewHandler(checker, health.WithResultWriter(&JSONResultWriter{}))
}
