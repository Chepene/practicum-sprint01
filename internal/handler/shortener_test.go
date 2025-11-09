package handler_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Chepene/practicum-sprint01/internal/app"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestShortenerHandler(t *testing.T) {
	type want struct {
		code        int
		response    string
		contentType string
	}
	tests := []struct {
		name string
		want want
	}{
		{
			name: "positive test #1",
			want: want{
				code:        201,
				response:    `{"status":"ok"}`,
				contentType: "application/json",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			app, err := app.NewApp()
			require.NoError(t, err)

			server := httptest.NewServer(app.Mux)
			defer server.Close()

			request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("https://www.yandex.ru"))
			w := httptest.NewRecorder()
			app.Handler.CreateHandler(w, request)

			res := w.Result()

			assert.Equal(t, test.want.code, res.StatusCode)

			defer res.Body.Close()

			resBody, err := io.ReadAll(res.Body)
			require.NoError(t, err)
			responseStr := string(resBody)

			assert.True(t, strings.HasPrefix(responseStr, "http://localhost:8080/"))
		})
	}
}

/*
func TestStatusHandler(t *testing.T) {
	type want struct {
		code        int
		response    string
		contentType string
	}
	tests := []struct {
		name string
		want want
	}{
		{
			name: "positive test #1",
			want: want{
				code:        200,
				response:    `{"status":"ok"}`,
				contentType: "application/json",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			app, err := app.NewApp()
			require.NoError(t, err)

			server := httptest.NewServer(app.Mux)
			defer server.Close()

			request := httptest.NewRequest(http.MethodPost, "/", nil)
			w := httptest.NewRecorder()
			app.Handler.CreateHandler(w, request)
			res := w.Result()

			assert.Equal(t, test.want.code, res.StatusCode)

			defer res.Body.Close()

			//resBody, err := io.ReadAll(res.Body)

			//require.NoError(t, err)
		})
	}
}
*/
