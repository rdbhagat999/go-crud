package external

type Product struct {
	ID                 int     `json:"id,omitempty"`
	Title              string  `json:"title,omitempty"`
	Price              float64 `json:"price,omitempty"`
	Quantity           int     `json:"quantity,omitempty"`
	Total              float64 `json:"total,omitempty"`
	DiscountPercentage float64 `json:"discountPercentage,omitempty"`
	DiscountedTotal    float64 `json:"discountedTotal,omitempty"`
	Thumbnail          string  `json:"thumbnail,omitempty"`
}
