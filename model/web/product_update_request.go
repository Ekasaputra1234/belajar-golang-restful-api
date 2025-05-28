package web

type ProductUpdateRequest struct {
	// Fields
	Name             string `json:"name"`
	JumlahPeliharaan uint   `json:"jumlah_peliharaan" validate:"gt=0"`
	UnitID           string `json:"unit_id"`
}
