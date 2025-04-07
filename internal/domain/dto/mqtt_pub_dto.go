package dto

type CommandTransferDTO struct {
	Command int64 `json:"id" xml:"id" form:"id" query:"id" validate:"required,gt=0"`
}
