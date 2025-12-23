package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRole string

const (
	RoleParent UserRole = "parent"
	RoleNakes  UserRole = "nakes"
	RoleAdmin  UserRole = "admin"
)

type User struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"uuid"`

	Name           string   `gorm:"type:varchar(255);not null" json:"name"`
	ProfilePicture *string  `gorm:"type:varchar(500)" json:"profile_picture"`
	Email          string   `gorm:"type:varchar(255);not null;uniqueIndex" json:"email"`
	Phone          *string  `gorm:"type:varchar(20);unique" json:"phone"`
	PasswordHash   string   `gorm:"type:varchar(255);not null" json:"-"`
	Role           UserRole `gorm:"type:varchar(20);default:'nakes';not null" json:"role"`
	IsVerified     bool     `gorm:"default:false" json:"is_verified"`
	Address        *string  `gorm:"type:text" json:"address"`

	// nakes
	NIK *string `gorm:"type:varchar(255)" json:"nik"`

	// nakes <- parent
	SupervisorID *uuid.UUID `gorm:"type:uuid" json:"supervisor_id"`
	Supervisor   *User      `gorm:"foreignKey:SupervisorID" json:"supervisor"`

	AssignedParents []User `gorm:"foreignKey:SupervisorID" json:"assigned_parents"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (User) TableName() string {
	return "users"
}
