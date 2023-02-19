package models

type Order struct {
	ID        string  `gorm:"index;column:id;primaryKey" json:"id,omitempty"`
	UserID    string  `gorm:"index;column:user_id" json:"-"`
	Price     float64 `gorm:"column:price" json:"price"`
	Status    string  `gorm:"column:status" json:"status"`
	CreatedAt uint64  `gorm:"column:created_at" json:"createdAt,omitempty"`
}

type OrderItem struct {
	ID      string `gorm:"index;column:id;primaryKey" json:"id,omitempty"`
	OrderID string `gorm:"index;column:order_id" json:"-"`
	BookID  string `gorm:"index;column:book_id" json:"-"`
	Qty     uint64 `gorm:"column:qty" json:"qty"`
}
