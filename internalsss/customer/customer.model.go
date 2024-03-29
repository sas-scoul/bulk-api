package customer

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name    string `json:"name"`
	Contact string `json:"contact"`
}
