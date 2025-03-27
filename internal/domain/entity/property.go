package entity

type Property struct {
	Model
	PropertyName string `gorm:"not null;size:255" json:"property_name"`
	Identifier   string `gorm:"not null;size:255" json:"identifier"`
	DataType     string `gorm:"not null;size:255" json:"data_type"`
	Desc         string `json:"desc"`
}
