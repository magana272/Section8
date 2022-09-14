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
	Homes Home   `json:"home"`
}

func (p *Person) CreatePerson() *Person {
	db.NewRecord(p)

	db.Create(&p)
	return p
}

// Get all People
func GetAllPeople() []Person {
	var People []Person
	db.Find(&People)
	return People
}

// Get a Person by Name
func GetPersonByName(name string) *Person {
	// This will break if name > 1
	var getPerson Person
	db.Where("name=?", name).Find(&getPerson)
	return &getPerson
}

// Deleteing a Person with ID
func DeletePersonWithId(id int) *Person {
	var deletePerson Person
	db.Where("id=?", id).Delete(&deletePerson)
	return &deletePerson
}

// GET Person
func GetPersonWithId(id int) *Person {
	var GetPerson Person
	db.Where("id=?", id).Find(&GetPerson)
	return &GetPerson
}

// Update All parametes
func (p *Person) UpdatePerson(id int) *Person {
	var updatedPerson Person
	personFromdb := GetPersonWithId(id)
	if personFromdb.Email != p.Email && p.Email != "" {
		db.Model(&Person{}).Where("id=?", id).Update("email", p.Email)
	}
	if personFromdb.Name != p.Name && p.Name != "" {
		db.Model(&Person{}).Where("id=?", id).Update("name", p.Name)
	}

	if personFromdb.Phone != p.Phone && p.Phone != "" {
		db.Model(&Person{}).Where("id=?", id).Update("phone", p.Phone)
	}
	db.Model(&Person{}).Where("id=?", id).Find(&updatedPerson)
	return &updatedPerson
}
