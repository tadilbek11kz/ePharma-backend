package pharmacy

type GetPharmacyAvailabilityRequest struct {
	Pharmacy
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}
