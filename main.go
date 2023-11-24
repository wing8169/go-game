package main

import (
	"context"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/websocket"
	"github.com/wing8169/go-game/views"
)

var (
	upgrader = websocket.Upgrader{}
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		component := views.Index()
		component.Render(ctx, w)
	})
	workDir, _ := os.Getwd()
	FileServer(r, "/css", http.Dir(filepath.Join(workDir, "css")))
	FileServer(r, "/fonts", http.Dir(filepath.Join(workDir, "fonts")))
	FileServer(r, "/static", http.Dir(filepath.Join(workDir, "static")))

	http.ListenAndServe(":3000", r)
}
