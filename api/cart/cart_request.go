package cart

type CartItemRequest struct {
	CartID string `gorm:"column:cart_id" json:"-"`
	BookID string `gorm:"index;column:book_id" json:"-"`
	Qty    uint64 `gorm:"column:qty" json:"qty"`
}
