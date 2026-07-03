package entity

import (
	"ostadbun/pkg/hash"
)

type User struct {
	Id           int  `json:"id"`
	UniversityId *int `json:"university_id"`
	MajorId      *int `json:"major_id"`
	AdminBy      *int `json:"admin_by"`
	//must be hash
	Email string `json:"email"`
	Name  string `json:"name"`
}

func (u *User) Email_Hashe() string {
	return hash.Hasher(u.Email)
}
