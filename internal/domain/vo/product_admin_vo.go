package vo

import "time"

type ProductAdminVo struct {
	ID          int64     `json:"id"`
	ProductKey  string    `json:"product_key"`
	ProductName string    `json:"product_name"`
	Desc        string    `json:"desc"`
	CreatedAt   time.Time `json:"created_at"`
}
