package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"path"
	"runtime/debug"
	"syscall"
	"time"

	"cloud.google.com/go/profiler"
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
		fx.Invoke(pprofiler),
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
	const defaultTimeout = 15 * time.Second
	const defaultTimeoutResponse = `{"errors": {"message": "Server timeout"}}`
	server := &http.Server{
		Handler: http.TimeoutHandler(router, defaultTimeout, defaultTimeoutResponse),
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

func pprofiler(logger *zap.Logger) {
	svc, rev := parseBuildInfo()
	cfg := profiler.Config{
		Service:        svc,
		ServiceVersion: rev,
	}
	logger.Info(
		"main: starting profiler",
		zap.String("service", cfg.Service),
		zap.String("version", cfg.ServiceVersion),
	)
	if err := profiler.Start(cfg); err != nil {
		logger.Error("main: could not start profiler", zap.Error(err))
		return
	}
}

func parseBuildInfo() (service, revision string) {
	defaultMod := "hn.hsblhsn.me"
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return defaultMod, ""
	}
	service = info.Main.Path
	var (
		commit string
		dirty  string
	)
	for _, v := range info.Settings {
		switch v.Key {
		case "vcs.revision":
			commit = v.Value[len(v.Value)-8:]
		case "vcs.modified":
			if v.Value == "true" {
				dirty = " (dirty)"
			}
		}
	}

	if commit != "" {
		revision = commit + dirty
	} else {
		revision = "unknown"
	}
	return path.Base(service), revision
}
