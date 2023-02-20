package test

import (
	"bytes"
	"encoding/json"
	"github.com/hiejulia/api-online-book-store/api/auth"
	"github.com/hiejulia/api-online-book-store/api/user"
	"github.com/hiejulia/api-online-book-store/clients"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostPublicRegister(t *testing.T) {
	clients.SetUpMain()
	Convey("Test TestPostPublicRegister succeed ", t, func() {
		r := SetRouter()
		auth.SetupMiddleware(r)
		user.SetupPrivacy()
		r.POST("/api/v1/users/register", user.PostPublicRegister)

		userRegister := user.LoginRequest{
			Email:    "hien@gmail.com",
			Password: "password",
		}
		jsonValue, _ := json.Marshal(userRegister)
		req, err := http.NewRequest("POST", "/api/v1/users/register", bytes.NewBuffer(jsonValue))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		So(err, ShouldBeNil)
	})
}

func TestPostPublicLogin(t *testing.T) {
	Convey("Test TestPostPublicLogin succeed ", t, func() {
		r := SetRouter()
		auth.SetupMiddleware(r)
		user.SetupPrivacy()
		r.POST("/api/v1/users/login", user.PostPublicLogin)

		userLogin := user.LoginRequest{
			Email:    "hien@gmail.com",
			Password: "password",
		}
		jsonValue, _ := json.Marshal(userLogin)
		req, err := http.NewRequest("POST", "/api/v1/users/login", bytes.NewBuffer(jsonValue))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		So(err, ShouldBeNil)
	})
}
