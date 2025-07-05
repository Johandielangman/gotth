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

func NewApp() *App {
	app := &App{
		config: MustLoadConfig(),
	}
	app.initLogger()
	app.loadRoutes()
	return app
}

func (a *App) initLogger() {
	logFile := &lumberjack.Logger{
		Filename:   "./logs/app.log",
		MaxSize:    10, // megabytes
		MaxBackups: 5,  // keep last 5 files
		MaxAge:     30, // days
		Compress:   true,
	}

	handler := slog.NewJSONHandler(
		io.MultiWriter(os.Stdout, logFile),
		&slog.HandlerOptions{
			Level: slog.LevelInfo,
		},
	)

	a.logger = slog.New(handler).With(
		slog.String("app", "gotth"),
		slog.String("version", "v1.0.0"),
		slog.String("env", a.config.Env),
	)
}

func (a *App) Start() {
	a.logger.Info("Starting the server", slog.String("port", a.config.Port))
	killSig := make(chan os.Signal, 1)

	signal.Notify(killSig, os.Interrupt, syscall.SIGTERM)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", a.config.Port),
		Handler: a.router,
	}

	go func() {
		err := srv.ListenAndServe()

		if errors.Is(err, http.ErrServerClosed) {
			fmt.Println("Server shutdown complete")
		} else if err != nil {
			fmt.Println("Error starting server:", err)
			os.Exit(1)
		}
	}()

	fmt.Println("Server started")
	<-killSig

	fmt.Println("Shutting down server")

	// Create a context with a timeout for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Attempt to gracefully shut down the server
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("Error during server shutdown:", err)
		os.Exit(1)
	}

	fmt.Println("Server shutdown complete")
	fmt.Println("Goodbye! 👋")
}
