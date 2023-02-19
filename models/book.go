package models

type Book struct {
	ID    string  `gorm:"index;column:id;primaryKey" json:"id,omitempty"`
	Name  string  `gorm:"column:name" json:"name"`
	Price float64 `gorm:"column:price" json:"price"`
	Qty   uint64  `gorm:"column:qty" json:"qty"`
}
