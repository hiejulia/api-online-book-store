package test

import (
	"bytes"
	"encoding/json"
	"github.com/hiejulia/api-online-book-store/api/auth"
	"github.com/hiejulia/api-online-book-store/api/cart"
	"github.com/hiejulia/api-online-book-store/api/user"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddItemToCart(t *testing.T) {
	Convey("Test TestAddItemToCart succeed ", t, func() {
		r := SetRouter()
		auth.SetupMiddleware(r)
		user.SetupPrivacy()

		r.GET("/api/v1/cart/1", cart.AddItemToCart)
		cartItem := cart.CartItemRequest{
			CartID: "1",
			BookID: "1",
			Qty:    10,
		}
		jsonValue, _ := json.Marshal(cartItem)
		req, err := http.NewRequest("POST", "/api/v1/orders/user/1", bytes.NewBuffer(jsonValue))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		So(err, ShouldBeNil)
	})
}

func TestGetItemsByCartId(t *testing.T) {
	Convey("Test TestGetItemsByCartId succeed ", t, func() {
		r := SetRouter()
		auth.SetupMiddleware(r)
		user.SetupPrivacy()
		r.GET("/api/v1/cart/1", cart.GetItemsByCartId)

		req, err := http.NewRequest("GET", "/api/v1/cart/1", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		So(err, ShouldBeNil)
	})
}

//
//func TestAddQuantityFromCart(t *testing.T) {
//	Convey("Test TestAddQuantityFromCart succeed ", t, func() {
//		r := SetRouter()
//		auth.SetupMiddleware(r)
//		user.SetupPrivacy()
//		r.POST("/api/v1/cart/1/book/1", cart.AddQuantityFromCart)
//
//		req, err := http.NewRequest("POST", "/api/v1/cart/1", nil)
//		w := httptest.NewRecorder()
//		r.ServeHTTP(w, req)
//		assert.Equal(t, http.StatusOK, w.Code)
//		So(err, ShouldBeNil)
//	})
//}
//
//func TestRemoveQuantityFromCart(t *testing.T) {
//	Convey("Test TestRemoveQuantityFromCart succeed ", t, func() {
//		r := SetRouter()
//		auth.SetupMiddleware(r)
//		user.SetupPrivacy()
//		r.POST("/api/v1/cart/1/book/1", cart.RemoveQuantityFromCart)
//
//		req, err := http.NewRequest("POST", "/api/v1/cart/1/book/1", nil)
//		w := httptest.NewRecorder()
//		r.ServeHTTP(w, req)
//		assert.Equal(t, http.StatusOK, w.Code)
//		So(err, ShouldBeNil)
//	})
//}
