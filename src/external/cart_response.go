package external

type CartResponse struct {
	Carts []Cart `json:"carts,omitempty"`
	Total int    `json:"total,omitempty"`
	Skip  int    `json:"skip,omitempty"`
	Limit int    `json:"limit,omitempty"`
}
