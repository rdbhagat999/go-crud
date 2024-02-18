package request

type UpdateTagRequest struct {
	Id   int    `validate:"required" json:"id"`
	Name string `validate:"required,min=5,max=200" json:"name"`
}
