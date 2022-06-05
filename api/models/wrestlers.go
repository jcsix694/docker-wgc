package models

import (
	"wgcapi/database"
	"wgcapi/logging"
	"wgcapi/requests"
	"wgcapi/structures"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

func CreateWrestlerData(w *requests.CreateWrestlerRequest) (structures.Wrestler, error) {
	wrestler := structures.Wrestler{Name: w.Name, UUID: uuid.New()}

	if err := database.DB.Create(&wrestler).Error; err != nil {
		return wrestler, err
	} else {
		return wrestler, nil
	}
}

func GetAllWrestlersData(w *[]structures.Wrestler) ([]structures.Wrestler, error) {
	var wrestler []structures.Wrestler

	if err := database.DB.Find(&wrestler).Error; err != nil {
		logging.Error(err.Error())
		return wrestler, err
	} else {
		return wrestler, nil
	}
}

func GetWrestlerData() {

}
