package models

type Rating struct {
	ID     string `gorm:"index;column:id;primaryKey" json:"id,omitempty"`
	UserID string `gorm:"column:user_id" json:"userId"`
	BookID string `gorm:"column:book_id" json:"bookId"`
	Rating uint32 `gorm:"column:rating" json:"rating"`
}
