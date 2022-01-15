package main

import (
	"io"
	"net/http"
	"strings"
	"text/template"

	"github.com/coolguy1771/rest-api/models"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

var (
	seq = 1
)

//----------
// Handlers
//----------

func createUser(c echo.Context) error {
	u := &models.ApiUser{}
	if err := c.Bind(u); err != nil {
		return err
	}
	seq++
	return c.JSON(http.StatusCreated, u)
}

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.ERROR)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.RequestID())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339}, method=${method}, uri=${uri}, status=${status}, requestID=${id}, host=${host}\n",
	}))
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
		LogLevel:  log.ERROR,
	}))
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Path(), "metrics") // Change "metrics" for your own path
		},
	}))
	//e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
	//	SigningKey: []byte(handler.Key),
	//	Skipper: func(c echo.Context) bool {
	//		// Skip authentication for signup and login requests
	//		if c.Path() == "/login" || c.Path() == "/signup" {
	//			return true
	//		}
	//		return false
	//	},
	//}))

	//e.Static("/", "static")
	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)
	e.POST("/api/v1/createUser", createUser)
	//h := &handler.Handler{DB: db}
	// Routes
	//e.POST("/api/v1/createuser", )
	//e.POST("/login", h.Login)
	//e.POST("/follow/:id", h.Follow)
	//e.POST("/posts", h.CreatePost)
	//e.GET("/feed", h.FetchPost)
	t := &Template{
		templates: template.Must(template.ParseGlob("static/templates/*.html")),
	}

	e.Renderer = t

	e.GET("/", Index)
	//	e.GET("/news", News)

	e.Logger.Fatal(e.Start(":1323"))
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Index(c echo.Context) error {
	return c.Render(http.StatusOK, "base", "index")
}

//func News(c echo.Context) error {
//	return c.Render(http.StatusOK, "base", "news")
//}
