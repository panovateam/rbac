package model

import (
	"time"
)
// PackageCache represents package model cache
type PackageCache struct {
	StartDate 			time.Time 		`json:"start_date"`
	EndDate 			time.Time 		`json:"end_date"`
	Price 				float64 		`json:"price"`
	Active 				bool 			`json:"active"`
	Type 				uint 			`json:"type"`
	CustomerNumber  	string 	 		`json:"customer_number"`
	Duration        	uint16			`json:"duration"`
	HaveTrialPackage	bool 			`json:"have_trial_package"`
	TrialDuration 		uint16			`json:"trial_duration"`
}