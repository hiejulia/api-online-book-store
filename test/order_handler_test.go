package test

import (
	"github.com/hiejulia/api-online-book-store/api/auth"
	"github.com/hiejulia/api-online-book-store/api/order"
	"github.com/hiejulia/api-online-book-store/api/user"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllOrdersByUser(t *testing.T) {
	Convey("Test GetAllOrdersByUser succeed ", t, func() {
		r := SetRouter()
		auth.SetupMiddleware(r)
		user.SetupPrivacy()
		r.GET("/api/v1/orders/user/1", order.GetAllOrdersByUser)

		req, err := http.NewRequest("GET", "/api/v1/orders/user/1", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		So(err, ShouldBeNil)
	})
}

func TestCreateOrder(t *testing.T) {
	Convey("Test CreateOrder succeed ", t, func() {
		r := SetRouter()
		auth.SetupMiddleware(r)
		user.SetupPrivacy()
		r.POST("/api/v1/cart/1", order.CreateOrder)

		req, err := http.NewRequest("POST", "/api/v1/cart/1", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		So(err, ShouldBeNil)
	})
}
