package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/blendle/zapdriver"
	"github.com/gorilla/mux"
	"github.com/hsblhsn/hn.hsblhsn.me/backend"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql"
	"github.com/hsblhsn/hn.hsblhsn.me/frontend"
	"github.com/pkg/errors"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	app := fx.New(
		fx.Provide(newHTTPClient),
		fx.Provide(newRouter),
		fx.Provide(newLogger),
		graphql.Module(),
		backend.Module(),
		frontend.Module(),
		fx.Invoke(httpServer),
	)
	app.Run()
}

func newRouter() *mux.Router {
	return mux.NewRouter()
}

func newHTTPClient() *http.Client {
	return &http.Client{Transport: http.DefaultTransport}
}

func newLogger() (*zap.Logger, error) {
	logger, err := zapdriver.NewProduction()
	if err != nil {
		return nil, errors.Wrap(err, "main: failed to build logger")
	}
	return logger, nil
}

func httpServer(
	lifecycle fx.Lifecycle,
	shutdowner fx.Shutdowner,
	router *mux.Router,
	logger *zap.Logger,
) {
	server := &http.Server{
		Handler: router,
		Addr:    ":8080",
	}
	shutdown := func() {
		err := shutdowner.Shutdown()
		if err != nil {
			logger.Fatal("main: could not shutdown", zap.Error(err))
		}
	}
	OnStart := func(context.Context) error {
		go func() {
			logger.Info("starting http server", zap.String("addr", ":8080"))
			if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
				shutdown()
				return
			}
		}()
		return nil
	}
	OnStop := func(ctx context.Context) error {
		if err := server.Shutdown(ctx); err != nil {
			return errors.Wrap(err, "main:error shutting down server")
		}
		return nil
	}
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		val := <-sig
		logger.Info("main: received shutdown signal", zap.String("signal", val.String()))
		shutdown()
	}()

	lifecycle.Append(fx.Hook{
		OnStart: OnStart,
		OnStop:  OnStop,
	})
}
