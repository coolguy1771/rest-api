package api

import (
	"net/http"

	"github.com/go-chi/render"

	m "github.com/coolguy1771/rest-api/pkg/middleware"
	"github.com/coolguy1771/rest-api/pkg/types"
)

// GetUser renders the user from the context
// @Summary Get user by id
// @Description GetUser returns a single user by id
// @Tags Users
// @Produce json
// @Param id path string true "user id"
// @Router /users/{id} [get]
// @Success 200 {object} types.User
// @Failure 400 {object} types.ErrResponse
// @Failure 404 {object} types.ErrResponse
func GetUser(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(m.UserCtxKey).(*types.User)

	if err := render.Render(w, r, user); err != nil {
		_ = render.Render(w, r, types.ErrRender(err))
		return
	}
}

// PutUser writes an user to the database
// @Summary Add an user to the database
// @Description PutUser writes an user to the database
// @Description To write a new user, leave the id empty. To update an existing one, use the id of the user to be updated
// @Tags Users
// @Produce json
// @Param id path string true "user id"
// @Param name path string true "user name"
// @Param discord path string true "user discord"
// @Router /users [put]
// @Success 200 {object} types.User
// @Failure 400 {object} types.ErrResponse
// @Failure 404 {object} types.ErrResponse
func PutUser(w http.ResponseWriter, r *http.Request) {
	user := &types.User{}
	if err := render.Bind(r, user); err != nil {
		_ = render.Render(w, r, types.ErrInvalidRequest(err))
		return
	}

	if err := DBClient.SetUser(user); err != nil {
		_ = render.Render(w, r, types.ErrInvalidRequest(err))
		return
	}

	if err := render.Render(w, r, user); err != nil {
		_ = render.Render(w, r, types.ErrRender(err))
		return
	}
}

// ListUsers returns all users in the database
// @Summary List all users
// @Description Get all users stored in the database
// @Tags Users
// @Produce json
// @Param page_id query string false "id of the page to be retrieved"
// @Router /users [get]
// @Success 200 {object} types.UserList
// @Failure 400 {object} types.ErrResponse
// @Failure 404 {object} types.ErrResponse
func ListUsers(w http.ResponseWriter, r *http.Request) {
	pageID := r.Context().Value(m.PageIDKey)
	if err := render.Render(w, r, DBClient.GetUsers(pageID.(int))); err != nil {
		_ = render.Render(w, r, types.ErrRender(err))
		return
	}
}

// GetUnit renders the unit from the context
// @Summary Get unit by id
// @Description GetUnit returns a single unit by id
// @Tags Units
// @Produce json
// @Param id path string true "unit id"
// @Router /unit/{id} [get]
// @Success 200 {object} types.Unit
// @Failure 400 {object} types.ErrResponse
// @Failure 404 {object} types.ErrResponse
func GetUnit(w http.ResponseWriter, r *http.Request) {
	unit := r.Context().Value(m.UnitCtxKey).(*types.Unit)

	if err := render.Render(w, r, unit); err != nil {
		_ = render.Render(w, r, types.ErrRender(err))
		return
	}
}

// PutUnit writes an unit to the database
// @Summary Add an unit to the database
// @Description PutUnit writes an unit to the database
// @Description To write a new unit, leave the id empty. To update an existing one, use the id of the unit to be updated
// @Tags Units
// @Produce json
// @Router /units [put]
// @Success 200 {object} types.Unit
// @Failure 400 {object} types.ErrResponse
// @Failure 404 {object} types.ErrResponse
func PutUnit(w http.ResponseWriter, r *http.Request) {
	unit := &types.Unit{}
	if err := render.Bind(r, unit); err != nil {
		_ = render.Render(w, r, types.ErrInvalidRequest(err))
		return
	}

	if err := DBClient.SetUnit(unit); err != nil {
		_ = render.Render(w, r, types.ErrInvalidRequest(err))
		return
	}

	if err := render.Render(w, r, unit); err != nil {
		_ = render.Render(w, r, types.ErrRender(err))
		return
	}
}

// ListUnits returns all units in the database
// @Summary List all units
// @Description Get all units stored in the database
// @Tags Units
// @Produce json
// @Param page_id query string false "id of the page to be retrieved"
// @Router /units [get]
// @Success 200 {object} types.UnitList
// @Failure 400 {object} types.ErrResponse
// @Failure 404 {object} types.ErrResponse
func ListUnits(w http.ResponseWriter, r *http.Request) {
	pageID := r.Context().Value(m.PageIDKey)
	if err := render.Render(w, r, DBClient.GetUnits(pageID.(int))); err != nil {
		_ = render.Render(w, r, types.ErrRender(err))
		return
	}
}
