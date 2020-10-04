package test

import (
	"GoLi/gin/action/router/setRouter"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserSave(t *testing.T) {
	name := "贾宏伟"
	router := setRouter.SetRouter()
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/user/" + name, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "name: " + name, w.Body.String())
}
