package user

// RegisterRequest ...
type RegisterRequest struct {
	Email    string `binding:"required,email" json:"email"`
	Password string `binding:"required" json:"password"`
}

// LoginRequest
type LoginRequest struct {
	Email    string `binding:"required,email" json:"email"`
	Password string `binding:"required" json:"password"`
}
