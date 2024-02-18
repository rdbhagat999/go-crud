package response

type TagResponse struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	UserID int    `json:"user_id"`
}
