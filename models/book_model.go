package models

// Book structure
type Book struct {
	ID          uint   `gorm:"primary_key" json:"id"`
	Price       int    `gorm:"not null" json:"price"`
	Title       string `gorm:"not null" json:"title"`
	Description string `gorm:"not null" json:"description"`
}
