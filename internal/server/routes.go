package server

import (
	"gotth/internal"
	"gotth/internal/views/about"
	"gotth/internal/views/home"
	"io/fs"
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

	a.registerFileServer(router)
	router.Route("/", a.registerHomeRouter)
	router.Route("/about", a.registerAboutRouter)

	a.router = router
}

func (a *App) registerFileServer(router chi.Router) {
	staticFS, err := fs.Sub(internal.StaticFiles, "static")
	if err != nil {
		panic(err)
	}

	staticHandler := http.FileServer(http.FS(staticFS))
	router.Handle("/static/*", http.StripPrefix("/static/", staticHandler))
}

func (a *App) registerHomeRouter(router chi.Router) {
	homeHandler := home.NewHomeHandler(a.logger)
	router.Get("/", homeHandler.ServeGetHTTP)
}

func (a *App) registerAboutRouter(router chi.Router) {
	aboutHandler := about.NewAboutHandler(a.logger)
	router.Get("/", aboutHandler.ServeGetHTTP)
}
