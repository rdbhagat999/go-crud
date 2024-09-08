package external

type CartResponse struct {
	Carts []Cart `json:"carts,omitempty"`
	Total int    `json:"total,omitempty"`
	Skip  int    `json:"skip,omitempty"`
	Limit int    `json:"limit,omitempty"`
}

type CartProduct struct {
	ID       int `json:"id,omitempty"`
	Quantity int `json:"quantity,omitempty"`
}
type AddCartRequest struct {
	UserID   int           `json:"userId,omitempty"`
	Products []CartProduct `json:"products,omitempty"`
}

type AddCartResponse struct {
	Products        []Product `json:"products,omitempty"`
	ID              int       `json:"id,omitempty"`
	UserId          int       `json:"userId,omitempty"`
	Total           float64   `json:"total,omitempty"`
	DiscountedTotal int       `json:"discountedTotal,omitempty"`
	TotalProducts   int       `json:"totalProducts,omitempty"`
	TotalQuantity   int       `json:"totalQuantity,omitempty"`
}

type UpdateCartResponse struct {
	Products        []Product `json:"products,omitempty"`
	ID              int       `json:"id,omitempty"`
	UserId          int       `json:"userId,omitempty"`
	Total           float64   `json:"total,omitempty"`
	DiscountedTotal int       `json:"discountedTotal,omitempty"`
	TotalProducts   int       `json:"totalProducts,omitempty"`
	TotalQuantity   int       `json:"totalQuantity,omitempty"`
}

type DeleteCartResponse struct {
	Products        []Product `json:"products,omitempty"`
	ID              int       `json:"id,omitempty"`
	UserId          int       `json:"userId,omitempty"`
	Total           float64   `json:"total,omitempty"`
	DiscountedTotal int       `json:"discountedTotal,omitempty"`
	TotalProducts   int       `json:"totalProducts,omitempty"`
	TotalQuantity   int       `json:"totalQuantity,omitempty"`
	IsDeleted       bool      `json:"isDeleted,omitempty"`
}
