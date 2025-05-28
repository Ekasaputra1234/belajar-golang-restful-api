package web

type ProductResponse struct {
	// Required Fields
	ID uint `json:"id"`

	// Fields
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Stock    int     `json:"stock"`
	Category string  `json:"category"`
}
