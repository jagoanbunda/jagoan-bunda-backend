// Package domain is used to store models / entity that are used for ORM and migration, represents table and its content, and also hook to GORM functions
package domain

import (
	"os"
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

// AfterFind adalah hook GORM yang dieksekusi setelah query SELECT berhasil.
// Gunakan fungsi ini untuk memanipulasi data struct sebelum digunakan oleh aplikasi.
func (u *User) AfterFind(tx *gorm.DB) (err error) {
	if u.ProfilePicture != nil {
		// Contoh logika: Menambahkan base URL ke nama file gambar
		// Dalam implementasi nyata, sebaiknya ambil base URL dari konfigurasi aplikasi
		baseURL := os.Getenv("BASE_URL")
		if baseURL == ""{
			baseURL = "http://0.0.0.0:8080"
		}
		modifiedPicture := baseURL + *u.ProfilePicture
		u.ProfilePicture = &modifiedPicture
	}
	return
}
