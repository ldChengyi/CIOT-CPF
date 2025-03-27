package entity

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        int64      `gorm:"primaryKey;autoIncrement" json:"id"`
	Ext       JSONMap    `gorm:"type:json" json:"ext"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at"`
}

// JSONMap 处理 json 类型字段
type JSONMap map[string]interface{}

// Scan 从数据库读取 json 数据
func (j *JSONMap) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("数据类型错误，无法转换为 []byte 类型")
	}
	return json.Unmarshal(bytes, j)
}

// Value 将 JSONMap 转换为 json 数据存储到数据库
func (j JSONMap) Value() (driver.Value, error) {
	if j == nil {
		return "{}", nil
	}
	return json.Marshal(j)
}

// 写入数据库前的默认操作
func (m *Model) BeforeCreate(db *gorm.DB) error {
	currentTime := time.Now()
	m.CreatedAt = currentTime
	m.UpdatedAt = currentTime

	if m.Ext == nil {
		m.Ext = make(map[string]interface{})
	}
	return nil
}

func (m *Model) BeforeUpdate(db *gorm.DB) error {
	m.UpdatedAt = time.Now()
	return nil
}
