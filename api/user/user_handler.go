package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/hiejulia/api-online-book-store/api/common"
	"github.com/hiejulia/api-online-book-store/clients"
	"github.com/hiejulia/api-online-book-store/models"
	"github.com/hiejulia/api-online-book-store/utils"
	"net/http"
	"strings"
)

var (
	ErrDuplicateEmail = errors.New("Requested email already exists.")
	ErrLogin          = errors.New("Incorrect username or password. Please try again.")
	ErrNotFound       = errors.New("Record not found.")
)

// PostPublicRegister godoc
// @Summary Register a user
// @Description Register a user
// @Tags users
// @Accept json
// @Produce json
// @Success 200
// @Failure 400
// @Router /users [post]
func PostPublicRegister(c *gin.Context) {
	req := new(RegisterRequest)
	if err := c.BindJSON(req); err != nil {
		common.Error(c, http.StatusBadRequest, err)
		return
	}

	var err error
	db := c.MustGet("db").(*clients.SQL)

	//check for duplicate email and user name
	userCheck := &models.User{}
	users := make([]models.User, 0)
	where := map[string]interface{}{
		"email = ?": req.Email,
	}
	if err := db.FindWhere(userCheck, where, &users); err != nil {
		common.Error(c, http.StatusInternalServerError, err)
		return
	}
	for _, v := range users {
		deleted := false

		if !deleted && strings.EqualFold(v.Email, req.Email) {
			common.Error(c, http.StatusBadRequest, ErrDuplicateEmail)
			return
		}
	}

	if err != nil {
		common.Error(c, http.StatusBadRequest, err)
		return
	}

	user := models.User{
		ID:        utils.ID(),
		Email:     req.Email,
		Password:  req.Password,
		CreatedAt: utils.Timestamp(),
	}
	if err = user.HashPassword(); err != nil {
		common.Error(c, http.StatusBadRequest, err)
		return
	}

	if err = db.Create(&user); err != nil {
		common.Error(c, http.StatusInternalServerError, err)
		return
	}
	common.SuccessJSON(c, "You will receive an email with further instructions.")
}

// PostPublicLogin godoc
// @Summary User login
// @Description User login
// @Tags users
// @Accept json
// @Produce json
// @Success 200
// @Failure 400
// @Router /users [post]
func PostPublicLogin(c *gin.Context) {
	req := new(LoginRequest)
	if err := c.BindJSON(req); err != nil {
		common.Error(c, http.StatusBadRequest, err)
		return
	}
	db := c.MustGet("db").(*clients.SQL)
	user := models.User{Email: req.Email}

	if err := db.First(&user); err != nil {
		common.Error(c, http.StatusUnauthorized, ErrLogin)
		return
	}
	//
	//token, err := auth.Token(&user, "user")
	//if err != nil {
	//	common.Error(c, http.StatusUnauthorized, ErrLogin)
	//	return
	//}

	if err := db.Update(user); err != nil {
		common.Error(c, http.StatusUnauthorized, ErrNotFound)
		return
	}
	common.SuccessJSON(c, gin.H{"token": "ok"})
}
