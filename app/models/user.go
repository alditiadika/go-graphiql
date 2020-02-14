package models

//UserModel struct
type UserModel struct {
	ID           uint16      `json:"id,omitempty"`
	Name         string      `json:"name,omitempty"`
	IsActive     bool        `json:"is_active,omitempty"`
	CreatedBy    interface{} `json:"created_by,omitempty"`
	ModifiedBy   interface{} `json:"modified_by,omitempty"`
	CreatedDate  interface{} `json:"created_date,omitempty"`
	ModifiedDate interface{} `json:"modified_date,omitempty"`
}
