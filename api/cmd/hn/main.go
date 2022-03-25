package main

import (
	"github.com/gorilla/mux"
	"github.com/hsblhsn/hn.hsblhsn.me/api"
	"github.com/hsblhsn/hn.hsblhsn.me/api/internal/grpc/readabilityserver"
	"github.com/hsblhsn/hn.hsblhsn.me/api/internal/servers"
	"github.com/hsblhsn/hn.hsblhsn.me/api/internal/spa"
	"github.com/hsblhsn/hn.hsblhsn.me/embedded"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer func() {
		err := logger.Sync()
		if err != nil {
			panic(err)
		}
	}()
	zap.ReplaceGlobals(logger)
	logger.Info("main: starting server...")
	go readabilityserver.Initialize()

	// Start the server.
	router := mux.NewRouter()

	api.RegisterRoutes(router.PathPrefix("/api/v1").Subrouter())
	spa.RegisterRoutes(router, embedded.Assets)

	if err := servers.Start(router); err != nil {
		logger.Fatal("main: could not start server", zap.Error(err))
	}
}
