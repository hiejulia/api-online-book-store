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

// UserCode ...
type UserCode struct {
	ID        string `gorm:"column:id;primaryKey" json:"-"`
	CreatedAt uint64 `gorm:"column:created_at" json:"-"`
	UserID    string `gorm:"index;column:builder_id" json:"-"`
	Code      string `gorm:"column:code" json:"-"`
	ExpiresAt uint64 `gorm:"column:expires_at" json:"-"`
}

type PermissionsResponse struct {
	Role        string   `json:"role"`
	Permissions []string `json:"permissions"`
}

//
//// HasRole ...
//func (a *User) HasRole(roleName string) bool {
//	roles := strings.Split(a.Role, constants.TOKEN_ROLE_SEPARATOR)
//	for _, role := range roles {
//		if role == roleName {
//			return true
//		}
//	}
//	return false
//}

// ComparePassword ...
func (a *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(password))
}

//
//// HasPermission ...
//func (a *User) HasPermission(perm uint64) bool {
//	return a.Permissions&perm > 0
//}
//
//// SetPermission ...
//func (a *User) SetPermission(perm uint64) {
//	a.Permissions = a.Permissions | perm
//}
//
//func (builder *Builder) GetPermissionList() PermissionsResponse {
//	result := []string{}
//	for i, perm_str := range LIST_PERMISSIONS_STRINGS {
//		if builder.HasPermission(1 << i) {
//			result = append(result, perm_str)
//		}
//	}
//	return PermissionsResponse{
//		Role:        builder.Role,
//		Permissions: result,
//	}
//}
//
//// UnsetPermission ...
//func (a *Builder) UnsetPermission(perm uint64) {
//	a.Permissions = a.Permissions & (^perm)
//}
//
//// HasPermissions ...
//func (a *Builder) HasPermissions(perms ...uint64) bool {
//	for _, perm := range perms {
//		if !a.HasPermission(perm) {
//			return false
//		}
//	}
//	return true
//}

// HashPassword ...
func (a *User) HashPassword() error {
	pwd, err := bcrypt.GenerateFromPassword([]byte(a.Password), 12)
	if err != nil {
		return err
	}

	a.Password = string(pwd)
	return nil
}
