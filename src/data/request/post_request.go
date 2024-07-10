package request

type CreatePostRequest struct {
	Title  string `validate:"required,min=5,max=255" json:"title"`
	Body   string `validate:"required,min=5,max=255" json:"body"`
	UserID int    `validate:"required" json:"user_id"`
}

type UpdatePostRequest struct {
	// ID    int    `validate:"required" json:"post_id"`
	Title string `validate:"required,min=5,max=255" json:"title"`
	Body  string `validate:"required,min=5,max=255" json:"body"`
}

type DeletePostRequest struct {
	ID int `validate:"required" json:"id"`
}
