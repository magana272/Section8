package models

import (
	"github.com/jinzhu/gorm"
)

type Person struct {
	gorm.Model
	// OwnerID uint   `gorm:"primaryKey" json:"OwnerID `
	Name  string `json:"name" ` // CompanyName
	Email string `json:"email"` // email
	Phone string `json:"phone"` // Phone
	Homes []Home `json:"home"`
}

func (p *Person) CreatePerson() *Person {
	db.NewRecord(p)

	db.Create(&p)
	return p
}
func GetAllPeople() []Person {
	var People []Person
	db.Find(&People)
	return People
}
func GetPersonByName(name string) *Person {
	// This will break if name > 1
	var getPerson Person
	db.Where("name=?", name).Find(&getPerson)
	return &getPerson
}
