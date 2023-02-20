package models

import "golang.org/x/crypto/bcrypt"

// User
type User struct {
	ID         string `gorm:"index;column:id;primaryKey" json:"id,omitempty"`
	Email      string `gorm:"column:email;uniqueIndex;size:255" json:"email,omitempty"`
	Password   string `gorm:"column:password" json:"-"`
	CreatedAt  uint64 `gorm:"column:created_at" json:"createdAt,omitempty"`
	UpdatedAt  uint64 `gorm:"column:updated_at" json:"updatedAt,omitempty"`
	VerifiedAt uint64 `gorm:"column:verified_at" json:"-"`
}

// ComparePassword ...
func (a *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(password))
}

// HashPassword ...
func (a *User) HashPassword() error {
	pwd, err := bcrypt.GenerateFromPassword([]byte(a.Password), 12)
	if err != nil {
		return err
	}

	a.Password = string(pwd)
	return nil
}
