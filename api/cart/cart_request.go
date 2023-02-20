package cart

type CartItemRequest struct {
	CartID string `gorm:"column:cart_id" json:"-"`
	BookID string `gorm:"column:book_id" json:"bookId"`
	Qty    uint64 `gorm:"column:qty" json:"qty"`
}
