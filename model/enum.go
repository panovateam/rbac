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

// Datatype of a property of thing
type DataType int8

const (
	Number = iota + 1
	String
	Json
	Image
	Boolean
	ThingObject
	XML
	DateTime
	Location
	Schedule
	HTML
)

func (s DataType) String() string {
	switch s {
	case Number:
		return "Number"
	case String:
		return "String"
	case Json:
		return "Json"
	case Image:
		return "Image"
	case Boolean:
		return "Boolean"
	case ThingObject:
		return "ThingObject"
	case XML:
		return "XML"
	case DateTime:
		return "DateTime"
	case Location:
		return "Location"
	case Schedule:
		return "Schedule"
	case HTML:
		return "HTML"
	default:
		return "Unknown"
	}
}
func GetDataTypeList() []Configure {
	var result = make([]Configure, 11)
	for i := 1; i < 12; i++ {
		id := i
		dataType := DataType(i)
		result[i-1] = Configure{
			&id,
			dataType.String(),
		}
	}

	return result
}
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
	OSLocusTemp
	OSLocusSOS
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
	case OSLocusTemp:
		return "OSLocusTemp"
	case OSLocusSOS:
		return "OSLocusSOS"
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
