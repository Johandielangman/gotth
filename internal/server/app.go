package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

type App struct {
	router http.Handler
	config *Config
}

func NewApp() *App {
	app := &App{
		config: MustLoadConfig(),
	}
	app.LoadRoutes()
	return app
}

func (a *App) Start() {
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
