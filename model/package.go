package model

import (
	"time"
)

// PackageCache represents package model cache
type PackageCache struct {
	UUID             string
	StartDate        time.Time
	EndDate          time.Time
	Price            uint64
	Active           bool
	Type             uint8
	CustomerNumber   string
	Duration         uint16
	HaveTrialPackage bool
	TrialDuration    uint16
	Meta             map[string]interface{}
	IntervalTime     uint16
	Quota            uint16
	ServiceName      string
}
