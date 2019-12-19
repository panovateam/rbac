package model

type Group struct {
	Name       string    `sql:"type:varchar(150)" json:"name"`
	Uuid       string    `sql:"type:uuid,notnull" json:"uuid"`
	CustomerID int       `sql:"type:bigint,notnull" json:"-"`

	Users []string `json:"users,omitempty"` // many to many relation
	Policies []Policy `json:"policies,omitempty"`
}
