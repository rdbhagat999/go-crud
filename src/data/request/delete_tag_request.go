package request

type DeleteTagRequest struct {
	Id int `validate:"required" json:"id"`
}
