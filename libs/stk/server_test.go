package stk_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/adharshmk96/fitsphere-be/libs/stk"
)

// Test server routes

func TestServerRoutes(t *testing.T) {
	s := stk.NewServer()

	sampleHandler := func(ctx *stk.Context) {
		method := ctx.Request.Method
		ctx.Writer.WriteHeader(200)
		ctx.Writer.Write([]byte(method))
	}

	s.Get("/test-get", sampleHandler)
	s.Post("/test-post", sampleHandler)
	s.Put("/test-put", sampleHandler)
	s.Delete("/test-delete", sampleHandler)
	s.Patch("/test-patch", sampleHandler)

	paramsHandler := func(ctx *stk.Context) {
		params := ctx.Params
		ctx.Writer.WriteHeader(200)
		ctx.Writer.Write([]byte(params.ByName("id")))
	}

	s.Get("/test/:id", paramsHandler)
	s.Post("/test/:id", paramsHandler)
	s.Put("/test/:id", paramsHandler)
	s.Delete("/test/:id", paramsHandler)
	s.Patch("/test/:id", paramsHandler)

	serverHandler := http.HandlerFunc(s.Router.ServeHTTP)

	t.Run("testing GET", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/test-get", nil)
		rr := httptest.NewRecorder()
		serverHandler.ServeHTTP(rr, req)

		res := rr.Result()
		body, _ := ioutil.ReadAll(res.Body)
		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, "GET", string(body))
	})

	t.Run("testing POST", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/test-post", nil)
		rr := httptest.NewRecorder()
		serverHandler.ServeHTTP(rr, req)

		res := rr.Result()
		body, _ := ioutil.ReadAll(res.Body)
		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, "POST", string(body))
	})

	t.Run("testing PUT", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/test-put", nil)
		rr := httptest.NewRecorder()
		serverHandler.ServeHTTP(rr, req)

		res := rr.Result()
		body, _ := ioutil.ReadAll(res.Body)
		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, "PUT", string(body))
	})

	t.Run("testing DELETE", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/test-delete", nil)
		rr := httptest.NewRecorder()
		serverHandler.ServeHTTP(rr, req)

		res := rr.Result()
		body, _ := ioutil.ReadAll(res.Body)
		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, "DELETE", string(body))
	})

	t.Run("testing PATCH", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPatch, "/test-patch", nil)
		rr := httptest.NewRecorder()
		serverHandler.ServeHTTP(rr, req)

		res := rr.Result()
		body, _ := ioutil.ReadAll(res.Body)
		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, "PATCH", string(body))
	})

	t.Run("testing GET with params returning the param", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/test/123", nil)
		rr := httptest.NewRecorder()
		serverHandler.ServeHTTP(rr, req)

		res := rr.Result()
		body, _ := ioutil.ReadAll(res.Body)
		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, "123", string(body))
	})

	t.Run("testing POST with params returning the param", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/test/123", nil)
		rr := httptest.NewRecorder()
		serverHandler.ServeHTTP(rr, req)

		res := rr.Result()
		body, _ := ioutil.ReadAll(res.Body)
		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, "123", string(body))
	})

	t.Run("testing PUT with params returning the param", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/test/123", nil)
		rr := httptest.NewRecorder()
		serverHandler.ServeHTTP(rr, req)

		res := rr.Result()
		body, _ := ioutil.ReadAll(res.Body)
		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, "123", string(body))
	})

	t.Run("testing DELETE with params returning the param", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/test/123", nil)
		rr := httptest.NewRecorder()
		serverHandler.ServeHTTP(rr, req)

		res := rr.Result()
		body, _ := ioutil.ReadAll(res.Body)
		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, "123", string(body))
	})

	t.Run("testing PATCH with params returning the param", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPatch, "/test/123", nil)
		rr := httptest.NewRecorder()
		serverHandler.ServeHTTP(rr, req)

		res := rr.Result()
		body, _ := ioutil.ReadAll(res.Body)
		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, "123", string(body))
	})

	t.Run("testing GET route with POST method should return method not allowed 405", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/test-get", nil)
		rr := httptest.NewRecorder()
		serverHandler.ServeHTTP(rr, req)

		res := rr.Result()
		body, _ := ioutil.ReadAll(res.Body)
		assert.Equal(t, http.StatusMethodNotAllowed, res.StatusCode)
		assert.Equal(t, "Method Not Allowed\n", string(body))
	})

	t.Run("testing POST route with GET method should return method not allowed 405", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/test-post", nil)
		rr := httptest.NewRecorder()
		serverHandler.ServeHTTP(rr, req)

		res := rr.Result()
		body, _ := ioutil.ReadAll(res.Body)
		assert.Equal(t, http.StatusMethodNotAllowed, res.StatusCode)
		assert.Equal(t, "Method Not Allowed\n", string(body))
	})

	t.Run("testing PUT route with POST method should return method not allowed 405", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/test-put", nil)
		rr := httptest.NewRecorder()
		serverHandler.ServeHTTP(rr, req)

		res := rr.Result()
		body, _ := ioutil.ReadAll(res.Body)
		assert.Equal(t, http.StatusMethodNotAllowed, res.StatusCode)
		assert.Equal(t, "Method Not Allowed\n", string(body))
	})

	t.Run("testing DELETE route with GET method should return method not allowed 405", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/test-delete", nil)
		rr := httptest.NewRecorder()
		serverHandler.ServeHTTP(rr, req)

		res := rr.Result()
		body, _ := ioutil.ReadAll(res.Body)
		assert.Equal(t, http.StatusMethodNotAllowed, res.StatusCode)
		assert.Equal(t, "Method Not Allowed\n", string(body))
	})

	t.Run("testing PATCH route with PUT method should return method not allowed 405", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/test-patch", nil)
		rr := httptest.NewRecorder()
		serverHandler.ServeHTTP(rr, req)

		res := rr.Result()
		body, _ := ioutil.ReadAll(res.Body)
		assert.Equal(t, http.StatusMethodNotAllowed, res.StatusCode)
		assert.Equal(t, "Method Not Allowed\n", string(body))
	})

}

// Test middlewares

func TestMiddlewares(t *testing.T) {

	firstMiddleware := func(next stk.HandlerFunc) stk.HandlerFunc {
		return func(ctx *stk.Context) {
			ctx.Writer.Header().Add("X-FirstMiddleware", "true")
			next(ctx)
		}
	}

	secondMiddleware := func(next stk.HandlerFunc) stk.HandlerFunc {
		return func(ctx *stk.Context) {
			ctx.Writer.Header().Add("X-SecondMiddleware", "true")
			next(ctx)
		}
	}

	myHandler := func(ctx *stk.Context) {
		fmt.Fprintln(ctx.Writer, "Hello, world!")
	}

	t.Run("server with two middlewares", func(t *testing.T) {
		s := stk.NewServer()

		s.Use(firstMiddleware)
		s.Use(secondMiddleware)

		s.Get("/", myHandler)

		req := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()

		s.Router.ServeHTTP(w, req)

		resp := w.Result()
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status OK, got %v", resp.Status)
		}

		if resp.Header.Get("X-FirstMiddleware") != "true" {
			t.Error("First middleware not executed")
		}

		if resp.Header.Get("X-SecondMiddleware") != "true" {
			t.Error("Second middleware not executed")
		}
	})

	t.Run("server with no middlewares", func(t *testing.T) {
		s := stk.NewServer()
		s.Get("/", myHandler)

		req := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()

		s.Router.ServeHTTP(w, req)

		resp := w.Result()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status OK, got %v", resp.Status)
		}

		if resp.Header.Get("X-FirstMiddleware") != "" {
			t.Error("First middleware executed when it shouldn't")
		}

		if resp.Header.Get("X-SecondMiddleware") != "" {
			t.Error("Second middleware executed when it shouldn't")
		}
	})

}
