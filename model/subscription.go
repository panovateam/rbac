package model

import (
	"time"
)

type SubscriptionStatus uint8

const (
	Pending = iota + 1
	Active
	OnHold
	Cancelled
	Expired
)

var SubscriptionStatuss = []SubscriptionStatus{
	Pending,
	Active,
	OnHold,
	Cancelled,
	Expired,
}

func (s SubscriptionStatus) String() string {
	switch s {
	case Pending:
		return "Pending"
	case Active:
		return "Active"
	case OnHold:
		return "OnHold"
	case Cancelled:
		return "Cancelled"
	case Expired:
		return "Expired"
	default:
		return ""
	}
}

//SubscriptionCache cache field customer_number+service_name
type SubscriptionCache struct {
	UUID             string
	PackageID        string
	StartDate        time.Time
	EndDate          time.Time
	Type             uint8
	Meta             map[string]interface{}
	Duration         uint16
	IntervalTime     uint16
	Quota            uint16
	OldPrice         float64
	CustomerNumber   string
	Service          string
	Status           SubscriptionStatus
	HaveTrialPackage bool
	TrialDuration    uint16
}
