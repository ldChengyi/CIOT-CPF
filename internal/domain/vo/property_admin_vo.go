package vo

import "time"

type PropertyAdminVo struct {
	ID           int64     `json:"id"`
	PropertyName string    `json:"property_name"`
	Identifier   string    `json:"identifier"`
	DataType     string    `json:"data_type"`
	Desc         string    `json:"desc"`
	CreatedAt    time.Time `json:"created_at"`
}
