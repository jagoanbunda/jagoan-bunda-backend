package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// EducationArticle merepresentasikan artikel edukasi untuk orang tua
type EducationArticle struct {
	ID           uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Title        string     `gorm:"type:varchar(255);not null" json:"title"`
	Category     string     `gorm:"type:varchar(50);not null" json:"category"`
	Content      string     `gorm:"type:text;not null" json:"content"`
	ThumbnailURL *string    `gorm:"type:varchar(500)" json:"thumbnail_url"`
	PublishedAt  *time.Time `gorm:"type:timestamp" json:"published_at"`
	ViewCount    int        `gorm:"type:integer;default:0" json:"view_count"`
	IsActive     bool       `gorm:"type:boolean;default:true" json:"is_active"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (EducationArticle) TableName() string {
	return "education_articles"
}
