package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"" json:"name"`
	DOB      string `json:"dob`
	Email    string `json:"email`
	UserName string `json:"username"`
	Password string `json:password`
}

func getUserWithID(id int) *User {
	var getUser User
	db.Where("id=?", id).Find(&getUser)
	return &getUser
}

//	func getUserWithSession(Session string) User {
//		var getUser User
//		db.Where("id=?", ).Find(getUser)
//		return getUser
//	}
func GetAllUsers() []User {
	var allUsers []User
	db.Find(&allUsers)
	return allUsers

}
func (u *User) CreateUser() User {
	db.NewRecord(u)
	db.Create(&u)
	return *u
}
func DeleteUser(id int) {
	//  form doc
	var DeleteU User
	db.Delete(DeleteU, id)
}
func GetUserWithCookie(c string) (*User, error) {
	var currSession = Session{}
	var currUser = &User{}

	db.Where("cookie=?", c).Find(&currSession)
	currUser = getUserWithID(int(currSession.UserId))
	fmt.Println(currUser)
	return currUser, nil

}
func GetUserWithEmail(email string) (*User, error) {
	var currUser = &User{}
	db.Where("email=?", email).Find(currUser)
	return currUser, nil

}
