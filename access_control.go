package models

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type AccessToken struct {
	jwt.RegisteredClaims
	UserID     uuid.UUID   `json:"user_id,omitempty"`
	RoleID     uuid.UUID   `json:"role_id,omitempty"`
	CompanyID  *uuid.UUID  `json:"company_id,omitempty"`
	ProgramIDs []uuid.UUID `json:"program_ids,omitempty"`
}

type Resource struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name string
}

type Permission struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name     string
	Group    string
	Contexts pq.StringArray `gorm:"type:text[]"` // Array of contexts

	Resources []Resource `gorm:"many2many:permission_resources;"`
}

type Role struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name        string
	Status      string
	IsSystem    bool
	CompanyID   *uuid.UUID   `gorm:"type:uuid"`
	Permissions []Permission `gorm:"many2many:role_permissions;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
