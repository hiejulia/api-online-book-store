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

// @Summary User register
// @ID user-register
// @Produce json
// @Success 200 {object}
// @Router /users/register
func PostPublicRegister(c *gin.Context) {
	req := new(RegisterRequest)
	if err := c.BindJSON(req); err != nil {
		common.Error(c, http.StatusBadRequest, err)
		return
	}

	var err error
	db := c.MustGet("db").(*clients.SQL)

	//check for duplicate email and builder name
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

// @Summary User login
// @ID user-login
// @Produce json
// @Success 200 {object}
// @Router /users/login
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

	token, err := Token(&user, "user")
	if err != nil {
		common.Error(c, http.StatusUnauthorized, ErrLogin)
		return
	}

	if err := db.Update(user); err != nil {
		common.Error(c, http.StatusUnauthorized, ErrNotFound)
		return
	}
	common.SuccessJSON(c, gin.H{"token": token})
}
