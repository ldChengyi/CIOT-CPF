package entity

// User 用户实体结构
type User struct {
	Model           // 继承自 Model 基础结构，包含 ID、Ext、CreatedAt、UpdatedAt、DeletedAt
	Username string `gorm:"type:varchar(100);not null;unique" json:"username"` // 用户名，最大长度 100，不能为空，唯一
	Password string `gorm:"type:varchar(255);not null" json:"password"`        // 密码，最大长度 255，不能为空
}
