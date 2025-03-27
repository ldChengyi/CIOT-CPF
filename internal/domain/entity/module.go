package entity

type Module struct {
	Model
	ModuleName string `gorm:"not null;size:255" json:"module_name"`
	Desc       string `json:"desc"`
}
