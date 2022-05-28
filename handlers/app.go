package handlers

import (
	"net/http"

	"github.com/batt0s/Ninja-Manga-Api/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

type App struct {
	Addr   string
	Router *chi.Mux
}

func (app *App) Init(mode string) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	database.InitDB(mode)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST"},
	}))
	r.Route("/api", func(api chi.Router) {
		api.Get("/manga", GetAllMangaHandler)
		api.Get("/manga/{id}", GetMangaByIdHandler)
		api.Get("/manga/tag/{tag}", GetAllMangaByTagHandler)
		api.Get("/manga/search/{keyword}", SearchMangaHandler)
		api.Post("/manga", CreateMangaHandler)
		api.Get("/chapter/{id}", GetChapterByIdHandler)
		api.Get("/manga/chapters/{mangaid}", GetChaptersByMangaHandler)
		api.Post("/chapter", CreateChapterHandler)
		api.Delete("/manga/{mangaid}", DeleteMangaHandler)
		api.Delete("/chapter/{chapterid}", DeleteChapterHandler)
	})
	app.Router = r
}

func (app *App) Run() {
	http.ListenAndServe(app.Addr, app.Router)
}
