package external

type Cart struct {
	ID              int       `json:"id,omitempty"`
	Products        []Product `json:"products,omitempty"`
	Total           float64   `json:"total,omitempty"`
	DiscountedTotal float64   `json:"discountedTotal,omitempty"`
	UserID          int       `json:"userId,omitempty"`
	TotalProducts   int       `json:"totalProducts,omitempty"`
	TotalQuantity   int       `json:"totalQuantity,omitempty"`
}
