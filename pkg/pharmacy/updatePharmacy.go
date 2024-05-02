package pharmacy

type UpdatePharmacyRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Location    string `json:"location"`
}
