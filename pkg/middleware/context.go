package middleware

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/coolguy1771/rest-api/pkg/db"
	"github.com/coolguy1771/rest-api/pkg/types"
)

type (
	// CustomKey is used to refer to the context key that stores custom values of this api to avoid overwrites
	CustomKey string
)

const (
	// UserCtxKey refers to the context key that stores the user
	UserCtxKey CustomKey = "user"
	// UnitCtxKey refers to the context key that stores the unit
	UnitCtxKey CustomKey = "unit"
)

var DBClient db.ClientInterface

func SetDBClient(c db.ClientInterface) {
	DBClient = c
}

// User middleware is used to load an User object from
// the URL parameters passed through as the request. In case
// the User could not be found, we stop here and return a 404.
func User(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user *types.User

		if id := chi.URLParam(r, "id"); id != "" {
			intID, err := strconv.Atoi(id)
			if err != nil {
				_ = render.Render(w, r, types.ErrInvalidRequest(err))
				return
			}
			user = DBClient.GetUserByID(intID)
		} else {
			_ = render.Render(w, r, types.ErrNotFound())
			return
		}
		if user == nil {
			_ = render.Render(w, r, types.ErrNotFound())
			return
		}

		ctx := context.WithValue(r.Context(), UserCtxKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Unit middleware is used to load an Unit object from
// the URL parameters passed through as the request. In case
// the Unit could not be found, we stop here and return a 404.
func Unit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var unit *types.Unit

		if id := chi.URLParam(r, "id"); id != "" {
			intID, err := strconv.Atoi(id)
			if err != nil {
				_ = render.Render(w, r, types.ErrInvalidRequest(err))
				return
			}
			unit = DBClient.GetUnitByID(intID)
		} else {
			_ = render.Render(w, r, types.ErrNotFound())
			return
		}
		if unit == nil {
			_ = render.Render(w, r, types.ErrNotFound())
			return
		}

		ctx := context.WithValue(r.Context(), UnitCtxKey, unit)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
