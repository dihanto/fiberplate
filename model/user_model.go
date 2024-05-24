package model

import "github.com/google/uuid"

type User struct {
	ID   uuid.UUID `gorm:"primaryKey" json:"id"`
	Name string    `gorm:"not null" json:"name" validate:"required"`
	Age  uint16    `gorm:"" json:"age"`
}

func (user *User) New(remote User) {
	user.ID = uuid.New()
	user.Name = remote.Name
	user.Age = remote.Age
}
