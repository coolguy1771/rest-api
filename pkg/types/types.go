package types

import (
	"net/http"
	"time"

	"github.com/go-chi/render"
)

// User is one instance of an user
type User struct {
	// The unique id of this item
	ID int `gorm:"type:SERIAL;PRIMARY_KEY" json:"id" example:"1"`
	// The name of this item
	Name string `gorm:"type:varchar;NOT NULL" json:"name" example:"Frosty Sigh"`
	// The price of this item
	DiscordID string `gorm:"type:varchar;NOT NULL" json:"discord" example:"12345678"`
} // @name User

// Render implements the github.com/go-chi/render.Renderer interface
func (a *User) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Bind implements the the github.com/go-chi/render.Binder interface
func (a *User) Bind(r *http.Request) error {
	return nil
}

// UserList contains a list of users
type UserList struct {
	// A list of users
	Items []*User `json:"users"`
	// The id to query the next page
	NextPageID int `json:"next_page_id,omitempty" example:"10"`
} // @name UserList

// Render implements the github.com/go-chi/render.Renderer interface
func (a *UserList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Unit is one instance of a unit
type Unit struct {
	// The unique id of this unit
	ID int `gorm:"type:SERIAL;primary_key;column:ID" json:"id" example:"1"`
	// DateTime is the date and time of this unit
	Name            string    `gorm:"not null;column:UNIT_NAME" json:"name" example:"Black Element"`
	Description     string    `gorm:"not null;column:UNIT_DESC" json:"desc" example:"A Semi Realism group based out of Reach`
	Owner           string    `gorm:"not null;column:OWNER" json:"owner" example:"Frosty Sigh`
	CreatedDatetime time.Time `gorm:"timestamp;not null;column:CREATED_DATETIME" json:"created"`
	UpdatedDatetime time.Time `gorm:"timestamp;column:UPDATED_DATETIME" json:"lastUpdated,omitempty" example:"0001-01-01 00:00:00+00"`
	UpdateBy        string    `gorm:"column:UPDATED_BY" json:"updatedBy,omitempty" example: "Frosty Sigh`
} // @name Unit

// Render implements the github.com/go-chi/render.Renderer interface
func (o *Unit) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Bind implements the the github.com/go-chi/render.Binder interface
func (o *Unit) Bind(r *http.Request) error {
	return nil
}

// UnitList contains a list of units
type UnitList struct {
	// A list of units
	Units []*Unit `json:"units"`
	// The id to query the next page
	NextPageID int `json:"next_page_id,omitempty" example:"10"`
} // @name OrderList

// Render implements the github.com/go-chi/render.Renderer interface
func (o *UnitList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// ErrResponse renderer type for handling all sorts of errors.
type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status" example:"Resource not found."`                                         // user-level status message
	AppCode    int64  `json:"code,omitempty" example:"404"`                                                 // application-specific error code
	ErrorText  string `json:"error,omitempty" example:"The requested resource was not found on the server"` // application-level error message, for debugging
} // @name ErrorResponse

// Render implements the github.com/go-chi/render.Renderer interface for ErrResponse
func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

// ErrInvalidRequest returns a structured http response for invalid requests
func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusBadRequest,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}

// ErrRender returns a structured http response in case of rendering errors
func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusUnprocessableEntity,
		StatusText:     "Error rendering response.",
		ErrorText:      err.Error(),
	}
}

// ErrNotFound returns a structured http response if a resource couln't be found
func ErrNotFound() render.Renderer {
	return &ErrResponse{
		HTTPStatusCode: http.StatusNotFound,
		StatusText:     "Resource not found.",
	}
}
