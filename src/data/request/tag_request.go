package request

type CreateTagRequest struct {
	Name   string `validate:"required,min=5,max=200" json:"name"`
	UserID int    `validate:"required" json:"user_id"`
}

type UpdateTagRequest struct {
	ID   int    `validate:"required" json:"id"`
	Name string `validate:"required,min=5,max=200" json:"name"`
}

type DeleteTagRequest struct {
	ID int `validate:"required" json:"id"`
}
