package model

// AuthUser represents data stored in JWT token for user
type AuthUser struct {
	ID             int        `json:"id,omitempty"`
	CustomerID     int        `json:"customerID,omitempty"`
	CustomerNumber string     `json:"customerNumber,omitempty"`
	Clients        []string   `json:"clients,omitempty"`
	Username       string     `json:"username,omitempty"`
	UserUuid       string     `json:"userUuid,omitempty"`
	Role           AccessRole `json:"-"`
}
