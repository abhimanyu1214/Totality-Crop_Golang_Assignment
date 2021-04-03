package main

import (
	"config"
	"entities"
	"fmt"
	"models"
)

func main() {
  db, err := config.GetMySQLDB()
	if err != nil {
		fmt.Println(err)
	} else {

		User_Model_Model := models.User_Model_Model{
			Db: db,
		}

		User_Model := entities.User_Model{
			id:      1,
			fname:   "Steve",
			city:    "LA",
			phone:   1234567890,
      height:  5.8
      Married: "True",
		}
		rowsAffected, err2 := User_Model_Model.Create(&User_Model)
		if err2 != nil {
			fmt.Println(err2)
		} else {
			fmt.Println("Rows Affected:", rowsAffected)
			fmt.Println("User_Model Info")
			fmt.Println("id:", User_Model.id)
			fmt.Println("fname:", User_Model.fname)
			fmt.Println("city:", User_Model.city)
			fmt.Println("phone:", User_Model.phone)
			fmt.Println("height:", User_Model.height)
      fmt.Println("Married:", User_Model.Married)
		}

	}
}
