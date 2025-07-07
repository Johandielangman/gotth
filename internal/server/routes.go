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

// A function that defines a new chi router.
// We then load a few middleware functions to the router, like CSP, loggers, tracers, etc
// and then register a few endpoints for endpoint handling
func (a *App) loadRoutes() {
	// ====> DEFINE THE CHI ROUTER
	router := chi.NewRouter()

	// =============== // MIDDLEWARE // ===============
	// We add a nice heartbeat
	router.Use(middleware.Heartbeat("/ping"))

	// Since we're dealing with a web app and not a json api, we set the content type
	router.Use(TextHTMLMiddleware)

	// Since we're using templ, we don't want to expose ourselves to possible injection attacks.
	// We therefore have the right content security policies in place.
	// see https://templ.guide/security/injection-attacks
	router.Use(CSPMiddleware)

	// Se register our structured logger here. We want to log to the same file
	router.Use(httplog.RequestLogger(a.logger, &httplog.Options{
		Level: slog.LevelInfo,
	}))

	// Not sure how useful this is yet, but I wanted to make a trace
	// middleware to track how requests chain
	router.Use(TraceMiddleware)

	// =============== // STATIC FILE SERVER // ===============
	a.registerFileServer(router)

	// =============== // ROUTE HANDLERS // ===============
	router.Route("/", a.registerHomeRouter)
	router.Route("/about", a.registerAboutRouter)

	// ====> ASSIGN ROUTER TO THE APP
	a.router = router
}

// We embed the static files within the single binary.
// I got the idea from this guy:
// https://github.com/ajaen4/goth-complete-setup/blob/main/internal/embed.go
// It makes sense. According to templ, you can also host a file
// server that points files ina folder relative to the binary:
// https://templ.guide/developer-tools/live-reload-with-other-tools#serving-static-assets
// But I wanted to compile everything to one file instead
func (a *App) registerFileServer(router chi.Router) {
	staticFS, err := fs.Sub(internal.StaticFiles, "static")
	if err != nil {
		panic(err)
	}

	staticHandler := http.FileServer(http.FS(staticFS))
	router.Handle("/static/*", http.StripPrefix("/static/", staticHandler))
}

// Handler for the home page
func (a *App) registerHomeRouter(router chi.Router) {
	homeHandler := home.NewHomeHandler(a.logger)
	router.Get("/", homeHandler.ServeGetHTTP)
	router.Post("/count/{count}", homeHandler.Count)
}

// Handler for the about page
func (a *App) registerAboutRouter(router chi.Router) {
	aboutHandler := about.NewAboutHandler(a.logger)
	router.Get("/", aboutHandler.ServeGetHTTP)
}
