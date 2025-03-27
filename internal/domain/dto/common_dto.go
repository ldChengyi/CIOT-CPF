package dto

type QueryDTO struct {
	Page     int                    `json:"page" xml:"page" form:"page" query:"page"`
	PageSize int                    `json:"page_size" xml:"page_size" form:"page_size" query:"page_size"`
	Filters  map[string]interface{} `json:"filters" xml:"filters" form:"filters" query:"filters"`
}

type IdDeleteDTO struct {
	ID int64 `json:"id" xml:"id" form:"id" query:"id" validate:"required,gt=0"`
}

type IdTransferDTO struct {
	ID int64 `json:"id" xml:"id" form:"id" query:"id" validate:"required,gt=0"`
}
