package products

import "github.com/jinzhu/gorm"

// Product model
type Product struct {
	gorm.Model
	Name  string
	Stock int32
}
