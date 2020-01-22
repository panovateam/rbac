package model

/*
Notification logs
- use security type  (safety,security) and gateway mac address to identity job id.
*/
type Notification struct {
	tableName struct{} `pg:",discard_unknown_columns"`
	Base
	ThingName          string             `json:"thingName,omitempty"`
	ThingDisplayName   string             `json:"thingDisplayName,omitempty"`
	ThingSerial        string             `json:"thing_serial,omitempty"`
	GatewayName        string             `json:"gatewayName,omitempty"`
	GatewayDisplayName string             `json:"gatewayDisplayName,omitempty"`
	GatewayMacAddress  string             `json:"gatewayMacAddress,omitempty"`
	ZoneName           string             `json:"zoneName,omitempty"`
	ZoneDisplayName    string             `json:"zoneDisplayName,omitempty"`
	Template           string             `json:"template,omitempty"`
	Status             NotificationStatus `sql:", default:0" json:"status,omitempty"`
	Type               NotificationType   `sql:", default:1" json:"type,omitempty"`
	CustomerNumber     string             `json:"customerNumber,omitempty"`
	MessageKey         string             `json:"message_key,omitempty"`
	State              int                `json:"state,omitempty"`
	DeviceID           int8               `json:"deviceId,omitempty"`
	Timezone           string             `json:"timezone,omitempty"`

	Localizes []Localize `sql:"-" pg:"fk:notification_id" json:"localizes,omitempty"`
}

// Localize logs
type Localize struct {
	Message string `json:"message"`
	Locale  string `json:"locale,omitempty"`
}
