package models

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	FirstName   string
	MiddleName  string
	LastName    string
	Email       string
	PhoneNumber string
	Password    string
	ImageUrl    string
	Status      string

	RoleID    uuid.UUID `gorm:"type:uuid"`
	CompanyID uuid.UUID `gorm:"type:uuid"`

	Role Role `gorm:"foreignkey:RoleID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (u *User) AfterFind(tx *gorm.DB) (err error) {
	u.ImageUrl, err = GenerateSignedURL(u.ImageUrl)
	return err
}
