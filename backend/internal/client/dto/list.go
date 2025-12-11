package dto

type GetClientsQuery struct {
	Limit int     `query:"limit" validate:"omitempty,min=1,max=100"`
	Page  int     `query:"page"  validate:"omitempty,min=1"`
	Sort  *string `query:"sort"`
}
