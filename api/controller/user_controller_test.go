package controller

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/Jeswyrne/chlnge/pkg/user"
	"github.com/patrickmn/go-cache"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {

	tests := []struct {
		name       string
		urlQuery   string
		method     string
		count      int
		statusCode int
	}{

		{"Test handler with no query parameters", "", "GET", 0, 404},
		{"Test handler with get method", "rick,marty", "GET", 2, 200},
		{"Test handler with options method", "", "OPTIONS", 0, 204},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			usr := NewUser(cache.New(cache.DefaultExpiration, cache.DefaultExpiration))

			var data user.InfoList
			request := httptest.NewRequest(test.method, "/users/info", nil)
			response := httptest.NewRecorder()

			qParam := request.URL.Query()
			qParam.Add("users", test.urlQuery)
			request.URL.RawQuery = qParam.Encode()

			usr.Handler(response, request)

			statusCode := response.Result().StatusCode
			assert.Equal(t, test.statusCode, statusCode)

			res := response.Body.String()

			json.Unmarshal([]byte(res), &data)
			assert.Equal(t, test.count, len(data))
		})
	}
}
