package domain

import (
	"time"

	"github.com/google/uuid"
)

// ASQConcern merepresentasikan kekhawatiran tambahan dari orang tua dalam screening ASQ-3
type ASQConcern struct {
	ID          uuid.UUID     `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	ScreeningID *uuid.UUID    `gorm:"type:uuid" json:"screening_id"`
	Screening   *ASQScreening `gorm:"foreignKey:ScreeningID" json:"screening"`

	HasVisionConcern   bool    `gorm:"type:boolean;default:false" json:"has_vision_concern"`
	HasHearingConcern  bool    `gorm:"type:boolean;default:false" json:"has_hearing_concern"`
	HasBehaviorConcern bool    `gorm:"type:boolean;default:false" json:"has_behavior_concern"`
	HasOtherConcern    bool    `gorm:"type:boolean;default:false" json:"has_other_concern"`
	ConcernDetails     *string `gorm:"type:text" json:"concern_details"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (ASQConcern) TableName() string {
	return "asq_concerns"
}
