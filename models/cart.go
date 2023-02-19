package models

type Cart struct {
	ID     string `gorm:"index;column:id;primaryKey" json:"id,omitempty"`
	UserID string `gorm:"index;column:user_id" json:"-"`
}

type CartItem struct {
	ID     string `gorm:"index;column:id;primaryKey" json:"id,omitempty"`
	CartID string `gorm:"index;column:cart_id" json:"-"`
	BookID string `gorm:"index;column:book_id" json:"-"`
	Qty    uint64 `gorm:"column:qty" json:"qty"`
}
