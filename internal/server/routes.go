package server

import (
	"gotth/internal/views/home"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func (a *App) LoadRoutes() {
	router := chi.NewRouter()

	router.Use(middleware.Heartbeat("/ping"))
	router.Use(middleware.Logger)

	fileServer := http.FileServer(http.Dir("../../static"))
	router.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	router.Route("/", a.registerHmeRouter)

	a.router = router
}

func (a *App) registerHmeRouter(router chi.Router) {
	homeHandler := home.NewHomeHandler()
	router.Get("/", homeHandler.ServeHTTP)
}
