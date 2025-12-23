package domain

import (
	"github.com/google/uuid"
)

// ASQResultType adalah enum untuk hasil evaluasi domain dalam ASQ-3
type ASQResultType string

const (
	ASQResultOnTrack     ASQResultType = "on_track"
	ASQResultMonitoring  ASQResultType = "monitoring"
	ASQResultBelowCutoff ASQResultType = "below_cutoff"
)

// ASQDomainResult merepresentasikan hasil per domain dalam screening ASQ-3 (5 hasil per screening)
type ASQDomainResult struct {
	ID          uuid.UUID     `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	ScreeningID *uuid.UUID    `gorm:"type:uuid" json:"screening_id"`
	Screening   *ASQScreening `gorm:"foreignKey:ScreeningID" json:"screening"`

	Domain         ASQDomain     `gorm:"type:varchar(20);not null" json:"domain"`
	TotalScore     float64       `gorm:"type:decimal(4,1);not null" json:"total_score"`  // Sum of 6 items (max 60)
	CutoffScore    float64       `gorm:"type:decimal(4,1);not null" json:"cutoff_score"` // From cutoffs table
	MonitoringZone float64       `gorm:"type:decimal(4,1);not null" json:"monitoring_zone"`
	Result         ASQResultType `gorm:"type:varchar(20);not null" json:"result"` // on_track, monitoring, below_cutoff
}

func (ASQDomainResult) TableName() string {
	return "asq_domain_results"
}
