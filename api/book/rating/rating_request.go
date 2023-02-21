package rating

// RateBookRequest ...
type RateBookRequest struct {
	UserId string `gorm:"column:user_id" json:"userId"`
	BookId string `gorm:"column:book_id" json:"bookId"`
	Rating uint32 `gorm:"column:rating" json:"rating"`
}
