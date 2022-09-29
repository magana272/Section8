package models

import (
	"github.com/jinzhu/gorm"
)

type Session struct {
	gorm.Model
	Cookie string `gorm:"" json:"cookie" ` // Cookie
	UserId uint   `json:"userid"`          // UserId
}

func (s *Session) CreateSession(u *User) *Session {
	s.UserId = u.ID
	db.NewRecord(s)
	db.Create(&s)
	return s
}

func GetSession(cookie string) *Session {
	var session = Session{}
	db.Where("cookie=?", cookie).Find(&session)
	return &session
}
func DeleteSession(cookie string) {
	db.Where("cookie=?", cookie).Delete(Session{})
}
