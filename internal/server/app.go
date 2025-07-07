// ~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~
//      /\_/\
//     ( o.o )
//      > ^ <
//
// Author: Johan Hanekom
// Date: July 2025
//
// ~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~

package server

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

type App struct {
	router http.Handler
	config *Config
	logger *slog.Logger
}

// Creates a new App instance. The App will:
//   - load in the config
//   - initialize the router, handlers and middleware
//   - create a structured logger (slog) that logs to a rotating log file and is shared
//     between the handlers
func NewApp() *App {
	// ====> CREATE A NEW INSTANCE
	app := &App{
		config: MustLoadConfig(),
	}

	// ====> INIT THE LOGGER
	app.initLogger()

	// ====> INIT THE ROUTES
	app.loadRoutes()
	return app
}

// Initializes a structured logger. We specifically create a structured logger since
// you can pass additional filter information to the logs. You can then use `jq` to tail
// for specific attributes. Structured logs are just single json-dumps strings.
// Since we're logging so verbosely, we want to add file rotation and retention
// to the files.
// example log: logger.Info("Starting the server", slog.String("port", a.config.Port))
func (a *App) initLogger() {
	// ====> DEFINE LOG ROTATION
	logFile := &lumberjack.Logger{
		Filename:   a.config.LogPath,
		MaxSize:    10, // megabytes
		MaxBackups: 5,  // keep last 5 files
		MaxAge:     30, // days
		Compress:   true,
	}

	// ====> DEFINE A NEW STOCK STANDARD JSON FILE HANDLER
	handler := slog.NewJSONHandler(
		io.MultiWriter(os.Stdout, logFile),
		&slog.HandlerOptions{
			Level: slog.LevelInfo,
		},
	)

	// ====> ADD SOME BASE ATTRIBUTES THAT WILL ALWAYS BE LOGGED
	a.logger = slog.New(handler).With(
		slog.String("app", a.config.AppName),
		slog.String("version", a.config.Version),
		slog.String("env", a.config.Env),
	)
}

// Starts the server. It does this by listening on the port specified in the config.
// It also handles graceful shutdowns by listening for OS signals. When a signal is received,
// it attempts to gracefully shut down the server within a timeout period.
// If the server does not shut down within the timeout, it will log an error and exit.
// Signals it listens to: os.Interrupt and syscall.SIGTERM.
func (a *App) Start() {
	a.logger.Info("Starting the server", slog.String("port", a.config.Port))

	// ====> CREATE A CHANNEL TO LISTEN FOR OS SIGNALS
	killSig := make(chan os.Signal, 1)
	signal.Notify(killSig, os.Interrupt, syscall.SIGTERM)

	// ====> CREATE A NEW HTTP SERVER
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", a.config.Port),
		Handler: a.router,
	}

	// ====> DEFINE A GO-ROUTINE TO START THE SERVER
	// This will run in the background and listen for incoming requests.
	go func() {
		err := srv.ListenAndServe()

		// If the server is closed gracefully, we log that and exit normally.
		// If there is an error starting the server, we log the error and exit with
		if errors.Is(err, http.ErrServerClosed) {
			a.logger.Info("Server closed gracefully")
		} else if err != nil {
			a.logger.Error("Error starting server", slog.String("error", err.Error()))
			os.Exit(1)
		}
	}()

	a.logger.Info("Server is running", slog.String("address", srv.Addr))

	// ====> WAIT FOR A SHUTDOWN SIGNAL
	<-killSig
	a.logger.Info("Received shutdown signal, shutting down server...")

	// ====> CREATE A CONTEXT WITH TIMEOUT FOR THE SHUTDOWN
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// ====> ATTEMPT TO SHUTDOWN THE SERVER GRACEFULLY
	if err := srv.Shutdown(ctx); err != nil {
		a.logger.Error("Error shutting down server", slog.String("error", err.Error()))
		os.Exit(1)
	}

	// ====> LOG THE SUCCESSFUL SHUTDOWN
	a.logger.Info("Server shutdown complete")
	a.logger.Info("Goodbye! 👋")
}
