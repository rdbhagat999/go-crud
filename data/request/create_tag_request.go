package request

type CreateTagRequest struct {
	Name string `validate:"required,min=5,max=200" json:"name"`
}
