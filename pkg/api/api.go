package api

import (
	"net/http"

	"github.com/coolguy1771/rest-api/pkg/db"
	"go.uber.org/zap"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger"

	m "github.com/coolguy1771/rest-api/pkg/middleware"
)

var DBClient db.ClientInterface

func SetDBClient(c db.ClientInterface) {
	DBClient = c
	m.SetDBClient(DBClient)
}

// GetRouter configures a chi router and starts the http server
// @title REST API Written in go
// @description This API is designed for Black Element's PERSCOM Project
// @contact.name Tyler Witlin
// @contact.email 
// @host localhost:8080
// @BasePath /
func GetRouter(log *zap.Logger, dbClient db.ClientInterface) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	SetDBClient(dbClient)
	if log != nil {
		r.Use(m.SetLogger(log))
	}
	buildTree(r)

	return r
}

func buildTree(r *chi.Mux) {
	r.HandleFunc("/swagger", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, r.RequestURI+"/", http.StatusMovedPermanently)
	})
	r.Get("/swagger*", httpSwagger.Handler())

	r.Route("/users", func(r chi.Router) {
		r.With(m.Pagination).Get("/", ListUsers)

		r.Route("/{id}", func(r chi.Router) {
			r.Use(m.User)
			r.Get("/", GetUser)
		})

		r.Put("/", PutUser)
	})

	r.Route("/units", func(r chi.Router) {
		r.With(m.Pagination).Get("/", ListUnits)

		r.Route("/{id}", func(r chi.Router) {
			r.Use(m.Unit)
			r.Get("/", GetUnit)
		})

		r.Put("/", PutUnit)
	})
}
