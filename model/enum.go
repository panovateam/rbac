package model

// SecurityType request or resposne
type SecurityType int8

// NotificationStatus type constant
type NotificationStatus int8

// Mode type constant
type Mode int8

// NotificationType request or resposne
type NotificationType int8

// CallStatus call event
type CallStatus string

// Type of call or message
type Type int8

// Message Type constant
const (
	Motion = iota
	Smoke
	Co
	SOS
	TempHumd
	DoorLock
	OSLocus
)

// Mode Type constant
const (
	Security = iota
	Safety
	Basic
)

// Status type constant
const (
	Initial = iota
	Running
	Done
	Timeout
	Canceled
)

func (s SecurityType) String() string {
	switch s {
	case Co:
		return "Co"
	case Smoke:
		return "Smoke"
	case SOS:
		return "SOS"
	case Motion:
		return "Motion"
	case TempHumd:
		return "TempHumd"
	case DoorLock:
		return "DoorLock"

	case OSLocus:
		return "OS Locus"
	default:
		return "Motion"
	}
}
func (s Mode) String() string {
	switch s {
	case Safety:
		return "Safety"
	case Security:
		return "Security"
	case Basic:
		return "Basic"
	default:
		return "Basic"
	}
}

func (s NotificationStatus) String() string {
	switch s {
	case Running:
		return "Running"
	case Done:
		return "Done"
	case Timeout:
		return "Timeout"
	case Canceled:
		return "Canceled"
	case Initial:
		return "Initial"
	default:
		return "Initial"
	}
}

// Message Type constant
const (
	Simple = iota
	DoorSensor
	MotionSensor
	SecurityBreach
	SafetyBreachCO
	SafetyBreachSOS
	SafetyBreachSmoke
	SecurityAlarm
	SafetyAlarm
	SafetyBreachTempHumd
	Complex
)

func (s NotificationType) String() string {
	switch s {
	case Complex:
		return "Complex"
	case DoorSensor:
		return "DoorSensor"
	case MotionSensor:
		return "MotionSensor"
	case SecurityBreach:
		return "SecurityBreach"
	case SafetyBreachCO:
		return "SafetyBreachCO"
	case SafetyBreachSOS:
		return "SafetyBreachSOS"
	case SafetyBreachSmoke:
		return "SafetyBreachSmoke"
	case SafetyBreachTempHumd:
		return "SafetyBreachTempHumd"
	case SecurityAlarm:
		return "SecurityAlarm"
	case SafetyAlarm:
		return "SafetyAlarm"
	case Simple:
		return "Simple"
	default:
		return "Simple"
	}
}

// Status type constant
const (
	Unknown = "unknown"
	Queued  = "queued"
	Failed  = "failed"
	// = "//" call status
	Initiated = "initiated"
	Ringing   = "ringing"
	Busy      = "busy"
	Answered  = "answered"
	Completed = "completed"
	// = "//" sms status
	Sent        = "sent"
	Delivered   = "delivered"
	Undelivered = "undelivered"
)

func (s CallStatus) String() string {
	switch s {
	case Queued:
		return "Queued"
	case Initiated:
		return "Initiated"
	case Ringing:
		return "Ringing"
	case Busy:
		return "Busy"
	case Answered:
		return "Answered"
	case Completed:
		return "Completed"
	case Sent:
		return "Sent"
	case Delivered:
		return "Delivered"
	case Undelivered:
		return "Undelivered"
	case Failed:
		return "Failed"
	case Unknown:
		return "Unknown"
	default:
		return "Unknown"
	}
}

// Call
const (
	Call = iota
	SMS
)

func (s Type) String() string {
	switch s {
	case SMS:
		return "SMS"
	case Call:
		return "Call"
	default:
		return "Call"
	}
}
