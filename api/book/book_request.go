package book

// BookRequest ...
type BookRequest struct {
	Name  string  `gorm:"column:name" json:"name"`
	Price float64 `gorm:"column:price" json:"price"`
	Qty   uint64  `gorm:"column:qty" json:"qty"`
}
