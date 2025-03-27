package vo

import "time"

type ModuleAdminVo struct {
	ID          int64     `json:"id"`
	ModuleName  string    `json:"module_name"`
	Desc        string    `json:"desc"`
	PropertyIDs []int64   `json:"property_ids"`
	CreatedAt   time.Time `json:"created_at"`
}
