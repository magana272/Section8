package models

import (
	"github.com/jinzhu/gorm"
	"github.com/magana272/Section8/pkg/config"
)

var db *gorm.DB

// `Home` belongs to `Person`, `OwnerID` is the foreign key
type Home struct {
	gorm.Model
	Owner    string `gorm:"foreignkey:Name;references:Owner"`
	City     string `json:"city"`
	Address  string `json:"address"`
	ZipCode  string `json:"zipcode"`
	UnitType string `json:"unittype"`
}

func init() {
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&Home{}, &Person{})

}

func (h *Home) CreateHome(p *Person) *Home {
	h.Owner = p.Name
	db.NewRecord(h)
	db.Create(&h)
	return h
}
func GetAllHome() []Home {
	var Homes []Home
	db.Find(&Homes)
	return Homes
}
func GetHomeById(id uint) (*Home, *gorm.DB) {
	var getHome Home
	db.Where("HomeID=?", id).Find(&getHome)
	return &getHome, db
}
func DeleteHome(id uint) *Home {
	var getHome Home
	db.Where("HomeID=?", id).Find(&getHome)
	return &getHome
}
