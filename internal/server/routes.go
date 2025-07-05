package server

import (
	"gotth/internal/views/about"
	"gotth/internal/views/home"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog/v3"
)

func (a *App) loadRoutes() {
	router := chi.NewRouter()

	router.Use(middleware.Heartbeat("/ping"))
	router.Use(httplog.RequestLogger(a.logger, &httplog.Options{
		Level: slog.LevelInfo,
	}))
	router.Use(TraceMiddleware)

	fileServer := http.FileServer(http.Dir("./internal/static"))
	router.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	router.Route("/", a.registerHomeRouter)
	router.Route("/about", a.registerAboutRouter)

	a.router = router
}

func (a *App) registerHomeRouter(router chi.Router) {
	homeHandler := home.NewHomeHandler(a.logger)
	router.Get("/", homeHandler.ServeHTTP)
}

func (a *App) registerAboutRouter(router chi.Router) {
	aboutHandler := about.NewAboutHandler(a.logger)
	router.Get("/", aboutHandler.ServeHTTP)
}
