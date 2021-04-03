package models

import (
	"database/sql"
	"entities"
)

type User_Model_Model struct {
	Db *sql.DB
}

func (User_Model_Model User_Model_Model) Create(User_Model *entities.User_Model) (int64, error) {
	result, err := User_Model_Model.Db.Exec("insert into User_Model(id, fname, city, phone, height, Married) values(?, ?, ?, ?)", User_Model.id, User_Model.fName, User_Model.city, User_Model.phone, User_Model.height, User_Model.Married)
	if err != nil {
		return 0, err
	} else {
		User_Model.Id, _ = result.LastInsertId()
		rowsAffected, _ := result.RowsAffected()
		return rowsAffected, nil
