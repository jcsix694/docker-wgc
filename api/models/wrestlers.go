package models

import (
	"wgcapi/database"
	"wgcapi/logging"
	"wgcapi/structures"

	_ "github.com/go-sql-driver/mysql"
)

func GetAllWrestlers(w *[]structures.Wrestler) ([]structures.Wrestler, error) {
	var wrestler []structures.Wrestler

	if err := database.DB.Find(&wrestler).Error; err != nil {
		logging.Error(err.Error())
		return wrestler, err
	} else {
		return wrestler, nil
	}
}
