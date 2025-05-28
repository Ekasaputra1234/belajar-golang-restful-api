package web

type ProductResponse struct {
	// Required Fields
	ID uint `json:"id"`

	// Fields
	Name             string `json:"name"`
	JumlahPeliharaan uint   `json:"jumlah_peliharaan" validate:"gt=0"`
	UnitID           string `json:"unit_id"`
}

type ProductShortResponse struct {
	// Required Fields
	ID uint `json:"id"`

	// Fields
	Name             string `json:"name"`
	JumlahPeliharaan uint   `json:"jumlah_peliharaan" validate:"gt=0"`
}
