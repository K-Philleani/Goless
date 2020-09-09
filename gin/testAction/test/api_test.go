package test

import (
	"Goless/gin/testAction/setInit"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
)

var router *gin.Engine

func init() {
	gin.SetMode(gin.TestMode)
	router = setInit.SetRouter()
}

func TestGetParamApi(t *testing.T) {
	username := "贾宏伟"
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/user/" + username, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Hello" + username, w.Body.String())
}

func TestGetQueryApi(t *testing.T) {
	username := "贾宏伟"
	age := 24
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/username?username=" + username + "&age=" + strconv.Itoa(age) , nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "hello: 贾宏伟/"+  strconv.Itoa(age), w.Body.String())
}

func TestGetHtmlApi(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/tem" , nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestTemPost(t *testing.T) {
	value := url.Values{}
	value.Add("email", "964202623@qq.com")
	value.Add("password", "1234")
	value.Add("password-again", "1234")
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/tem/register", bytes.NewBufferString(value.Encode()))
	req.Header.Add("Content-type", "application/x-www-form-urlencoded;param=value")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestTemPostEmailAndPwd(t *testing.T) {
	value := url.Values{}
	value.Add("email", "964202623@qq.com")
	value.Add("password", "19951220")
	value.Add("password-again", "19951220")
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/tem/register", bytes.NewBufferString(value.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; param=value")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

