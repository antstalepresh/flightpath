package server

import (
	"io"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/steinfletcher/apitest"
)

func TestAPICall(t *testing.T) {
	testCases := []struct {
		name         string
		paths        string
		expectedPath string
		status       int
	}{
		{
			name:         "single_path_valid",
			paths:        `[["SFO", "EWR"]]`,
			expectedPath: `["SFO", "EWR"]`,
			status:       http.StatusOK,
		},
		{
			name:         "multiple_path_valid_1",
			paths:        `[["ATL", "EWR"], ["SFO", "ATL"]]`,
			expectedPath: `["SFO", "EWR"]`,
			status:       http.StatusOK,
		},
		{
			name:         "multiple_path_valid_2",
			paths:        `[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]`,
			expectedPath: `["SFO", "EWR"]`,
			status:       http.StatusOK,
		},
		{
			name:         "invalid_path_1",
			paths:        `[["IND", "IND"]]`,
			expectedPath: `failed to calculate: same path`,
			status:       http.StatusInternalServerError,
		},
		{
			name:         "invalid_path_2",
			paths:        `[["IND", "EWR"], ["EWR", "IND"]]`,
			expectedPath: `failed to calculate: invalid path`,
			status:       http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			apitest.New().
				HandlerFunc(FiberToHandlerFunc(New(fiber.New()).fiber)).
				Post("/calculate").
				Body(tc.paths).
				Expect(t).
				Body(tc.expectedPath).
				Status(tc.status).
				End()
		})
	}
}

func FiberToHandlerFunc(app *fiber.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := app.Test(r)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		// copy headers
		for k, vv := range resp.Header {
			for _, v := range vv {
				w.Header().Add(k, v)
			}
		}
		w.WriteHeader(resp.StatusCode)

		// copy body
		if _, err := io.Copy(w, resp.Body); err != nil {
			panic(err)
		}
	}
}
