package main

import (
	"github.com/sas-scoul/bulk-api/database"
	"github.com/sas-scoul/bulk-api/items/customer"
	"github.com/sas-scoul/bulk-api/items/user"
)

func main() {

	err := database.DB.AutoMigrate(user.User{}, customer.Customer{})
	if err != nil {
		panic("Error migrating: " + err.Error())
	}

}
