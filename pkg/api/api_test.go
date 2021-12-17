package api

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/coolguy1771/rest-api/pkg/types"

	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"github.com/coolguy1771/rest-api/pkg/api/mocks"
)

var (
	testUser1 = types.User{
		ID:    1,
		Name:  "Frosty Sigh",
		DiscordID: "12345678",
	}
	testUser2 = types.User{
		ID:    2,
		Name:  "Runic Moose",
		DiscordID: "87654321",
	}
)

// TestGetRouter ensures that the router contains all expected routes
func TestGetRouter(t *testing.T) {
	log, _ := zap.NewProduction(zap.WithCaller(false))
	r := GetRouter(log, nil)

	testcases := map[string]struct {
		method string
		path   string
	}{
		"GET /users": {
			method: http.MethodGet,
			path:   "/users",
		},
		"GET /users/{id}": {
			method: http.MethodGet,
			path:   "/users/id",
		},
		"GET /units": {
			method: http.MethodGet,
			path:   "/units",
		},
		"GET /units/{id}": {
			method: http.MethodGet,
			path:   "/units/id",
		},
		"GET /swagger": {
			method: http.MethodGet,
			path:   "/swagger",
		},
	}

	for name, test := range testcases {
		t.Run(name, func(t *testing.T) {
			got := r.Match(chi.NewRouteContext(), test.method, test.path)
			assert.Equal(t, true, got, fmt.Sprintf("not found: %s '%s'", test.method, test.path))
		})
	}
}

//go:generate $GOPATH/bin/mockgen -destination=./mocks/db.go -package=mocks github.com/coolguy1771/rest-api/pkg/db ClientInterface

func getDBClientMock(t *testing.T) *mocks.MockClientInterface {
	ctrl := gomock.NewController(t)
	dbClient := mocks.NewMockClientInterface(ctrl)

	dbClient.EXPECT().GetUsers(gomock.Eq(0)).Return(&types.UserList{
		Items: []*types.User{
			&testUser1,
			&testUser2,
		},
	})

	dbClient.EXPECT().GetUsers(gomock.Eq(1)).Return(&types.UserList{
		Items: []*types.User{
			&testUser2,
		},
	})

	dbClient.EXPECT().GetUserByID(gomock.Eq(1)).Return(&testUser1).AnyTimes()

	dbClient.EXPECT().SetUser(gomock.Any()).DoAndReturn(func(user *types.User) error {
		if user.ID == 0 {
			user.ID = 1
		}
		return nil
	}).AnyTimes()

	return dbClient
}

// TestEndpoints ensures the expected results upon requests
func TestEndpoints(t *testing.T) {
	r := GetRouter(nil, getDBClientMock(t))
	ts := httptest.NewServer(r)
	defer ts.Close()

	testcases := map[string]struct {
		method   string
		path     string
		body     string
		header   http.Header
		wantCode int
		wantBody string
	}{
		"GET /users": {
			method:   http.MethodGet,
			path:     "/users",
			wantCode: http.StatusOK,
			wantBody: `{"users":[{"id":1,"name":"Frosty Sigh","discord":"12345678"},{"id":2,"name":"Runic Moose","discord":"87654321"}]}`,
		},
		"GET /users?page_id=1": {
			method:   http.MethodGet,
			path:     "/users?page_id=1",
			wantCode: http.StatusOK,
			wantBody: `{"users":[{"id":2,"name":"Runic Moose","discord":"87654321"}]}`,
		},
		"GET /users/{id}": {
			method:   http.MethodGet,
			path:     "/users/1",
			wantCode: http.StatusOK,
			wantBody: `{"id":1,"name":"Frosty Sigh","discord":"12345678"}`,
		},
		"PUT /users": {
			method: http.MethodPut,
			path:   "/users",
			body:   `{"name":"Frosty Sigh","discord":"12345678"}`,
			header: map[string][]string{
				"Content-Type": {"application/json"},
			},
			wantCode: http.StatusOK,
			wantBody: `{"id":1,"name":"Frosty Sigh","discord":"12345678"}`,
		},
		"Page Not Found": {
			method:   http.MethodGet,
			path:     "/blah",
			wantCode: http.StatusNotFound,
			wantBody: "404 page not found",
		},
	}

	for name, test := range testcases {
		t.Run(name, func(t *testing.T) {
			body := bytes.NewReader([]byte(test.body))
			gotResponse, gotBody := testRequest(t, ts, test.method, test.path, body, test.header)
			assert.Equal(t, test.wantCode, gotResponse.StatusCode)
			if test.wantBody != "" {
				assert.Equal(t, test.wantBody, gotBody, "body did not match")
			}
		})
	}
}

// testRequest is a helper function to exectute the http request against the server
func testRequest(t *testing.T, ts *httptest.Server, method, path string, body io.Reader, header http.Header) (*http.Response, string) {
	t.Helper()
	req, err := http.NewRequest(method, ts.URL+path, body)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}
	req.Header = header

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	respBody = bytes.TrimSpace(respBody)

	return resp, string(respBody)
}
