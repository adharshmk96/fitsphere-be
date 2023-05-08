package stk_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/adharshmk96/fitsphere-be/libs/stk"
	"github.com/julienschmidt/httprouter"
)

type TestPayload struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func TestJSONResponse(t *testing.T) {
	testCases := []struct {
		name        string
		data        interface{}
		expectedErr error
	}{
		{
			name:        "Invalid JSON",
			data:        make(chan int),
			expectedErr: stk.ErrInternalServer,
		},
		{
			name: "Struct data",
			data: TestPayload{
				Message: "Hello, this is a JSON response!",
				Status:  http.StatusOK,
			},
			expectedErr: nil,
		},
		{
			name: "Map data",
			data: map[string]interface{}{
				"message": "Hello, this is a JSON response!",
				"status":  http.StatusOK,
			},
			expectedErr: nil,
		},
		{
			name:        "nil",
			data:        nil,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			request, _ := http.NewRequest("GET", "/", nil)
			responseRec := httptest.NewRecorder()

			router := httprouter.New()
			router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
				context := &stk.Context{
					Request: r,
					Writer:  w,
				}
				context.JSONResponse(tc.data)
			})

			router.ServeHTTP(responseRec, request)

			if tc.expectedErr != nil {
				expectedErr := tc.expectedErr.Error()
				if responseRec.Body.String() != string(expectedErr) {
					t.Errorf("Expected error to be %q but got %q", expectedErr, responseRec.Body.String())
				}
			} else {
				expectedJSON, _ := json.Marshal(tc.data)
				if responseRec.Body.String() != string(expectedJSON) {
					t.Errorf("Expected JSON data to be %q but got %q", expectedJSON, responseRec.Body.String())
				}
			}
		})
	}
}
