package dto

type ModuleDTO struct {
	ID          int64   `json:"id" xml:"id" form:"id" query:"id"`
	ModuleName  string  `json:"module_name" xml:"module_name" form:"module_name" query:"module_name"`
	Desc        string  `json:"desc" xml:"desc" form:"desc" query:"desc"`
	PropertyIDs []int64 `json:"property_ids" xml:"property_ids" form:"property_ids" query:"property_ids"`
}
