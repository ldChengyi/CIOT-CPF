package entity

// Product 表示产品信息，继承了 Model 基础结构，包含了与产品相关的字段
type Product struct {
	Model
	UserID        int64  `json:"user_id"`
	ProductKey    string `gorm:"uniqueIndex;not null;size:255" json:"product_key"`
	ProductName   string `gorm:"not null;size:255" json:"product_name"`
	ProductSecret string `gorm:"not null;size:255" json:"product_secret"`
	Desc          string `gorm:"type:text" json:"desc"`
}
