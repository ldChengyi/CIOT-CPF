package dto

type PropertyDTO struct {
	ID           int64  `json:"id" xml:"id" form:"id" query:"id"`
	PropertyName string `json:"property_name" xml:"property_name" form:"property_name" query:"property_name"`
	Identifier   string `json:"identifier" xml:"identifier" form:"identifier" query:"identifier"`
	DataType     string `json:"data_type" xml:"data_type" form:"data_type" query:"data_type"`
	Desc         string `json:"desc" xml:"desc" form:"desc" query:"desc"`
}
