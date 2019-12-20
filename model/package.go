package model

import (
	"time"
)

// PackageCache represents package model cache
type PackageCache struct {
	UUID             string                 ` json:"uuid" validate:"omitempty,uuid4"`
	StartDate        time.Time              `json:"start_date"`
	EndDate          time.Time              `json:"end_date"`
	Price            uint64                 `json:"price"`
	Active           bool                   `json:"active"`
	Type             uint8                  `json:"type"`
	CustomerNumber   string                 `json:"customer_number"`
	Duration         uint16                 `json:"duration"`
	HaveTrialPackage bool                   `json:"have_trial_package"`
	TrialDuration    uint16                 `json:"trial_duration"`
	Meta             map[string]interface{} `json:"meta" validate:"omitempty"`
	IntervalTime     uint16                 `json:"interval_time" validate:"omitempty,gte=0"` // Month
	Quota            uint16                 `json:"quota" validate:"omitempty,gte=0"`
	ServiceName      string                 `pg:"-" json:"service_name" validate:"required"`
}
