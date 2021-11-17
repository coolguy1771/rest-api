package auth

import (
	"context"
	"net/http"
	"testing"

	"github.com/coolguy1771/rest-api/internal/test"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

func TestCurrentUser(t *testing.T) {
	ctx := context.Background()
	assert.Nil(t, CurrentUser(ctx))
	ctx = WithUser(ctx, "100", "test")
	identity := CurrentUser(ctx)
	if assert.NotNil(t, identity) {
		assert.Equal(t, "100", identity.getId())
		assert.Equal(t, "test", identity.getName())
	}
}

func TestHandler(t *testing.T) {
	assert.NotNil(t, Handler("test"))
}

func Test_HandleToken(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://example.com", nil)
	ctx, _ := test.MockRoutingContext(req)
	assert.Nil(t, CurrentUser(ctx.Request.Context()))

	err := HandleToken(ctx, &jwt.Token{
		Claims: jwt.MapClaims{
			"id":   "100",
			"name": "test",
		},
	})
	assert.Nil(t, err)
	identity := CurrentUser(ctx.Request.Context())
	if assert.NotNil(t, identity) {
		assert.Equal(t, "100", identity.getId())
		assert.Equal(t, "test", identity.getName())
	}
}

func TestMocks(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://example.com", nil)
	ctx, _ := test.MockRoutingContext(req)
	assert.NotNil(t, MockAuthHandler(ctx))
	req.Header = MockAuthHeader()
	ctx, _ = test.MockRoutingContext(req)
	assert.Nil(t, MockAuthHandler(ctx))
	assert.NotNil(t, CurrentUser(ctx.Request.Context()))
}
